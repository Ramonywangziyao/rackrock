package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"rackrock/utils"
)

type RockConfig struct {
	Jwt  *Jwt   `yaml:"jwt"`
	Aes  *Aes   `yaml:"aes"`
	Db   *DB    `yaml:"db"`
	Log  *Log   `yaml:"log"`
	Port string `yaml:"port"`
}

type Option func(cfg *RockConfig)

func CheckJwt() Option {
	return func(cfg *RockConfig) {
		utils.IsTrue(cfg.Jwt != nil, "conf jwt should not empty")
		cfg.Jwt.Check()
	}
}

func CheckAes() Option {
	return func(cfg *RockConfig) {
		utils.IsTrue(cfg.Aes != nil, "conf aes should not empty")
		cfg.Aes.Check()
	}
}

func CheckDB() Option {
	return func(cfg *RockConfig) {
		utils.IsTrue(cfg.Db != nil, "conf db should not empty")
		cfg.Db.Check()
	}
}

func CheckLog() Option {
	return func(cfg *RockConfig) {
		utils.IsTrue(cfg.Log != nil, "conf log should not empty")
		cfg.Log.Check()
	}
}

func (cfg *RockConfig) Check() {
	for _, option := range []Option{CheckJwt(), CheckAes(), CheckDB(), CheckLog()} {
		option(cfg)
	}
}

var Cfg RockConfig

// Init resolve config
func Init() {
	var cfgPath = "rock.yml"

	var fullPath, _ = filepath.Abs(cfgPath)
	yamlLoad(fullPath, &Cfg)

}

func yamlLoad(path string, store interface{}) {
	var bytes, err = os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(bytes, store); err != nil {
		panic(err)
	}
}
