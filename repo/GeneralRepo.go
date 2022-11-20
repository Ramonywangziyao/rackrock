package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
)

func InsertTag(db *gorm.DB, tag model.Tag) (int64, error) {
	err := db.Create(&tag).
		Error

	return tag.Id, err
}
