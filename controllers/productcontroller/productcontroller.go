package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsualzm/go-restapi-gin/model"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []model.Product

	model.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {

	var product model.Product
	id := c.Param("id")

	if err := model.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Data tidak ditemukan"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Create(c *gin.Context) {

	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	model.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Update(c *gin.Context) {
	var product model.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	if model.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasi di perbaharui"})

}

func Delete(c *gin.Context) {

	var product model.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if model.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasi di hapus"})

}
