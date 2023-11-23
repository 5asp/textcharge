package publishing

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/aheadIV/textcharge/app/app"
	kitlog "github.com/go-kit/log"
	"github.com/go-rel/rel"
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

func (s *RpcService) Insert(args *app.App, reply *int) error {
	err := s.Repo.Insert(s.ctx, args)
	if err != nil {
		return err
	}
	*reply = args.ID
	return nil
}
func (s *RpcService) InsertAppUser(args *app.AppUser, reply *int) error {
	err := s.Repo.Insert(s.ctx, args)
	if err != nil {
		return err
	}
	*reply = args.ID
	return nil
}

func (s *RpcService) Update(args *app.App, reply *int) error {
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

func (s *RpcService) Delete(ID *string, reply *int) error {
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

func (s *RpcService) FindByID(ID *int, reply *app.App) error {
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

func RegisterRPCService(log kitlog.Logger, repo rel.Repository, config *viper.Viper) {
	rpcService := &RpcService{
		Log:  log,
		Repo: repo,
		ctx:  context.Background(),
	}
	rpc.Register(rpcService)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", config.GetString("app.rpc"))
	if e != nil {
		log.Log("listen error:", e)
	}
	go http.Serve(l, nil)
}
