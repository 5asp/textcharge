package publishing

import (
	"context"
	"net"
	"net/http"
	"net/rpc"

	kitlog "github.com/go-kit/log"

	"github.com/spf13/viper"
)

type RpcService struct {
	ctx context.Context
	Log kitlog.Logger
}

func RegisterRPCService(log kitlog.Logger, config *viper.Viper) {
	rpcService := &RpcService{
		Log: log,
		ctx: context.Background(),
	}
	rpc.Register(rpcService)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", config.GetString("queue.rpc"))
	if e != nil {
		log.Log("listen error:", e)
	}
	go http.Serve(l, nil)
}
