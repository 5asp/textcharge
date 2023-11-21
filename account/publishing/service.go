package publishing

import (
	"errors"
	"net/rpc"
	"os/user"

	kitlog "github.com/go-kit/log"
	"github.com/spf13/viper"
)

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidToken = errors.New("invalid token")
)

type Service interface {
	Login(req *loginRequest) (res loginResponse, err error)
	Register(req *registerRequest) (res registerResponse, err error)
}
type service struct {
	client *rpc.Client
}

func New(log kitlog.Logger, config *viper.Viper) Service {
	client, err := rpc.DialHTTP("tcp", config.GetString("account.rpcURL"))
	if err != nil {
		log.Log("rpc error:", err)
	}
	return &service{
		client: client,
	}
}

func (s *service) Login(req *loginRequest) (res loginResponse, err error) {
	return
}
func (s *service) Register(req *registerRequest) (res registerResponse, err error) {
	if req.Account == "" {
		return res, errors.New("account != nil")
	}
	if req.Password == "" {
		return res, errors.New("account != nil")
	}

	var user user.User
	err = s.client.Call("RpcService.FindByAccount", req.Account, &user)
	if err != nil {
		return res, errors.New("rpc account != nil")
	}
	return
}
