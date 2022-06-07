package headwater

import (
	"context"
	"fmt"
)

type requestHandler[R any, T any] func(context context.Context, request R) (T, error)

type Receiver[R any, T any] struct {
	handler requestHandler[R, T]
}

func (eh *Receiver[R, T]) Send(context context.Context, request R) (T, error) {
	if eh.handler == nil {
		return GetZero[T](), fmt.Errorf("no handler for request: %T", request)
	}
	return eh.handler(context, request)
}

func (eh *Receiver[E, T]) SetHandler(handler requestHandler[E, T]) {
	eh.handler = handler
}
