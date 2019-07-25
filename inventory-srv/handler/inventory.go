package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	inventory "github.com/songxuexian/gogomicro/inventory-srv/proto/inventory"
)

type Inventory struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Inventory) Call(ctx context.Context, req *inventory.Request, rsp *inventory.Response) error {
	log.Log("Received Inventory.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Inventory) Stream(ctx context.Context, req *inventory.StreamingRequest, stream inventory.Inventory_StreamStream) error {
	log.Logf("Received Inventory.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&inventory.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Inventory) PingPong(ctx context.Context, stream inventory.Inventory_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&inventory.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
