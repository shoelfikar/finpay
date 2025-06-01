package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const SECRET_KEY = "152f895991a4924be4735227a4701756"


func AuthInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
    }

    keys := md.Get("x-api-key")
    if len(keys) == 0 || keys[0] != SECRET_KEY {
        return nil, status.Errorf(codes.Unauthenticated, "invalid API key")
    }

    // Authorized; continue
    return handler(ctx, req)
}