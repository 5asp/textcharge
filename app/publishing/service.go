package publishing

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/rpc"

	"github.com/aheadIV/textcharge/account/user"
	"github.com/aheadIV/textcharge/app/app"
	kitlog "github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

//	r.Methods("POST").Path("/create").Handler(createAppHandler)
//
// r.Methods("POST").Path("/remove").Handler(removeAppHandler)
// r.Methods("POST").Path("/info").Handler(appInfoHandler)
// r.Methods("POST").Path("/update").Handler(updateAppHandler)
type Service interface {
	CreateApp(req *createAppRequest) (res *createAppResponse, err error)
	AppInfo(req *appInfoRequest) (res *appInfoResponse, err error)
	UpdateApp(req *updateAppRequest) (result *int, err error)
	RemoveApp(req *removeAppRequest) (result *int, err error)
}
type service struct {
	appClient     *rpc.Client
	accountClient *rpc.Client
}

func New(log kitlog.Logger, config *viper.Viper) Service {
	appClient, err := rpc.DialHTTP("tcp", config.GetString("app.rpcURL"))

	if err != nil {
		log.Log("app rpc error:", err)
	}
	accountClient, err := rpc.DialHTTP("tcp", config.GetString("account.rpcURL"))

	if err != nil {
		log.Log("account rpc error:", err)
	}
	return &service{
		appClient:     appClient,
		accountClient: accountClient,
	}
}

func (s *service) CreateApp(req *createAppRequest) (res *createAppResponse, err error) {
	if req.UserID == 0 {
		err = errors.New("missing userID")
		return nil, err
	}
	accountData := &user.User{}
	err = s.accountClient.Call("RpcService.FindByID", req.UserID, accountData)
	if err != nil {
		err = errors.New("callFindAccountByID err")
		return nil, err
	}

	if accountData.Status == 0 {
		err = errors.New("account disabled")
		return nil, err
	}
	uuid := uuid.New()
	key := uuid.String()
	h := md5.New()

	io.WriteString(h, key)
	appData := &app.App{
		Secret: hex.EncodeToString(h.Sum(nil)),
		Status: 1,
	}
	var appID int
	err = s.appClient.Call("RpcService.Insert", appData, &appID)
	fmt.Println(err)
	if err != nil && appID == 0 {
		err = errors.New("call create app err")
		return nil, err
	}
	appUserData := &app.AppUser{
		AppID:  appID,
		UserID: accountData.ID,
	}
	var result int
	err = s.appClient.Call("RpcService.InsertAppUser", appUserData, &result)
	if err != nil {
		err = errors.New("call create app user err")
		return nil, err
	}

	return &createAppResponse{
		AppID:  appID,
		Secret: appData.Secret,
	}, nil
}

func (s *service) AppInfo(req *appInfoRequest) (res *appInfoResponse, err error) {

	return
}
func (s *service) UpdateApp(req *updateAppRequest) (result *int, err error) {

	return
}
func (s *service) RemoveApp(req *removeAppRequest) (result *int, err error) {

	return
}
