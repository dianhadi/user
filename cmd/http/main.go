package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dianhadi/user/pkg/database"
	"github.com/dianhadi/user/pkg/log"
	"github.com/dianhadi/user/pkg/redis"

	"github.com/dianhadi/user/internal/config"
	"github.com/dianhadi/user/internal/handler/helper"
	handlerAuth "github.com/dianhadi/user/internal/handler/http/auth"
	handlerUser "github.com/dianhadi/user/internal/handler/http/user"
	repoAuth "github.com/dianhadi/user/internal/repo/auth"
	repoUser "github.com/dianhadi/user/internal/repo/user"
	usecaseAuth "github.com/dianhadi/user/internal/usecase/auth"
	usecaseUser "github.com/dianhadi/user/internal/usecase/user"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.elastic.co/apm/module/apmchi"
)

const (
	serviceName = "user-service"
)

func main() {
	log.New(serviceName)

	log.Info("Get Configuration")
	appConfig, err := config.GetConfig("config/main.yaml")
	if err != nil {
		panic(err)
	}
	err = config.LoadPublicKey("files/public_key.pem")
	if err != nil {
		panic(err)
	}

	log.Info("Connect to Database")
	db := database.New(appConfig.Database.Host, appConfig.Database.Port, appConfig.Database.Username, appConfig.Database.Password, appConfig.Database.DBName)

	log.Info("Connect to Redis")
	redis, err := redis.New(appConfig.Redis.Host, appConfig.Redis.Port, appConfig.Redis.Password)
	if err != nil {
		panic(err)
	}

	log.Info("Init Repo")
	repoUser, err := repoUser.New(db, redis)
	if err != nil {
		panic(err)
	}
	repoAuth, err := repoAuth.New(redis)
	if err != nil {
		panic(err)
	}

	log.Info("Init Usecase")
	usecaseAuth, err := usecaseAuth.New(repoUser, repoAuth)
	if err != nil {
		panic(err)
	}
	usecaseUser, err := usecaseUser.New(repoUser)
	if err != nil {
		panic(err)
	}

	log.Info("Init Handler")
	handlerAuth, err := handlerAuth.New(usecaseAuth)
	if err != nil {
		panic(err)
	}
	handlerUser, err := handlerUser.New(usecaseUser)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(apmchi.Middleware())
	r.Use(helper.Common)
	r.Use(helper.Recover)

	authToken := chi.Chain(handlerAuth.AuthenticateMiddleware)
	jwtToken := chi.Chain(helper.JWTMiddleware)

	log.Info("Register Route")
	r.Post("/v1/register", handlerUser.Register)
	r.Post("/v1/login", handlerAuth.Login)
	r.Post("/v1/authenticate", handlerAuth.Authenticate) // for external service

	r.Get("/v1/user/{id:[0-9a-f]{8}-(?:[0-9a-f]{4}-){3}[0-9a-f]{12}}", handlerUser.GetUserByID)
	r.With(jwtToken...).Get("/v1/user/{username}", handlerUser.GetUserByUsername)

	r.With(authToken...).Patch("/v1/change-password", handlerUser.ChangePassword)

	r.Handle("/metrics", promhttp.Handler())

	log.Infof("Starting server on port %s...", appConfig.Server.Port)
	startServer(":"+appConfig.Server.Port, r)
}

func startServer(port string, r http.Handler) {

	srv := http.Server{
		Addr:    port,
		Handler: r,
	}

	// Create a channel that listens on incomming interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	// Graceful shutdown
	go func() {
		// Wait for a new signal on channel
		<-signalChan
		// Signal received, shutdown the server
		log.Info("shutting down..")

		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		srv.Shutdown(ctx)

		// Check if context timeouts, in worst case call cancel via defer
		select {
		case <-time.After(21 * time.Second):
			log.Info("not all connections done")
		case <-ctx.Done():
		}
	}()

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
