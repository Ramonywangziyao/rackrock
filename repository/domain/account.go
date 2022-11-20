package domain

import "time"

type Account struct {
	Model
	Username      string     `json:"username"`
	Password      string     `json:"password"`
	Status        int8       `json:"status"`
	LastLoginTime *time.Time `json:"last_login_time"`
}

func (self *Account) TableName() string {
	return "account"
}
