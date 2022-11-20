package config

import (
	"rackrock/utils"
)

type Aes struct {
	Key string `yaml:"key"`
}

func (aes *Aes) Check() {
	var length int = len(aes.Key)

	utils.IsTrue(utils.InterfaceContains(length, []int{16, 24, 32}),
		"aes.key len should be 16、24、32, current len %d", length)
}
