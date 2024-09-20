package routes

import (
	"github.com/gin-gonic/gin"
	"my-module/controllers"
)

type ProductRouteController struct {
	postController controllers.PostController
}

func NewRouteProductController(productController controllers.PostController) ProductRouteController {
	return ProductRouteController{productController}
}

func (pc *ProductRouteController) ProductRoute(rg *gin.RouterGroup) {
	router := rg.Group("posts")
	//middleware
	router.GET("/", pc.postController.FindPosts)
	router.POST("/", pc.postController.CreatePost)
}
