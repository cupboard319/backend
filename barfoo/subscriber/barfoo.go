package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	barfoo "barfoo/proto/barfoo"
)

type Barfoo struct{}

func (e *Barfoo) Handle(ctx context.Context, msg *barfoo.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *barfoo.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
