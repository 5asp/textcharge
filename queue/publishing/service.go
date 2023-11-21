package publishing

import "net/rpc"

type Service interface {
	ReviewList()
	ReviewAccpct()
	// Login(req *loginRequest) (res loginResponse, err error)
	// Register(req *registerRequest) (res registerResponse, err error)
}
type service struct {
	client *rpc.Client
}

func (s *service) ReviewList() {

}

func (s *service) ReviewAccpct() {

}
