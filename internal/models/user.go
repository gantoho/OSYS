package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"type:varchar(20);not null;unique;"`
	Password string `json:"password" form:"password" gorm:"type:varchar(255);not null;"`
	Phone    string `json:"phone" form:"phone" gorm:"type:varchar(11);"`
}

func (User) TableName() string {
	return "user"
}
