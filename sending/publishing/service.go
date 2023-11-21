package publishing

import (
	"errors"
	"net/rpc"

	kitlog "github.com/go-kit/log"

	"github.com/spf13/viper"
)

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidToken = errors.New("invalid token")
)

type Service interface {
	// Send(ctx context.Context, req *loginRequest) (res loginResponse, err error)
	// Register(ctx context.Context, req *registerRequest) (res registerResponse, err error)
}
type service struct {
	client *rpc.Client
	log    kitlog.Logger
	config *viper.Viper
}

func New(log kitlog.Logger, config *viper.Viper) Service {
	client, err := rpc.DialHTTP("tcp", config.GetString("queue.rpcURL"))
	if err != nil {
		log.Log("rpc error:", err)
	}
	return &service{
		client: client,
		log:    log,
		config: config,
	}
}
