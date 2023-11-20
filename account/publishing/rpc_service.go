package publishing

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/aheadIV/textcharge/account/user"
	kitlog "github.com/go-kit/log"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	Success = 1
	Fail    = 0
)

type RpcService struct {
	Repo rel.Repository
	ctx  context.Context
	Log  kitlog.Logger
}

func (s *RpcService) Insert(args *user.User, reply *int) error {
	err := s.Repo.Insert(s.ctx, args)
	if err != nil {
		return err
	}
	*reply = Success
	return nil
}

func (s *RpcService) Update(args *user.User, reply *int) error {
	if args.ID != 0 {
		err := s.Repo.Update(s.ctx, args)
		if err != nil {
			*reply = Fail
			return err
		}
		*reply = Success
	}
	return nil
}

func (s *RpcService) Delete(userID *int, reply *int) error {
	if userID != nil {
		err := s.Repo.Delete(s.ctx, &user.User{
			ID: *userID,
		})
		if err != nil {
			*reply = Fail
			return err
		}
		*reply = Success
	}
	return nil
}

func (s *RpcService) FindByID(userID *int, reply *user.User) error {
	if userID != nil {
		err := s.Repo.Find(s.ctx, reply, rel.Eq("id", userID))
		if err != nil {
			*reply = user.User{}
			return err
		}
	}
	return nil
}

func (s *RpcService) FindByAccount(account string, reply *user.User) error {
	if account != "" {
		if err := s.Repo.Find(s.ctx, reply, where.Eq("account", account)); err != nil {
			if errors.Is(err, rel.ErrNotFound) {
				fmt.Println(2323232)
				return nil
			}
			return err
		}
	}
	return nil
}

func RegisterRPCService(log kitlog.Logger, repo rel.Repository, config *viper.Viper) {
	rpcService := &RpcService{
		Log:  log,
		Repo: repo,
		ctx:  context.Background(),
	}
	rpc.Register(rpcService)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", config.GetString("account.rpc"))
	if e != nil {
		log.Log("listen error:", e)
	}
	go http.Serve(l, nil)
}
