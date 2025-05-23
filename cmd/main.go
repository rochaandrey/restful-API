package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rochaandrey/restful-API.git/controller"
	"github.com/rochaandrey/restful-API.git/db"
	"github.com/rochaandrey/restful-API.git/repository"
	"github.com/rochaandrey/restful-API.git/usecase"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//inicializacao da camada de repository,usecase e controller
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/products", ProductController.GetProducts)
	server.Run(":8080")
}
