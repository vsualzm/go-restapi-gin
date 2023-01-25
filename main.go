package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vsualzm/go-restapi-gin/controllers/productcontroller"
	"github.com/vsualzm/go-restapi-gin/model"
)

func main() {
	r := gin.Default()
	model.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
