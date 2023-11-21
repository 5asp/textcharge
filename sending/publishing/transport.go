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
	sendHandler := kithttp.NewServer(
		makeSendEndpoint(s),
		decodeSendRequest,
		encodeSendResponse,
	)

	router.Handler(http.MethodPost, "/send", sendHandler)
}

func decodeSendRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req sendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("Login err: %s", err)
	}
	return req, nil
}

func encodeSendResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(sendResponse)
	if !ok {
		return fmt.Errorf("Login failed: %s", "1")
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}
