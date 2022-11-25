package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
)

func GetMemberIdByMemberDetail(db *gorm.DB, whereClause string) (uint64, error) {
	var memberId uint64
	err := db.Table("membership").
		Where(whereClause).
		First(&memberId).
		Error

	return memberId, err
}

func BatchInsertMembers(db *gorm.DB, members []model.Member) error {
	err := db.Create(&members).
		Error

	return err
}
