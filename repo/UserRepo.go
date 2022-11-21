package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
	"rackrock/setting"
)

var db = setting.DB
var tableName = "user"

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
		Where("user_id = ?", userId).
		First(&user).
		Error

	return user, err
}

func GetUserAccessLevelByUserId(db *gorm.DB, userId uint64) (int, error) {
	var accessLevel int

	err := db.Table(tableName).
		Select("access_level").
		Where("id = ?", userId).
		First(&accessLevel).
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
