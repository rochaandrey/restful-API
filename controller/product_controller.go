package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rochaandrey/restful-API.git/model"
	"github.com/rochaandrey/restful-API.git/usecase"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "batata",
			Price: 20,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
