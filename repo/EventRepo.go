package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
)

func GetTotalEventCountById(db *gorm.DB, userId uint64) (int, error) {
	var count int64

	err := db.Table("event").
		Where("user_id = ?", userId).
		Count(&count).
		Error

	return int(count), err
}

func GetEventIdsByUserId(db *gorm.DB, userId uint64) ([]uint64, error) {
	eventIds := make([]uint64, 0)

	err := db.Table("event").
		Select("id").
		Where("user_id = ?", userId).
		Find(&eventIds).
		Error

	return eventIds, err
}

func GetTotalAmountSoldByEventIds(db *gorm.DB, eventIds []uint64) (int, error) {
	var amount int

	err := db.Table("sales s").
		Joins("left join items i on a.item_id = i.id").
		Select("sum(i.sale_price)").
		Where("i.event_id in ?", eventIds).
		Find(&amount).
		Error

	return amount, err
}

func GetTotalItemSoldByEventIds(db *gorm.DB, eventIds []uint64) (int, error) {
	var itemCount int

	err := db.Table("sales s").
		Joins("left join items i on a.item_id = i.id").
		Select("count(i.sale_price)").
		Where("i.event_id in ?", eventIds).
		Find(&itemCount).
		Error

	return itemCount, err
}

func GetEvents(db *gorm.DB, whereClause, sortOrder string, offset, pageSize int) ([]model.Event, error) {
	var events = make([]model.Event, 0)
	err := db.Table("event").
		Where(whereClause).
		Order(sortOrder).
		Offset(offset).
		Limit(pageSize).
		Error

	return events, err
}

func InsertEvent(db *gorm.DB, event model.Event) (uint64, error) {
	err := db.Create(&event).
		Error

	return event.Id, err
}
