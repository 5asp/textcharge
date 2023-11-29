package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	addpb "github.com/aheadIV/textcharge/t_platform/pb"
	"github.com/aheadIV/textcharge/user-service/publishing"
	"github.com/go-kit/kit/examples/addsvc/pkg/addendpoint"
	"github.com/go-kit/kit/examples/addsvc/pkg/addservice"
	"github.com/go-kit/kit/examples/addsvc/pkg/addtransport"
	kitlog "github.com/go-kit/log"
	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	configPath := filepath.Clean("./app.yaml")

	var config *viper.Viper
	{
		config = viper.New()
		config.SetConfigFile(configPath)

		if err := config.ReadInConfig(); err != nil {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
	var repo rel.Repository
	{
		adapter, _ := postgres.Open(config.GetString("db.testdb"))

		// initialize rel's repo.
		repo = rel.New(adapter)
	}

	var logger kitlog.Logger
	{
		logger = kitlog.NewJSONLogger(os.Stderr)
		logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	}
	var (
		service     = addservice.New(logger, ints, chars)
		endpoints   = addendpoint.New(service, logger, duration, tracer, zipkinTracer)
		httpHandler = addtransport.NewHTTPHandler(endpoints, tracer, zipkinTracer, logger)
		grpcServer  = addtransport.NewGRPCServer(endpoints, tracer, zipkinTracer, logger)
	)

	publishing.RegisterRPCService(logger, repo)

	// 设置并发控制
	// concurrency := runtime.NumCPU()
	// var wg sync.WaitGroup
	// wg.Add(concurrency)

	// 启动并发的RPC服务
	var g group.Group
	{
		rpc.HandleHTTP()
		addr := config.GetString("account.rpc")
		l, e := net.Listen("tcp", addr)
		if e != nil {
			logger.Log("listen error:", e)
		}
		g.Add(func() error {
			logger.Log("account", "RPC", "addr", addr)
			return http.Serve(l, nil)
		}, func(error) {
			l.Close()
		})
	}

	{
		grpcServer := addtransport.NewGRPCServer(endpoints, tracer, zipkinTracer, logger)

		platformRpc := config.GetString("account.rpc")
		// The gRPC listener mounts the Go kit gRPC server we created.
		grpcListener, err := net.Listen("tcp", platformRpc)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", platformRpc)
			// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
			// the here demonstrated zipkin tracing middleware.
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			addpb.RegisterTappInfoServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	logger.Log("exit", g.Run())
}
