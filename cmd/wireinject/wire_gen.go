// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wireinject

import (
	"context"
	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/hyper/provider/http"
	"github.com/hyper-micro/hyper/provider/rpc"
	"github.com/hyper-micro/project-layout/internal/handler/restful"
	rpc2 "github.com/hyper-micro/project-layout/internal/handler/rpc"
	"github.com/hyper-micro/project-layout/internal/repository"
	"github.com/hyper-micro/project-layout/internal/server"
	"github.com/hyper-micro/project-layout/internal/service"
	"github.com/hyper-micro/project-layout/internal/service/svcctx"
)

// Injectors from wireinject.go:

func NewHttpServer(ctx context.Context, conf config.Config) (*server.HttpServer, func(), error) {
	provider := http.NewProvider(conf)
	accountRepository := repository.NewAccountRepository()
	greeterServiceCtx := svcctx.NewGreeterServiceCtx(ctx, accountRepository)
	greeterService := service.NewGreeterService(greeterServiceCtx)
	greeterRestfulHandler := restful.NewGreeterRestfulHandler(greeterService)
	restfulHandlerSet := &server.RestfulHandlerSet{
		Greeter: greeterRestfulHandler,
	}
	httpServer := server.NewHttpServer(conf, provider, restfulHandlerSet)
	return httpServer, func() {
	}, nil
}

func NewRpcServer(ctx context.Context, conf config.Config) (*server.RpcServer, func(), error) {
	provider := rpc.NewProvider(conf)
	accountRepository := repository.NewAccountRepository()
	greeterServiceCtx := svcctx.NewGreeterServiceCtx(ctx, accountRepository)
	greeterService := service.NewGreeterService(greeterServiceCtx)
	greeterRpcServerHandler := rpc2.NewGreeterRpcServerHandler(greeterService)
	rpcHandlerSet := &server.RpcHandlerSet{
		Greeter: greeterRpcServerHandler,
	}
	rpcServer := server.NewRpcServer(conf, provider, rpcHandlerSet)
	return rpcServer, func() {
	}, nil
}
