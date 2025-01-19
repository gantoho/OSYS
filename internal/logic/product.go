package logic

import (
	"fmt"
	"github.com/gantoho/osys/internal/models"
	"github.com/gantoho/osys/internal/tools"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func GetProductByNo(c *gin.Context) {
	productNo := c.PostForm("product_no")
	var product models.Product
	err := models.DB.Where("product_no=?", productNo).First(&product).Error
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
		Data:    product,
	})
}

func GetProductByID(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	err := models.DB.Take(&product, id).Error
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
		Data:    product,
	})
}

func GetProduct(c *gin.Context) {
	var products []models.Product
	err := models.DB.Find(&products).Error
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
		Data:    products,
	})
}

type AProduct struct {
	ProductNo   string `json:"product_no"`
	ProductName string `form:"product_name" json:"product_name" binding:"required"`
	ShopID      int64  `json:"shop_id"`
}

func AddProduct(c *gin.Context) {
	var aProduct AProduct
	err := c.ShouldBind(&aProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	uuidV4 := uuid.New()
	aProduct.ProductNo = uuidV4.String()
	fmt.Printf("%+v \n", aProduct)
	err = models.DB.Create(&models.Product{
		ProductNo:   aProduct.ProductNo,
		ProductName: aProduct.ProductName,
		Model: models.Model{
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		},
		ShopID: aProduct.ShopID,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, tools.ECode{
		Code:    http.StatusCreated,
		Message: "Product successfully added",
		Data:    aProduct,
	})
}

func EditProduct(c *gin.Context) {
	id := c.Param("id")
	productNameNew := c.PostForm("product_name_new")
	var eProduct models.Product
	err := models.DB.Find(&eProduct, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	if eProduct.ProductName == productNameNew {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "New Product Name == Old Product Name",
		})
		return
	}
	eProduct.ProductName = productNameNew
	err = models.DB.Save(&eProduct).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "Product successfully edited",
		Data:    eProduct,
	})
}

func DelProduct(c *gin.Context) {
	id := c.Param("id")
	err := models.DB.Delete(&models.Product{}, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "Product successfully deleted",
	})
}
