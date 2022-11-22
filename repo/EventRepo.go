package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
)

var eventTableName = "event"
var salesTableName = "sales"

func GetTotalEventCountById(db *gorm.DB, userId uint64) (int, error) {
	var count int64

	err := db.Table(eventTableName).
		Where("user_id = ?", userId).
		Count(&count).
		Error

	return int(count), err
}

func GetEventIdsByUserId(db *gorm.DB, userId uint64) ([]uint64, error) {
	eventIds := make([]uint64, 0)

	err := db.Table(eventTableName).
		Select("id").
		Where("user_id = ?", userId).
		Find(&eventIds).
		Error

	return eventIds, err
}

func GetTotalAmountSoldByEventIds(db *gorm.DB, eventIds []uint64) (int, error) {
	var amount int

	err := db.Table("sales s").
		Joins("left join items i on s.item_id = i.id").
		Select("sum(i.sale_price)").
		Where("i.event_id in ?", eventIds).
		Find(&amount).
		Error

	return amount, err
}

func GetTotalItemSoldByEventIds(db *gorm.DB, eventIds []uint64) (int, error) {
	var itemCount int

	err := db.Table("sales s").
		Joins("left join items i on s.item_id = i.id").
		Select("count(i.sale_price)").
		Where("i.event_id in ?", eventIds).
		Find(&itemCount).
		Error

	return itemCount, err
}

func GetEventByEventId(db *gorm.DB, eventId uint64) (model.Event, error) {
	var event = model.Event{}

	err := db.Table(eventTableName).
		Where("id = ?", eventId).
		First(&event).
		Error

	return event, err
}

func GetEvents(db *gorm.DB, whereClause, sortOrder string, offset, pageSize int) ([]model.Event, error) {
	var events = make([]model.Event, 0)
	err := db.Table(eventTableName).
		Where(whereClause).
		Order(sortOrder).
		Offset(offset).
		Limit(pageSize).
		Error

	return events, err
}

func GetEventsCountByUserId(db *gorm.DB, userId uint64) (int64, error) {
	var count int64

	err := db.Table(eventTableName).
		Select("count(id)").
		Where("user_id = ?", userId).
		Find(&count).
		Error

	return count, err
}

func InsertEvent(db *gorm.DB, event model.Event) (uint64, error) {
	err := db.Create(&event).
		Error

	return event.Id, err
}
