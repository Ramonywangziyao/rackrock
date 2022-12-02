package context

import (
	"github.com/gin-gonic/gin"
)

type HandleFunction func(ctx *gin.Context) error

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

func OperateHandler(ctx *gin.Context, handlers []HandleFunction) error {
	var err error
	var handle HandleFunction
	for _, handle = range handlers {
		if err = handle(ctx); err != nil {
			return err
		}
	}

	return nil
}
