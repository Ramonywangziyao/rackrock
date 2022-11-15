package utils

import (
	"database/sql"
	"fmt"
)

func DBConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:960415@(localhost:3306)/rackrock")
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Connection failed.")
		panic(err.Error())
	}

	fmt.Println("Connected.")

	return db
}
