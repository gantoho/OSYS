package models

type Goods struct {
	Model
	GoodsNo   string `json:"goods_no" form:"goods_no" gorm:"type:varchar(20);not null;unique"`
	Goodsname string `json:"goodsname" form:"goodsname" gorm:"type:varchar(20);not null;"`
}
