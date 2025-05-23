package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rochaandrey/restful-API.git/controller"
	"github.com/rochaandrey/restful-API.git/usecase"
)

func main() {
	server := gin.Default()

	MyProductUseCase := usecase.NewProductUseCase()

	ProductController := controller.NewProductController(MyProductUseCase)

	server.GET("/products", ProductController.GetProducts)
	server.Run(":8080")
}
