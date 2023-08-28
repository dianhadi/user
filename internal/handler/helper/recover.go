package helper

import (
	"fmt"
	"net/http"

	"github.com/dianhadi/golib/log"

	"github.com/go-stack/stack"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				ctx := r.Context()

				err, ok := p.(error)
				if !ok {
					err = fmt.Errorf("%v", p)
				}

				var stackTrace stack.CallStack
				// Get the current stacktrace but trim the runtime
				traces := stack.Trace().TrimRuntime()

				// Format the stack trace removing the clutter from it
				for i := 0; i < len(traces); i++ {
					t := traces[i]
					tFunc := t.Frame().Function

					// Opentelemetry is recovering from the panics on span.End defets and throwing them again
					// we don't want this noise to appear on our logs
					if tFunc == "runtime.gopanic" || tFunc == "go.opentelemetry.io/otel/sdk/trace.(*span).End" {
						continue
					}

					// This call is made before the code reaching our handlers, we don't want to log things that are coming before
					// our own code, just from our handlers and donwards.
					if tFunc == "net/http.HandlerFunc.ServeHTTP" {
						break
					}
					stackTrace = append(stackTrace, t)
				}

				fields := log.Fields{
					"request-id": ctx.Value("request-id"),
					"stack":      fmt.Sprintf("%+v", stackTrace),
				}
				log.ErrorWithFields(err.Error(), fields)

				log.Info("Recovered from panic:", p)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
