package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
)

var tagTableName = "tag"

func InsertTag(db *gorm.DB, tag model.Tag) (uint64, error) {
	err := db.Create(&tag).
		Error

	return tag.Id, err
}

func GetTagById(db *gorm.DB, id uint64) (model.Tag, error) {
	var tag model.Tag

	err := db.Table(tagTableName).
		Where("id = ?", id).
		First(&tag).
		Error

	return tag, err
}

func GetTagsByUserId(db *gorm.DB, userId uint64) ([]model.Tag, error) {
	var tags = make([]model.Tag, 0)

	err := db.Table(tagTableName).
		Where("user_id = ?", userId).
		Find(&tags).
		Error

	return tags, err
}
