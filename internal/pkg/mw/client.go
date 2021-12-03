package mwclient

import (
	"context"
	"github.com/ozonmp/omp-bot/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	appNameHeader    = "x-app-name"
	appVersionHeader = "x-app-version"
)

// AddAppInfoUnary add info about client
func AddAppInfoUnary(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.AppendToOutgoingContext(ctx, appNameHeader, config.AppName)
	ctx = metadata.AppendToOutgoingContext(ctx, appVersionHeader, config.AppVersion)
	return invoker(ctx, method, req, reply, cc, opts...)
}
