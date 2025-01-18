package models

type Product struct {
	Model
	ProductNo   string `json:"product_no" form:"product_no" gorm:"type:varchar(50);not null;unique"`
	ProductName string `json:"product_name" form:"product_name" gorm:"type:varchar(20);not null;"`

	Orders []Order `gorm:"many2many:order2product;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;"`
}
