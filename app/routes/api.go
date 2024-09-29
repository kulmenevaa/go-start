package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kulmenevaa/go-start/app/controllers"
	database "github.com/kulmenevaa/go-start/app/db"
	"github.com/kulmenevaa/go-start/app/repositories"
	"github.com/kulmenevaa/go-start/app/services"
)

func ApiRoutes(prefix string, router *gin.Engine) {
	db := database.ConnectDB()
	apiGroup := router.Group(prefix)
	{
		posts := apiGroup.Group("/posts")
		{
			postRepo := repositories.NewPostRepository(db)
			postService := services.NewPostService(postRepo)
			postController := controllers.NewPostController(postService)

			posts.GET("/all", postController.GetPostList)
			posts.GET("/:id", postController.GetPost)
			posts.POST("/create", postController.CreatePost)
			posts.PUT("/update/:id", postController.UpdatePost)
			posts.DELETE("/delete/:id", postController.DeletePost)
		}
	}
}
