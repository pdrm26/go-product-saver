package routes

import (
	"github.com/gin-gonic/gin"
	"my-module/controllers"
)

type ProductRouteController struct {
	productController controllers.ProductController
}

func NewRouteProductController(productController controllers.ProductController) ProductRouteController {
	return ProductRouteController{productController}
}

func (pc *ProductRouteController) ProductRoute(rg *gin.RouterGroup) {
	router := rg.Group("products")
	//middleware
	router.GET("/", pc.productController.FindProducts)
}
