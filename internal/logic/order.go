package logic

import (
	"errors"
	"fmt"
	"github.com/gantoho/osys/internal/models"
	"github.com/gantoho/osys/internal/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func GetOrderByID(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	err := models.DB.Preload("Products").Take(&order, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, tools.ECode{
				Code:    http.StatusOK,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    order,
	})
}

type AOrderParams struct {
	UserID     int64 `json:"user_id"`
	ProductsID []int `json:"products_id"`
	ShopID     int64 `json:"shop_id"`
}

func AddOrder(c *gin.Context) {
	var aOrderParams AOrderParams
	err := c.ShouldBindJSON(&aOrderParams)
	fmt.Printf("%+v \n", aOrderParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	var products []models.Product
	err = models.DB.Find(&products, aOrderParams.ProductsID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	order := models.Order{
		UserID: aOrderParams.UserID,
		Model: models.Model{
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		},
		OrderNo:  strconv.FormatInt(time.Now().Unix(), 10),
		Products: products,
		ShopID:   aOrderParams.ShopID,
	}

	err = models.DB.Create(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    order,
	})
}

func GetOrderByNo(c *gin.Context) {
	orderNo := c.PostForm("order_no")
	var order models.Order
	err := models.DB.Preload("Products").Preload("User").Where("order_no = ?", orderNo).Take(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    order,
	})
}

func GetOrderByUserID(c *gin.Context) {
	userID := c.Param("userId")
	var order models.Order
	err := models.DB.Preload("Products").Preload("User").Where("user_id = ?", userID).Take(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    order,
	})
}

func DelOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	err := models.DB.Where("id = ?", id).Delete(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
	})
}
