package publishing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SendEndpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		SendEndpoint: makeSendEndpoint(s),
	}
}

type sendRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type sendResponse struct {
	// Token string `json:"token,omitempty"`
	Err string `json:"err,omitempty"`
}

func makeSendEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// req := request.(sendRequest)
		// v, err := s.Login(ctx, &req)
		// if err != nil {
		// 	return v, nil
		// }
		// return v, nil

		return nil, nil
	}
}
