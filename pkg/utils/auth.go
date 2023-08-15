package utils

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// should be placed somewhere
var secretKey = []byte("secret12345")

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func GenerateToken(username string) string {
	tokenBytes := make([]byte, 32)
	rand.Read(tokenBytes)
	return base64.StdEncoding.EncodeToString(tokenBytes)
}

func GetTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		return authHeader[len("Bearer "):]
	}
	return ""
}
