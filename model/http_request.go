package model

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Account        string `json:"account"`
	NickName       string `json:"nickname"`
	Password       string `json:"password"`
	BrandId        string `json:"brand_id"`
	InvitationCode string `json:"invitation_code"`
}

type LogOutRequest struct {
	UserId string `json:"user_id"`
}

type CreateEventRequest struct {
	EventName string `json:"event_name"`
	EventType int    `json:"event_type"`
	UserId    string `json:"user_id"`
	CreatorId string `json:"creator_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	City      string `json:"city"`
	TagId     string `json:"tag_id"`
}

type CreateBrandRequest struct {
	Brand           string `json:"event_name"`
	IndustryCode    int    `json:"industry_code"`
	SubindustryCode int    `json:"subindustry_code"`
}

type ImportEventDataRequest struct {
	UserId  string `json:"user_id"`
	EventId string `json:"event_id"`
}

type ImportMemberRequest struct {
	UserId string `json:"user_id"`
}

type CreateShareLinkRequest struct {
	ReportPassword string `json:"report_password"`
	EventId        string `json:"event_id"`
}

type CreateTagRequest struct {
	UserId string `json:"user_id"`
	Tag    string `json:"tag"`
}
