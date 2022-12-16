package model

import "time"

type Member struct {
	Id       uint64    `json:"id" gorm:"id" form:"id"`
	Name     string    `json:"name" gorm:"name" form:"name"`
	Nickname string    `json:"nickname" gorm:"nickname" form:"nickname"`
	Phone    string    `json:"phone" gorm:"phone" form:"phone"`
	Gender   int       `json:"gender" gorm:"gender" form:"gender"`
	Source   int       `json:"source" gorm:"source" form:"source"`
	City     string    `json:"city" gorm:"city" form:"city"`
	Dob      time.Time `json:"dob" gorm:"dob" form:"dob"`
	//CreateTime time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	//ModifyTime time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	//IsDeleted  uint      `json:"is_deleted" gorm:"is_deleted" form:"is_deleted"`
}

type User struct {
	Id            uint64    `json:"id" gorm:"id" form:"id"`
	BrandId       uint64    `json:"brand_id" gorm:"brand_id" form:"brand_id"`
	Account       string    `json:"account" gorm:"account" form:"account"`
	Nickname      string    `json:"nickname" gorm:"nickname" form:"nickname"`
	Password      string    `json:"password" gorm:"password" form:"password"`
	AccessLevel   int       `json:"access_level" gorm:"access_level" form:"access_level"`
	LastLoginTime time.Time `json:"last_login_time" gorm:"last_login_time" form:"last_login_time"`
	//CreateTime    time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	//ModifyTime    time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	//IsDeleted     uint      `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
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
	Id              uint64 `json:"id" gorm:"id" form:"id"`
	Brand           string `json:"brand" gorm:"brand" form:"brand"`
	IndustryCode    int    `json:"industry_code" gorm:"industry_code" form:"industry_code"`
	SubindustryCode int    `json:"subindustry_code" gorm:"subindustry_code" form:"subindustry_code"`
	//CreateTime      time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	//ModifyTime      time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	//IsDeleted       uint      `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type BrandInfo struct {
	Id          string `json:"id"`
	Brand       string `json:"brand"`
	Industry    string `json:"industry"`
	Subindustry string `json:"subindustry"`
}

type Tag struct {
	Id     uint64 `json:"id" gorm:"id" form:"id"`
	Tag    string `json:"tag" gorm:"tag" form:"tag"`
	UserId uint64 `json:"user_id" gorm:"user_id" form:"user_id"`
	//CreateTime time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	//ModifyTime time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	//IsDeleted  uint      `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type TagInfo struct {
	Id  string `json:"id"`
	Tag string `json:"tag"`
}

