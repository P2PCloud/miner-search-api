package middleware

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCRecover interceptor.
func GRPCRecover(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
	defer func() {
		if p := recover(); p != nil {
			errMsg := fmt.Sprintf("recovered from panic: %v", p)
			err = status.Errorf(codes.Internal, errMsg)
			logrus.New().WithError(errors.New(errMsg)).WithFields(
				logrus.Fields{
					"stack_trace": string(debug.Stack()),
					"panic":       true,
				},
			)

			// TODO metric Count panic recovery
		}
	}()
	return handler(ctx, req)
}
