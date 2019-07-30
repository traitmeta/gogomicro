package payment

import (
	"context"
	"github.com/google/uuid"
	"github.com/micro/go-micro/util/log"
	proto "github.com/songxuexian/gogomicro/payment-srv/proto/payment"
	"time"
)

// sendPayDoneEvt 发送支付事件
func (s *service) sendPayDoneEvt(orderId int64, state int32) {
	// 构建事件
	ev := &proto.PayEvent{
		Id:       uuid.New().String(),
		SendTime: time.Now().Unix(),
		OrderId:  orderId,
		State:    state,
	}

	log.Logf("[sendPayDoneEvt] 发送支付事件，%+v\n", ev)

	// 广播
	if err := payPublisher.Publish(context.Background(), ev); err != nil {
		log.Logf("[sendPayDoneEvt] 异常: %v", err)
	}
}
