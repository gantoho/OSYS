package models

type Order struct {
	Model
	OrderNo string `json:"order_no" form:"order_no" gorm:"type:varchar(20);not null;unique;"`

	Products []Product `gorm:"many2many:order2product"`
}
