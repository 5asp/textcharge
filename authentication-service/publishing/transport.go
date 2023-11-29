package publishing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func MakeHttpHandler(log kitlog.Logger, s Service) *mux.Router {
	loginHandler := kithttp.NewServer(
		makeLoginEndpoint(s),
		decodeLoginRequest,
		encodeResponse,
	)

	registerHandler := kithttp.NewServer(
		makeRegisterEndpoint(s),
		decodeRegisterRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/login").Handler(loginHandler)
	r.Methods("POST").Path("/register").Handler(registerHandler)
	return r
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeLoginRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("decodeLoginRequest: %s", err)
	}
	return req, nil
}

func decodeRegisterRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("decodeRegisterRequest %s", err)
	}
	return req, nil
}
