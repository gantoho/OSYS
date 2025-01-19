package models

type Order struct {
	Model
	OrderNo string `json:"order_no" form:"order_no" gorm:"type:varchar(20);not null;unique;"`

	Products []Product `gorm:"many2many:order2product"`

	UserID int64
	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;"`

	ShopID int64
	Shop   Shop `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE;"`
}

func (Order) TableName() string {
	return "order"
}
