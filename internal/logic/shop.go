package logic

import (
	"github.com/gantoho/osys/internal/models"
	"github.com/gantoho/osys/internal/tools"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func GetShopByID(c *gin.Context) {
	var shop models.Shop
	id := c.Param("id")
	err := models.DB.Where("id = ?", id).First(&shop).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, tools.ECode{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    shop,
	})
}

func GetShops(c *gin.Context) {
	var shops []models.Shop
	err := models.DB.Find(&shops).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, tools.ECode{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    shops,
	})
}

func GetRandomShops(c *gin.Context) {
	numStr := c.Param("num")
	num, _ := strconv.Atoi(numStr)
	var shops []models.Shop
	err := models.DB.Find(&shops).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, tools.ECode{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	if len(shops) <= num {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    http.StatusOK,
			Message: "success",
			Data:    shops,
		})
		return
	}
	if num == 0 {
		num = 10
	}

	var numShops []models.Shop
	for {
		i := rand.Intn(len(shops))
		numShops = append(numShops, shops[i])
		if len(numShops) == num {
			break
		}
	}

	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    numShops,
	})
}

func AddShop(c *gin.Context) {
	var shop models.Shop
	err := c.ShouldBind(&shop)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	shop.ShopNo = strconv.FormatInt(time.Now().Unix(), 10)
	shop.Model = models.Model{
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	err = models.DB.Create(&shop).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, tools.ECode{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    shop,
	})
}
