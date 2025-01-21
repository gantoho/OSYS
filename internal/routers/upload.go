package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func initUpload(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		v1.POST("/uploads", func(c *gin.Context) {
			// multiple files
			form, _ := c.MultipartForm()
			files := form.File["files"]

			for _, file := range files {
				// save each uploaded file to specific location
				dst := fmt.Sprintf("./uploads/%s", file.Filename)
				if err := c.SaveUploadedFile(file, dst); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			}

			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("%d files uploaded successfully", len(files)),
				"files":   files,
			})
		})
		v1.POST("/upload", func(c *gin.Context) {
			// single file
			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// save uploaded file to specific location
			dst := fmt.Sprintf("./uploads/%s", file.Filename)
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "file uploaded successfully",
				"file":    file.Filename,
			})
		})
		v1.GET("/f/:filename", func(c *gin.Context) {
			// 获取 URL 中的文件名参数
			filename := c.Param("filename")

			// 指定文件所在的目录
			filepath := "./uploads/" + filename

			// 检查文件是否存在
			if _, err := os.Stat(filepath); os.IsNotExist(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
				return
			}

			// 发送文件给客户端
			c.File(filepath)
		})
	}
}
