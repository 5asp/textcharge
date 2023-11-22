package publishing

import (
	"context"
	"errors"
	"fmt"
	"net/rpc"

	"github.com/aheadIV/textcharge/account/user"
	"github.com/aheadIV/textcharge/auth/security"

	kitlog "github.com/go-kit/log"
	"github.com/spf13/viper"
)

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidToken = errors.New("invalid token")
)

type Service interface {
	Login(ctx context.Context, req *loginRequest) (res loginResponse, err error)
	Register(ctx context.Context, req *registerRequest) (res registerResponse, err error)
}
type service struct {
	client *rpc.Client
	log    kitlog.Logger
	config *viper.Viper
}

func New(log kitlog.Logger, config *viper.Viper) Service {
	client, err := rpc.DialHTTP("tcp", config.GetString("account.rpcURL"))
	if err != nil {
		log.Log("rpc error:", err)
	}
	return &service{
		client: client,
		log:    log,
		config: config,
	}
}

func (s *service) Login(ctx context.Context, req *loginRequest) (res loginResponse, err error) {
	if req.Account == "" {
		return res, errors.New("account != nil")
	}

	account := CheckAccountExist(s, req.Account)
	if account == nil {
		res.Err = "account not exits."
		return res, errors.New("account not exits.")
	}

	if req.Password == "" {
		return res, errors.New("password != nil")
	}
	if !security.CheckPassword(account.Password, s.config.GetString("auth.salt")+req.Password+s.config.GetString("auth.salt")) {
		return res, errors.New("password not match.")
	}

	token, err := security.NewToken(req.Account)
	if err != nil {
		res.Err = "create token faild."
		return res, err
	}
	res.Token = token
	return res, nil
}

const (
	Success = 1
	Fail    = 0
)

func (s *service) Register(ctx context.Context, req *registerRequest) (res registerResponse, err error) {
	if req.Account == "" {
		return res, errors.New("account != nil")
	}

	account := CheckAccountExist(s, req.Account)

	fmt.Println(account)
	if account.ID > 0 {
		res.Err = "account exits."
		return res, errors.New("account exits.")
	}

	if req.Password == "" {
		return res, errors.New("password != nil")
	}

	register := &user.User{
		Account:  req.Account,
		Password: security.CreatePassword(s.config.GetString("salt") + req.Password + s.config.GetString("salt")),
	}
	var status int
	err = s.client.Call("RpcService.Insert", register, &status)
	fmt.Println("status:", status)
	if err != nil && status == Fail {
		s.log.Log(err)
		return res, nil
	}
	res.Code = Success
	return res, nil
}

func CheckAccountExist(s *service, account string) (result *user.User) {
	result = &user.User{}
	if account != "" {
		err := s.client.Call("RpcService.FindByAccount", account, result)
		if err != nil {
			s.log.Log(err)
			return
		}
	}
	return
}
