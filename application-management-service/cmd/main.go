package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/aheadIV/textcharge/application-management-service/publishing"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/handlers"

	"github.com/go-rel/postgres"

	"github.com/go-rel/rel"
	"github.com/spf13/viper"
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

	publishing.RegisterRPCService(logger, repo)

	// 设置并发控制
	concurrency := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(concurrency)

	// 启动并发的RPC服务
	go func() {
		defer wg.Done()
		rpc.HandleHTTP()
		l, e := net.Listen("tcp", config.GetString("app.rpc"))
		if e != nil {
			logger.Log("listen error:", e)
		}
		logger.Log("service:", "app was started")
		err := http.Serve(l, nil)
		if err != nil {
			logger.Log("serve error:", err)
		}
	}()

	// 启动并发的HTTP服务
	go func() {
		defer wg.Done()
		service := publishing.New(logger, config)
		r := publishing.MakeHttpHandler(logger, service)
		loggedRouter := handlers.LoggingHandler(os.Stdout, r)
		addr := config.GetString("app.address")
		logger.Log("msg", "HTTP", "addr", addr)
		logger.Log("err", http.ListenAndServe(addr, handlers.CompressHandler(loggedRouter)))
	}()
	// 等待所有RPC服务完成
	wg.Wait()
}
