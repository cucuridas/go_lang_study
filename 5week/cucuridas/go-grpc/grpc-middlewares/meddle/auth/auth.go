package auth

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CustomAuthFunc(ctx context.Context)(context.Context,error) {
	token, err := grpc_auth.AuthFromMD(ctx,"bearer")

	if err != nil {
		return nil, err
	}

	if token != "customToken" {
		return nil, status.Errorf(codes.Unauthenticated,"invalid auth token: %v",err)
	}

	newCtx := context.WithValue(ctx,"token",token)

	return newCtx, nil
}
