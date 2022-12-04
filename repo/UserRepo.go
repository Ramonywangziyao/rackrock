package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
	"time"
)

var tableName = "user"

func InsertUser(db *gorm.DB, user model.User) error {
	err := db.Table(tableName).
		Create(&user).
		Error

	return err
}

func UpdateLoginTimeByUserId(db *gorm.DB, userId uint64) error {
	time := time.Now()

	err := db.Table(tableName).
		Where("id = ?", userId).
		Update("last_login_time", time).
		Error

	return err
}

func GetUserNickNameById(db *gorm.DB, userId uint64) (string, error) {
	var nickname string

	err := db.Table(tableName).
		Select("nickname").
		Where("user_id = ?", userId).
		Find(&nickname).
		Error

	return nickname, err
}

func GetUserByUserId(db *gorm.DB, userId uint64) (model.User, error) {
	var user model.User

	err := db.Table(tableName).
		Where("id = ?", userId).
		First(&user).
		Error

	return user, err
}

func GetUserAccessLevelByUserId(db *gorm.DB, userId uint64) (int, error) {
	var accessLevel int

	err := db.Table(tableName).
		Select("access_level").
		Where("id = ?", userId).
		Find(&accessLevel).
		Error

	return accessLevel, err
}

func GetUserList(db *gorm.DB) ([]model.User, error) {
	var users = make([]model.User, 0)

	err := db.Table(tableName).
		Find(&users).
		Error

	return users, err
}

func GetUserByAccount(db *gorm.DB, account string) (model.User, error) {
	var user = model.User{}

	err := db.Table(tableName).
		Where("account = ?", account).
		First(&user).
		Error

	return user, err
}
