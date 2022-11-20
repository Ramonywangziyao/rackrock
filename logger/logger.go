package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"rackrock/config"
	"rackrock/utils"
	"strings"
	"time"
)

var Logger = logrus.New()

func Init() {
	var logCfg = config.Cfg.Log

	Logger.SetFormatter(new(formatter))
	level, _ := logrus.ParseLevel(logCfg.Level)
	Logger.SetLevel(level)

	if file := logCfg.LogFile; file != nil {
		// write
		var ofile, err = os.OpenFile(file.GetFileName(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(fmt.Sprintf("create log file err: %s", err.Error()))
		}

		Logger.Out = ofile
	}
}

type formatter struct {
}

// Format %date{yyyy-MM-dd HH:mm:ss.SSS}|%t|%-5level|%X{traceId}|%C{0}#%M:%L|%msg%n
func (log *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var timestamp = time.Now().Local().Format(utils.DateFormat)

	var msg = ""
	for k, v := range entry.Data {
		msg = msg + fmt.Sprintf("|%s: %s|", k, v)
	}

	var logMsg = fmt.Sprintf("%s|%s|%s|%s\n", timestamp, strings.ToLower(entry.Level.String()), msg, entry.Message)
	return []byte(logMsg), nil
}
