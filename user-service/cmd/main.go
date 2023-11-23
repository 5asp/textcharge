package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/aheadIV/textcharge/user-service/publishing"
	kitlog "github.com/go-kit/log"
	"github.com/oklog/oklog/pkg/group"

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
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}
