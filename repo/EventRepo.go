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
		Where("user_id = ? or creator_id = ?", userId, userId).
		Count(&count).
		Error

	return int(count), err
}

func GetEventsByUserId(db *gorm.DB, userId uint64) ([]model.Event, error) {
	events := make([]model.Event, 0)

	err := db.Table(eventTableName).
		Where("user_id = ?", userId).
		Find(&events).
		Error

	return events, err
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

	err := db.Debug().
		Table(eventTableName).
		Where(whereClause).
		Order(sortOrder).
		Offset(offset).
		Limit(pageSize).
		Find(&events).
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

func GetItemByItemDetail(db *gorm.DB, whereClause string) (model.EventItem, error) {
	var item model.EventItem
	err := db.Table("items").
		Where(whereClause).
		First(&item).
		Error
	return item, err
}

func GetSaleRecordsByOrderId(db *gorm.DB, orderId, paidPrice string) ([]model.SaleRecord, error) {
	var saleRecords = make([]model.SaleRecord, 0)
	err := db.Table("sales").
		Where("order_id = ? and paid_price = ?", orderId, paidPrice).
		Find(&saleRecords).
		Error

	return saleRecords, err
}

func InsertEvent(db *gorm.DB, event model.Event) (uint64, error) {
	err := db.Table("event").
		Create(&event).
		Error

	return event.Id, err
}

func UpdateEventReportStatusByEventId(db *gorm.DB, eventId uint64) error {
	err := db.Table("event").
		Where("id = ?", eventId).
		Update("report_status", 1).
		Error
	return err
}

func UpdateReturnStatus(db *gorm.DB, ids []uint64) error {
	err := db.Table("sales").
		Where("id in (?)", ids).
		Update("is_return", 1).
		Error
	return err
}

func BatchInsertEventItems(db *gorm.DB, items []model.EventItem) error {
	err := db.Table("items").
		Create(&items).
		Error

	return err
}

func BatchInsertEventSales(db *gorm.DB, sales []model.SaleRecord) error {
	err := db.Table("sales").
		Create(&sales).
		Error

	return err
}
