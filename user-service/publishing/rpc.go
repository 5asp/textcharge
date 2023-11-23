package publishing

import (
	"context"
	"errors"
	"net/rpc"

	"github.com/aheadIV/textcharge/account/user"
	kitlog "github.com/go-kit/log"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	_ "github.com/lib/pq"
)

const (
	Success = 1
	Fail    = 0
)

type RpcUserService struct {
	Repo rel.Repository
	ctx  context.Context
	Log  kitlog.Logger
}

func (s *RpcUserService) Insert(args *user.User, reply *int) error {
	err := s.Repo.Insert(s.ctx, args)
	if err != nil {
		return err
	}
	*reply = args.ID
	return nil
}

func (s *RpcUserService) Update(args *user.User, reply *int) error {
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

func (s *RpcUserService) Delete(userID *int, reply *int) error {
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

func (s *RpcUserService) FindByID(userID *int, reply *user.User) error {
	if userID != nil {
		err := s.Repo.Find(s.ctx, reply, rel.Eq("id", userID))
		if errors.Is(err, rel.ErrNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func (s *RpcUserService) CreateUserApp(args *user.UserApp, reply *int) error {
	err := s.Repo.Insert(s.ctx, args)
	if err != nil {
		return err
	}
	*reply = args.ID
	return nil
}

func (s *RpcUserService) FindByAccount(account string, reply *user.User) error {
	if account != "" {
		if err := s.Repo.Find(s.ctx, reply, where.Eq("account", account)); err != nil {
			if errors.Is(err, rel.ErrNotFound) {
				return nil
			}
			return err
		}
	}
	return nil
}

func RegisterRPCService(log kitlog.Logger, repo rel.Repository) {
	rpcService := &RpcUserService{
		Log:  log,
		Repo: repo,
		ctx:  context.Background(),
	}
	rpc.Register(rpcService)
}
