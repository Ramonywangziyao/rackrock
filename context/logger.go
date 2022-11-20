package context

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"rackrock/logger"
	"rackrock/utils"
	"reflect"
)

type LogField struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Method   string `json:"method"`
	Path     string `json:"path"`
}

func (field *LogField) merge(target LogField) {
	var dValue = reflect.ValueOf(target)
	var sValue = reflect.ValueOf(field).Elem()

	var numFields = dValue.NumField()
	for i := 0; i < numFields; i++ {
		var field = dValue.Field(i)
		if utils.IsEmptyStr(field.String()) {
			continue
		}

		sValue.Field(i).Set(field)
	}
}

func LoggerHandle(ctx context.Context) (err error) {

	var logFields = LogField{}

	var login = GetLoginUser(ctx)
	logFields.merge(
		LogField{
			UserId:   fmt.Sprintf("%v", login.ID),
			UserName: login.UserName,
		},
	)

	var ginCtx = ctx.(*gin.Context)
	var req = ginCtx.Request
	logFields.merge(
		LogField{
			Method: req.Method,
			Path:   req.URL.Path,
		},
	)
	var fields = logrus.Fields{}
	utils.MustUnmarshal(utils.MustMarshal(logFields), &fields)

	logger.Logger.WithFields(fields).Info(string(utils.MustMarshal(GetResponse(ctx))))
	return
}
