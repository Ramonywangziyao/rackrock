package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
)

func GetSoldItemDetailByEventId(db *gorm.DB, whereClause string) ([]model.SaleRecordDetail, error) {
	var soldItemDetail = make([]model.SaleRecordDetail, 0)

	err := db.Table("sales s").
		Joins("left join items i on s.item_id = i.id").
		Joins("left join membership m on s.member_id = m.id").
		Select("s.id as id, m.id as member_id, m.name, m.nickname, m.phone, m.gender, " +
			"m.source as member_source, m.city, s.order_id, s.order_time, s.coupon_used, " +
			"s.source as sale_source, s.is_return, s.item_id, i.event_id, i.brand, i.sku, " +
			"i.barcode, i.retail_price, i.sale_price, i.discount, i.season, i.category, " +
			"i.color, i.size, i.inventory").
		Where(whereClause).
		Order("s.order_id desc").
		Find(&soldItemDetail).
		Error

	return soldItemDetail, err
}

func GetSoldItemDetailByEventIdWithOrder(db *gorm.DB, whereClause, sortBy string) ([]model.SaleRecordDetail, error) {
	var soldItemDetail = make([]model.SaleRecordDetail, 0)

	err := db.Table("sales s").
		Joins("left join items i on s.item_id = i.id").
		Joins("left join membership m on s.member_id = m.id").
		Select("s.id as id, m.id as member_id, m.name, m.nickname, m.phone, m.gender, " +
			"m.source as member_source, m.city, s.order_id, s.order_time, s.coupon_used, " +
			"s.source as sale_source, s.is_return, s.item_id, i.event_id, i.brand, i.sku, " +
			"i.barcode, i.retail_price, i.sale_price, i.discount, i.season, i.category, " +
			"i.color, i.size, i.inventory").
		Where(whereClause).
		Order(sortBy).
		Find(&soldItemDetail).
		Error

	return soldItemDetail, err
}

func GetTotalItemCountByEventId(db *gorm.DB, eventId uint64) (int64, error) {
	var totalItemCount int64

	err := db.Table("items").
		Select("sum(inventory)").
		Where("event_id = ?", eventId).
		Find(&totalItemCount).
		Error

	return totalItemCount, err
}

func GetRankItems(db *gorm.DB, selects, whereClause, groupBy, sortBy string, offset, pageSize int) ([]model.RankRecord, error) {
	var rankRecords = make([]model.RankRecord, 0)

	err := db.Table("sales s").
		Joins("left join items i on s.item_id = i.id").
		Select(selects).
		Where(whereClause).
		Group(groupBy).
		Order(sortBy).
		Offset(offset).
		Limit(pageSize).
		Find(&rankRecords).
		Error

	return rankRecords, err
}

func GetRankTotalCount(db *gorm.DB, selects, whereClause, groupBy, sortBy string) ([]model.RankRecord, error) {
	var rankRecords = make([]model.RankRecord, 0)

	err := db.Table("sales s").
		Joins("left join items i on s.item_id = i.id").
		Select(selects).
		Where(whereClause).
		Group(groupBy).
		Order(sortBy).
		Find(&rankRecords).
		Error

	return rankRecords, err
}
