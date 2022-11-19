package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	username := "root"   //账号
	password := "960415" //密码
	host := "localhost"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "rackrock" //数据库名

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	fmt.Println("Connected.")

	return db
}
