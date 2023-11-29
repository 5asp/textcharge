package publishing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateAppEndpoint endpoint.Endpoint
	AppInfoEndpoint   endpoint.Endpoint
	RemoveAppEndpoint endpoint.Endpoint
	UpdateAppEndpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateAppEndpoint: makeCreateAppEndpoint(s),
		AppInfoEndpoint:   makeAppInfoEndpoint(s),
		RemoveAppEndpoint: makeRemoveAppEndpoint(s),
		UpdateAppEndpoint: makeUpdateAppEndpoint(s),
	}
}

type createAppRequest struct {
	UserID int `json:"userID"`
}

type createAppResponse struct {
	AppID  int    `json:"appID"`
	Secret string `json:"secret"`
}

func makeCreateAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createAppRequest)
		v, err := s.CreateApp(&req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

type appInfoRequest struct {
	AppID int `json:"appID"`
}

type appInfoResponse struct {
	AppID  int    `json:"appID"`
	Secret string `json:"secret"`
	Status int    `json:"status"`
	ID     int    `json:"id"`
}

func makeAppInfoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(appInfoRequest)
		v, err := s.AppInfo(&req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

type removeAppRequest struct {
	AppID int `json:"appID"`
}

func makeRemoveAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(removeAppRequest)
		v, err := s.RemoveApp(&req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

type updateAppRequest struct {
	AppID  int `json:"appID"`
	Status int `json:"status"`
}

func makeUpdateAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateAppRequest)
		v, err := s.UpdateApp(&req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
