package queue

import "github.com/nats-io/nats.go"

func NewQueue(url string) *nats.Conn {
	nc, _ := nats.Connect(nats.DefaultURL)
	return nc
}
