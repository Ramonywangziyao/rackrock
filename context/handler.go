package context

import (
	"context"
)

type HandleFunction func(ctx context.Context) error

var (
	BeforeHandler []HandleFunction
	AfterHandler  []HandleFunction
)

func AddBeforeHandler(handle HandleFunction) {
	BeforeHandler = append(BeforeHandler, handle)
}

func AddAfterHandler(handle HandleFunction) {
	AfterHandler = append(AfterHandler, handle)
}

func OperateHandler(ctx context.Context, handlers []HandleFunction) error {
	for _, handle := range handlers {
		if err := handle(ctx); err != nil {
			return err
		}
	}

	return nil
}
