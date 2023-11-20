package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aheadIV/textcharge/auth/publishing"
	kitlog "github.com/go-kit/log"
	"github.com/julienschmidt/httprouter"

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
	var logger kitlog.Logger
	{
		logger = kitlog.NewJSONLogger(os.Stderr)
		logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	}
	var router *httprouter.Router
	{
		router = httprouter.New()
	}

	var service publishing.Service
	{
		service = publishing.New(logger, config)
	}

	publishing.MakeHttpHandler(router, service)

	logger.Log("msg", "HTTP", "addr", config.GetString("auth.address"))
	logger.Log("err", http.ListenAndServe(config.GetString("auth.address"), router))
}
