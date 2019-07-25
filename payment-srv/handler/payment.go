package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	payment "github.com/songxuexian/gogomicro/payment-srv/proto/payment"
)

type Payment struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Payment) Call(ctx context.Context, req *payment.Request, rsp *payment.Response) error {
	log.Log("Received Payment.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Payment) Stream(ctx context.Context, req *payment.StreamingRequest, stream payment.Payment_StreamStream) error {
	log.Logf("Received Payment.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&payment.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Payment) PingPong(ctx context.Context, stream payment.Payment_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&payment.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