type Event struct {
	Id            uint64    `json:"id" gorm:"id" form:"id"`
	TagId         uint64    `json:"tag_id" gorm:"tag_id" form:"tag_id"`
	UserId        uint64    `json:"user_id" gorm:"user_id" form:"user_id"`
	EventName     string    `json:"event_name" gorm:"event_name" form:"event_name"`
	City          string    `json:"city" gorm:"city" form:"city"`
	Type          int       `json:"type" gorm:"type" form:"type"`
	LastDays      int       `json:"last_days" gorm:"last_days" form:"last_days"`
	TotalQuantity int       `json:"total_quantity" gorm:"total_quantity" form:"total_quantity"`
	TotalSku      int       `json:"total_sku" gorm:"total_sku" form:"total_sku"`
	ReportStatus  uint      `json:"report_status" gorm:"report_status" form:"report_status"`
	CreatorId     uint64    `json:"creator_id" gorm:"creator_id" form:"creator_id"`
	StartTime     time.Time `json:"start_time" gorm:"start_time" form:"start_time"`
	EndTime       time.Time `json:"end_time" gorm:"end_time" form:"end_time"`
	//CreateTime    time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	//ModifyTime    time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	//IsDeleted     uint      `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type EventInfo struct {
	Id        string    `json:"id"`
	Tag       TagInfo   `json:"tag"`
	EventName string    `json:"event_name"`
	City      string    `json:"city"`
	Type      EventType `json:"type"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
}

type EventType struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

type EventItem struct {
	Id          uint64  `json:"id" gorm:"id" form:"id"`
	EventId     uint64  `json:"event_id" gorm:"event_id" form:"event_id"`
	Brand       string  `json:"brand" gorm:"brand" form:"brand"`
	Name        string  `json:"name" gorm:"name" form:"name"`
	Sku         string  `json:"sku" gorm:"sku" form:"sku"`
	Barcode     string  `json:"barcode" gorm:"barcode" form:"barcode"`
	RetailPrice int     `json:"retail_price" gorm:"retail_price" form:"retail_price"`
	SalePrice   int     `json:"sale_price" gorm:"sale_price" form:"sale_price"`
	Discount    float32 `json:"discount" gorm:"discount" form:"discount"`
	Season      string  `json:"season" gorm:"season" form:"season"`
	Category    string  `json:"category" gorm:"category" form:"category"`
	Color       string  `json:"color" gorm:"color" form:"color"`
	Size        string  `json:"size" gorm:"size" form:"size"`
	Inventory   int     `json:"inventory" gorm:"inventory" form:"inventory"`
	//CreateTime  time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	//ModifyTime  time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	//IsDeleted   uint      `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type SaleRecord struct {
	Id         uint64    `json:"id" gorm:"id" form:"id"`
	MemberId   uint64    `json:"member_id" gorm:"member_id" form:"member_id"`
	ItemId     uint64    `json:"item_id" gorm:"item_id" form:"item_id"`
	OrderId    string    `json:"order_id" gorm:"order_id" form:"order_id"`
	OrderTime  time.Time `json:"order_time" gorm:"order_time" form:"order_time"`
	CouponUsed int       `json:"coupon_used" gorm:"coupon_used" form:"coupon_used"`
	Source     int       `json:"source" gorm:"source" form:"source"`
	Quantity   int       `json:"quantity" gorm:"quantity" form:"quantity"`
	PaidPrice  int       `json:"paid_price" gorm:"paid_price" form:"paid_price"`
	IsReturn   int       `json:"is_return" gorm:"is_return" form:"is_return"`
	//CreateTime time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	//ModifyTime time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	//IsDeleted  uint      `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type ReturnRecords struct {
	Id        uint64 `json:"id" gorm:"id" form:"id"`
	OrderId   string `json:"order_id" gorm:"order_id" form:"order_id"`
	PaidPrice int    `json:"paid_price" gorm:"paid_price" form:"paid_price"`
	Quantity  int    `json:"quantity" gorm:"quantity" form:"quantity"`
}

type SaleRecordDetail struct {
	Id           uint64    `json:"id" gorm:"id" form:"id"`
	MemberId     uint64    `json:"member_id" gorm:"member_id" form:"member_id"`
	Name         string    `json:"name" gorm:"name" form:"name"`
	Nickname     string    `json:"nickname" gorm:"nickname" form:"nickname"`
	Phone        string    `json:"phone" gorm:"phone" form:"phone"`
	Gender       int       `json:"gender" gorm:"gender" form:"gender"`
	MemberSource int       `json:"member_source" gorm:"member_source" form:"member_source"`
	City         string    `json:"city" gorm:"city" form:"city"`
	OrderId      string    `json:"order_id" gorm:"order_id" form:"order_id"`
	OrderTime    time.Time `json:"order_time" gorm:"order_time" form:"order_time"`
	CouponUsed   int       `json:"coupon_used" gorm:"coupon_used" form:"coupon_used"`
	SaleSource   int       `json:"sale_source" gorm:"sale_source" form:"sale_source"`
	IsReturn     int       `json:"is_return" gorm:"is_return" form:"is_return"`
	EventId      uint64    `json:"event_id" gorm:"event_id" form:"event_id"`
	Brand        string    `json:"brand" gorm:"brand" form:"brand"`
	Sku          string    `json:"sku" gorm:"sku" form:"sku"`
	Barcode      string    `json:"barcode" gorm:"barcode" form:"barcode"`
	RetailPrice  int       `json:"retail_price" gorm:"retail_price" form:"retail_price"`
	SalePrice    int       `json:"sale_price" gorm:"sale_price" form:"sale_price"`
	Discount     float32   `json:"discount" gorm:"discount" form:"discount"`
	Season       string    `json:"season" gorm:"season" form:"season"`
	Category     string    `json:"category" gorm:"category" form:"category"`
	Color        string    `json:"color" gorm:"color" form:"color"`
	Size         string    `json:"size" gorm:"size" form:"size"`
	Inventory    int       `json:"inventory" gorm:"inventory" form:"inventory"`
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
	X string `json:"x"`
	Y string `json:"y"`
}

type RankRecord struct {
	Brand    string `json:"brand" gorm:"brand" form:"brand"`
	Name     string `json:"name" gorm:"name" form:"name"`
	Sku      string `json:"sku" gorm:"sku" form:"sku"`
	Category string `json:"category" gorm:"category" form:"category"`
	Color    string `json:"color" gorm:"color" form:"color"`
	Size     string `json:"size" gorm:"size" form:"size"`
	Quantity int    `json:"quantity" gorm:"quantity" form:"quantity"`
}

type Rank struct {
	Rank     string `json:"rank"`
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
}

type DailyDetail struct {
	Date string `json:"date"`
	CoreMetric
	ReturnAmount float32 `json:"return_amount"`
	Growth       float32 `json:"growth_to_yesterday"`
}

type Industry struct {
	Id                 uint64 `json:"id" gorm:"id" form:"id"`
	IndustryCode       int    `json:"industry_code" gorm:"industry_code" form:"industry_code"`
	ParentIndustryCode int    `json:"parent_industry_code" gorm:"parent_industry_code" form:"parent_industry_code"`
	IndustryLevel      int    `json:"industry_level" gorm:"industry_level" form:"industry_level"`
	Industry           string `json:"industry" gorm:"industry" form:"industry"`
	English            string `json:"english" gorm:"english" form:"english"`
	//CreateTime         time.Time `json:"create_time" gorm:"create_time" form:"create_time"`
	//ModifyTime         time.Time `json:"modify_time" gorm:"modify_time" form:"modify_time"`
	//IsDeleted          uint      `json:"is_deleted" gorm:"is_deleted"  form:"is_deleted"`
}

type IndustryInfo struct {
	IndustryCode  int               `json:"industry_code"`
	Industry      string            `json:"industry"`
	English       string            `json:"english"`
	Subindustries []SubindustryInfo `json:"subindustries"`
}

type SubindustryInfo struct {
	SubIndustryCode int    `json:"subindustry_code"`
	SubIndustry     string `json:"subindustry"`
	English         string `json:"english"`
}
