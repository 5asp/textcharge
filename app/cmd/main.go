package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aheadIV/textcharge/app/publishing"
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
	publishing.RegisterRPCService(logger, repo, config)

	var service publishing.Service
	{
		service = publishing.New(logger, config)
	}

	r := publishing.MakeHttpHandler(logger, service)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	addr := config.GetString("app.address")

	logger.Log("msg", "HTTP", "addr", addr)
	logger.Log("err", http.ListenAndServe(addr, handlers.CompressHandler(loggedRouter)))
}
