package queue_test

import (
	"fmt"
	"testing"

	"github.com/aheadIV/textcharge/queue"
)

func TestNewQueue(t *testing.T) {
	q := queue.NewQueue("ssss")

	if q == nil {
		fmt.Println("connect nats err.")
	}
	fmt.Println(q)
}
