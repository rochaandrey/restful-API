package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rochaandrey/restful-API.git/controller"
)

func main() {
	server := gin.Default()

	ProductUseCase := usecase.newProductUseCase()

	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/products", ProductController.GetProducts)
	server.Run(":8080")
}
