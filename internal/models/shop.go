package models

type Shop struct {
	Model
	ShopNo   string `json:"shop_no" form:"shop_no" gorm:"type:varchar(20);not null;unique;"`
	ShopName string `json:"shop_name" form:"shop_name" gorm:"type:varchar(20);not null;unique;"`
}

func (Shop) TableName() string {
	return "shop"
}
