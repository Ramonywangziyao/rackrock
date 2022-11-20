package config

import (
	"path"
	"rackrock/utils"
	"strings"
)

type Log struct {
	Level   string   `yaml:"level"`
	LogFile *LogFile `yaml:"logFile"`
}

type LogFile struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func (log *Log) Check() {
	utils.IsTrue(utils.InterfaceContains(strings.ToLower(log.Level),
		[]string{"info", "debug", "warn", "error"}), "log level is not valid")
}

func (file *LogFile) GetFileName() string {
	var fpath, fname string

	if fp := file.Path; !utils.IsEmptyStr(fp) {
		fpath = fp
	} else {
		fpath = "./"
	}

	if fn := file.Name; !utils.IsEmptyStr(fn) {
		fname = fn
	} else {
		fname = "rock.log"
	}
	return path.Join(fpath, fname)
}
