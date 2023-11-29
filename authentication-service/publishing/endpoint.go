package publishing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	LoginEndpoint    endpoint.Endpoint
	RegisterEndpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		LoginEndpoint:    makeLoginEndpoint(s),
		RegisterEndpoint: makeRegisterEndpoint(s),
	}
}

type loginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"err,omitempty"`
}

func makeLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		v, err := s.Login(ctx, &req)
		if err != nil {
			return v, nil
		}
		return v, nil
	}
}

type registerRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type registerResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err,omitempty"`
}

func makeRegisterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerRequest)
		v, err := s.Register(ctx, &req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
