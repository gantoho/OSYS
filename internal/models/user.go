package models

import "time"

type User struct {
	ID          int64     ` json:"id" form:"id" gorm:"primary_key;auto_increment"`
	Username    string    `json:"username" form:"username" gorm:"type:varchar(20);not null;unique;"`
	Password    string    `json:"-" form:"password" gorm:"type:varchar(255);not null;"`
	Email       string    `json:"email" form:"email" gorm:"type:varchar(255);"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

func (User) TableName() string {
	return "user"
}
