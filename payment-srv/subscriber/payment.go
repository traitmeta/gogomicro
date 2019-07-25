package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	payment "github.com/songxuexian/gogomicro/payment-srv/proto/payment"
)

type Payment struct{}

func (e *Payment) Handle(ctx context.Context, msg *payment.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *payment.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
