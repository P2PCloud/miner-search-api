package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

// HTTPRecover middleware
func HTTPRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				logrus.New().WithError(fmt.Errorf("%v", p)).WithFields(
					logrus.Fields{
						"stack_trace": string(debug.Stack()),
					},
				)

				w.WriteHeader(http.StatusInternalServerError)
				buf := bytes.NewBuffer([]byte(fmt.Sprintf("recovered from panic: %v\n", p)))
				_, _ = buf.Write(debug.Stack())
				_, _ = w.Write(buf.Bytes())

				// TODO metric Count panic recovery
			}
		}()
		next.ServeHTTP(w, r)
	})
}
