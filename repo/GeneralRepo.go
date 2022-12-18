package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
)

var tagTableName = "tag"

func InsertTag(db *gorm.DB, tag model.Tag) (uint64, error) {
	err := db.Table("tag").
		Create(&tag).
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

func GetAllTags(db *gorm.DB) ([]model.Tag, error) {
	var tags = make([]model.Tag, 0)

	err := db.Table(tagTableName).
		Find(&tags).
		Error

	return tags, err
}

func GetTagIdsByTag(db *gorm.DB, tag string, userId uint64) ([]uint64, error) {
	var tagIds = make([]uint64, 0)

	err := db.Table(tagTableName).
		Select("id").
		Where("tag = ? and user_id = ?", tag, userId).
		Find(&tagIds).
		Error

	return tagIds, err
}
