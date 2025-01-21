package models

type User struct {
	Model
	Username string `json:"username" form:"username" gorm:"type:varchar(20);not null;unique;"`
	Password string `json:"-" form:"password" gorm:"type:varchar(255);not null;"`
	Email    string `json:"email" form:"email" gorm:"type:varchar(255);"`
	Avatar   string `json:"avatar" form:"avatar" gorm:"type:varchar(255);"`
	Motto    string `json:"motto" form:"motto" gorm:"type:varchar(255);"`
}

func (User) TableName() string {
	return "user"
}
