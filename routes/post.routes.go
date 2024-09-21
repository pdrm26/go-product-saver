package routes

import (
	"github.com/gin-gonic/gin"
	"my-module/controllers"
)

type PostRouteController struct {
	postController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}

func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {
	router := rg.Group("posts")
	//middleware
	router.GET("/", pc.postController.FindPosts)
	router.POST("/", pc.postController.CreatePost)
}
