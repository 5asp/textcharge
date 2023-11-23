package publishing

import (
	"context"
	"fmt"
	"net/rpc"

	"github.com/aheadIV/textcharge/application-management-service/app"
	kitlog "github.com/go-kit/log"
	"github.com/go-rel/rel"
	_ "github.com/lib/pq"
)

const (
	Success = 1
	Fail    = 0
)

type RpcAppService struct {
	Repo rel.Repository
	ctx  context.Context
	Log  kitlog.Logger
}

func (s *RpcAppService) Insert(args *app.App, reply *int) error {
	err := s.Repo.Insert(s.ctx, args)
	if err != nil {
		return err
	}
	*reply = args.ID
	return nil
}

func (s *RpcAppService) Update(args *app.App, reply *int) error {
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

func (s *RpcAppService) Delete(ID *string, reply *int) error {
	// if ID != nil {
	// 	err := s.Repo.Delete(s.ctx, &app.App{
	// 		ID: *ID,
	// 	})
	// 	if err != nil {
	// 		*reply = Fail
	// 		return err
	// 	}
	// 	*reply = Success
	// }
	return nil
}

func (s *RpcAppService) FindByID(ID *int, reply *app.App) error {
	if ID != nil {
		id := *ID
		fmt.Println(id)

		err := s.Repo.Find(s.ctx, reply, rel.Eq("id", id))
		if err != nil {
			*reply = app.App{}
			return err
		}
	}
	return nil
}

func RegisterRPCService(log kitlog.Logger, repo rel.Repository) {
	rpcService := &RpcAppService{
		Log:  log,
		Repo: repo,
		ctx:  context.Background(),
	}
	rpc.Register(rpcService)
}
