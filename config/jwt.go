package config

import "rackrock/utils"

type Jwt struct {
	Key        string `yaml:"key"`
	ExpireTime uint64 `yaml:"expire-time"`
}

func (jwt *Jwt) Check() {
	utils.IsTrue(utils.IsEmptyStr(jwt.Key), "jwt-key not empty")
	utils.IsTrue(jwt.ExpireTime != 0, "jwt-expire time not empty")
}
