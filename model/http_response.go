package model

import "time"

type BrandListResponse struct {
	Brands []BrandInfo `json:"brands"`
}

type AccountDetailResponse struct {
	Account  string    `json:"account"`
	Nickname string    `json:"nickname"`
	Brand    BrandInfo `json:"brand"`
}

type DashboardBasicResponse struct {
	Nickname        string `json:"nickname"`
	TotalEvent      int    `json:"total_event"`
	TotalAmountSold int    `json:"total_amount_sold"`
	TotalItemSold   int    `json:"total_item_sold"`
}

type EventListResponse struct {
	Events      []EventInfo `json:"events"`
	CurrentPage int         `json:"current_page"`
	TotalPage   int         `json:"total_page"`
	PageSize    int         `json:"page_size"`
}

type LoginResponse struct {
	Account     string    `json:"account"`
	Nickname    string    `json:"nickname"`
	LoginIp     string    `json:"login_ip"`
	LoginTime   time.Time `json:"login_time"`
	Token       string    `json:"token"`
	AccessLevel int       `json:"access_level"`
}

type UserListResponse struct {
	Users []UserInfo `json:"users"`
}

type TagListResponse struct {
	Tags []TagInfo `json:"tags"`
}

type ReportResponse struct {
	ReportStatus     int             `json:"report_status"`
	EventInfo        EventInfo       `json:"event_info"`
	CurrentStartTime string          `json:"current_start_time"`
	CurrentEndTime   string          `json:"current_end_time"`
	BrandInfo        BrandInfo       `json:"brand_info"`
	CoreMetric       CoreMetric      `json:"core_metric"`
	SecondaryMetric  SecondaryMetric `json:"secondary_metric"`
	Distribution     Distribution    `json:"distribution"`
}

type RankingResponse struct {
	Ranks       []Rank `json:"ranks"`
	CurrentPage int    `json:"current_page"`
	TotalPage   int    `json:"total_page"`
	PageSize    int    `json:"page_size"`
}

type DailyDetailResponse struct {
	Detail []DailyDetail `json:"detail"`
}

type CityResponse struct {
	Cities        []string `json:"cities"`
	CitiesEnglish []string `json:"cities_english"`
}

type IndustryResponse struct {
	Industries []IndustryInfo `json:"industries"`
}

type RockResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
