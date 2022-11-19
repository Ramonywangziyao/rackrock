package model

import "time"

type Member struct {
	Id         int64     `json:"id" gorm:"id" form:"id"`
	Name       string    `json:"name" gorm:"name" form:"name"`
	Nickname   string    `json:"nickname" gorm:"nickname" form:"nickname"`
	Phone      string    `json:"phone" gorm:"phone" form:"phone"`
	Gender     int       `json:"gender" gorm:"gender" form:"gender"`
	Source     int       `json:"source" gorm:"source" form:"source"`
	City       string    `json:"city" gorm:"city" form:"city"`
	CreateTime time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	ModifyTime time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	IsDeleted  int       `json:"is_deleted" gorm:"is_deleted" form:"is_deleted"`
}

type User struct {
	Id          int64     `json:"id" gorm:"id" form:"id"`
	BrandId     int64     `json:"brand_id" gorm:"brand_id" form:"brand_id"`
	Account     string    `json:"account" gorm:"account" form:"account"`
	Nickname    string    `json:"nickname" gorm:"nickname" form:"nickname"`
	Password    string    `json:"password" gorm:"password" form:"password"`
	AccessLevel int       `json:"access_level" gorm:"access_level" form:"access_level"`
	CreateTime  time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	ModifyTime  time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	IsDeleted   int       `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type UserInfo struct {
	Id           string    `json:"id"`
	Brand        BrandInfo `json:"brand"`
	Account      string    `json:"account"`
	Nickname     string    `json:"nickname"`
	CombinedName string    `json:"combined_name"`
	AccessLevel  int       `json:"access_level"`
}

type Brand struct {
	Id          int64     `json:"id" gorm:"id" form:"id"`
	Brand       string    `json:"brand" gorm:"brand" form:"brand"`
	Industry    int       `json:"industry" gorm:"industry" form:"industry"`
	Subindustry int       `json:"subindustry" gorm:"subindustry" form:"subindustry"`
	CreateTime  time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	ModifyTime  time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	IsDeleted   int       `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type BrandInfo struct {
	Id          string `json:"id"`
	Brand       string `json:"brand"`
	Industry    string `json:"industry"`
	Subindustry string `json:"subindustry"`
}

type Tag struct {
	Id         int64     `json:"id" gorm:"id" form:"id"`
	Tag        string    `json:"tag" gorm:"tag" form:"tag"`
	CreateTime time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	ModifyTime time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	IsDeleted  int       `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type TagInfo struct {
	Id  string `json:"id"`
	Tag string `json:"tag"`
}

type Event struct {
	Id            int64     `json:"id" gorm:"id" form:"id"`
	TagId         int64     `json:"tag_id" gorm:"tag_id" form:"tag_id"`
	UserId        int64     `json:"user_id" gorm:"user_id" form:"user_id"`
	EventName     string    `json:"event_name" gorm:"event_name" form:"event_name"`
	City          string    `json:"city" gorm:"city" form:"city"`
	Type          int       `json:"type" gorm:"type" form:"type"`
	LastDays      int       `json:"last_days" gorm:"last_days" form:"last_days"`
	TotalQuantity int       `json:"total_quantity" gorm:"total_quantity" form:"total_quantity"`
	TotalSku      int       `json:"total_sku" gorm:"total_sku" form:"total_sku"`
	ReportStatus  int       `json:"report_status" gorm:"report_status" form:"report_status"`
	StartTime     time.Time `json:"start_time" gorm:"start_time" form:"start_time"`
	EndTime       time.Time `json:"end_time" gorm:"end_time" form:"end_time"`
	CreateTime    time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	ModifyTime    time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	IsDeleted     int       `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type EventInfo struct {
	Id        string  `json:"id"`
	Tag       TagInfo `json:"tag"`
	EventName string  `json:"event_name"`
	City      string  `json:"city"`
	Type      string  `json:"type"`
	StartTime string  `json:"start_time"`
	EndTime   string  `json:"end_time"`
}

type EventItem struct {
	Id          int64     `json:"id" gorm:"id" form:"id"`
	EventId     int64     `json:"event_id" gorm:"event_id" form:"event_id"`
	Brand       string    `json:"brand" gorm:"brand" form:"brand"`
	Sku         string    `json:"sku" gorm:"sku" form:"sku"`
	Barcode     string    `json:"barcode" gorm:"barcode" form:"barcode"`
	RetailPrice int       `json:"retail_price" gorm:"retail_price" form:"retail_price"`
	SalePrice   int       `json:"sale_price" gorm:"sale_price" form:"sale_price"`
	Discount    float32   `json:"discount" gorm:"discount" form:"discount"`
	Season      string    `json:"season" gorm:"season" form:"season"`
	Category    string    `json:"category" gorm:"category" form:"category"`
	Color       string    `json:"color" gorm:"color" form:"color"`
	Size        string    `json:"size" gorm:"size" form:"size"`
	Inventory   int       `json:"inventory" gorm:"inventory" form:"inventory"`
	CreateTime  time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	ModifyTime  time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	IsDeleted   int       `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type SaleRecord struct {
	Id         int64     `json:"id" gorm:"id" form:"id"`
	MemberId   int64     `json:"member_id" gorm:"member_id" form:"member_id"`
	ItemId     int64     `json:"item_id" gorm:"item_id" form:"item_id"`
	OrderId    string    `json:"order_id" gorm:"order_id" form:"order_id"`
	OrderTime  time.Time `json:"order_time" gorm:"order_time" form:"order_time"`
	CouponUsed int       `json:"coupon_used" gorm:"coupon_used" form:"coupon_used"`
	Source     int       `json:"source" gorm:"source" form:"source"`
	IsReturn   int       `json:"is_return" gorm:"is_return" form:"is_return"`
	CreateTime time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	ModifyTime time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	IsDeleted  int       `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type CoreMetric struct {
	ItemSold   int     `json:"item_sold"`
	OrderSold  int     `json:"order_sold"`
	AmountSold float32 `json:"amount_sold"`
	Conversion float32 `json:"conversion"`
}

type SecondaryMetric struct {
	ReturnAmount    float32 `json:"return_amount"`
	AverageSku      float32 `json:"average_sku"`
	AverageItem     float32 `json:"average_item"`
	AverageAmount   float32 `json:"average_amount"`
	AveragePrice    float32 `json:"average_price"`
	AverageDiscount float32 `json:"average_discount"`
	MaxDiscount     float32 `json:"max_discount"`
	MinDiscount     float32 `json:"min_discount"`
}

type Distribution struct {
	PriceDistribution    []DistributionItem `json:"price_distribution"`
	DiscountDistribution []DistributionItem `json:"discount_distribution"`
}

type DistributionItem struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Rank struct {
	Rank int    `json:"rank"`
	Item string `json:"item"`
}

type DailyDetail struct {
	Date string `json:"date"`
	CoreMetric
	ReturnAmount float32 `json:"return_amount"`
	Growth       float32 `json:"growth_to_yesterday"`
}
