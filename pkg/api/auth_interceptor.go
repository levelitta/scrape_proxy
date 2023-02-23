package api

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var authFailedErr = fmt.Errorf("auth failed")

type AuthInterceptor struct {
	token string
}

func NewAuthInterceptor(token string) *AuthInterceptor {
	return &AuthInterceptor{
		token: token,
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, authFailedErr
		}

		tokens := md.Get("token")
		if len(tokens) < 1 {
			return nil, authFailedErr
		}
		if interceptor.token != tokens[0] {
			return nil, authFailedErr
		}

		return handler(ctx, req)
	}
}
