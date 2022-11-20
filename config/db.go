package config

import (
	"fmt"
	"rackrock/utils"
)

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"db-name"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	//LogMode     bool   `yaml:"logMode"`
	MaxIdleConn int `yaml:"max-idle-conn"`
	MaxOpenConn int `yaml:"max-open-conn"`
}

func (db *DB) Dsn() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true", db.UserName, db.Password, db.Host, db.Port, db.DbName)
}

func (db *DB) Check() {
	utils.IsTrue(!utils.IsEmptyStr(db.Host) && db.Port != 0 && !utils.IsEmptyStr(db.DbName) && !utils.IsEmptyStr(db.UserName), "db param empty")
}
