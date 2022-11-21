package domain

import "time"

type Model struct {
	Id         uint64     `json:"id"`
	CreateTime *time.Time `json:"create_time"`
	//Creator    uint64     `json:"creator"`
	ModifyTime *time.Time `json:"modify_time"`
}
