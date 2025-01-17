package models

import "time"

type Model struct {
	ID          int64     ` json:"id" form:"id" gorm:"primary_key;auto_increment"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}
