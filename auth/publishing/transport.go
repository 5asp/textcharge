package publishing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

type Router interface {
	Handle(method, path string, handler http.Handler)
}

func MakeHttpHandler(router *httprouter.Router, s Service) {
	loginHandler := kithttp.NewServer(
		makeLoginEndpoint(s),
		decodeLoginRequest,
		encodeLoginResponse,
	)

	registerHandler := kithttp.NewServer(
		makeRegisterEndpoint(s),
		decodeRegisterRequest,
		encodeRegisterResponse,
	)

	router.Handler(http.MethodPost, "/login", loginHandler)
	router.Handler(http.MethodPost, "/register", registerHandler)
}

func decodeLoginRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("Login err: %s", err)
	}
	return req, nil
}

func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(loginResponse)
	if !ok {
		return fmt.Errorf("Login failed: %s", "1")
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}

func decodeRegisterRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("decodeRegisterRequest %s", err)
	}
	return req, nil
}

func encodeRegisterResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(registerResponse)
	if !ok {
		return fmt.Errorf("encodeRegisterResponse failed cast response")
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}
