package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
	"rackrock/setting"
)

var db = setting.DB

func GetUserNickNameById(db *gorm.DB, userId int64) (string, error) {
	var nickname string

	err := db.Table("user").
		Select("nickname").
		Where("user_id = ?", userId).
		Find(&nickname).
		Error

	return nickname, err
}

func GetUserByUserId(db *gorm.DB, userId int64) (model.User, error) {
	var user model.User

	err := db.Table("user").
		Where("user_id = ?", userId).
		First(&user).
		Error

	return user, err
}