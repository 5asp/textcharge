package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aheadIV/textcharge/sending/publishing"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/handlers"

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

	var service publishing.Service
	{
		service = publishing.New(logger, config)
	}

	r := publishing.MakeHttpHandler(logger, service)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	addr := config.GetString("send.address")

	logger.Log("msg", "HTTP", "addr", addr)
	logger.Log("err", http.ListenAndServe(addr, handlers.CompressHandler(loggedRouter)))
}
