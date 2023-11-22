package publishing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func MakeHttpHandler(log kitlog.Logger, s Service) *mux.Router {
	createAppHandler := kithttp.NewServer(
		makeCreateAppEndpoint(s),
		decodeCreateAppRequest,
		encodeResponse,
	)

	removeAppHandler := kithttp.NewServer(
		makeRemoveAppEndpoint(s),
		decodeRemoveAppRequest,
		encodeResponse,
	)

	appInfoHandler := kithttp.NewServer(
		makeAppInfoEndpoint(s),
		decodeAppInfoRequest,
		encodeResponse,
	)

	updateAppHandler := kithttp.NewServer(
		makeUpdateAppEndpoint(s),
		decodeUpdateAppRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/create").Handler(createAppHandler)
	r.Methods("POST").Path("/update").Handler(updateAppHandler)
	r.Methods("GET").Path("/remove/{id:[0-9]+}").Handler(removeAppHandler)
	r.Methods("GET").Path("/info/{id:[0-9]+}").Handler(appInfoHandler)
	return r
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateAppRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req createAppRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("decodeLoginRequest: %s", err)
	}
	return req, nil
}

func decodeUpdateAppRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req updateAppRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("decodeLoginRequest: %s", err)
	}
	return req, nil
}

func decodeRemoveAppRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := strconv.Atoi(vars["id"])
	if id <= 0 || ok != nil {
		return nil, fmt.Errorf("removeAppRequest: %s", err)
	}
	return &removeAppRequest{AppID: id}, nil
}

func decodeAppInfoRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := strconv.Atoi(vars["id"])
	if id <= 0 || ok != nil {
		return nil, fmt.Errorf("appInfoRequest: %s", err)
	}
	return &appInfoRequest{AppID: id}, nil
}
