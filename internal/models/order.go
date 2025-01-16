package models

type Order struct {
	Model
	OrderNo string `json:"order_no" form:"order_no" gorm:"type:varchar(20);not null;unique;"`

	Goods []Goods `gorm:"many2many:order2goods"`
}
