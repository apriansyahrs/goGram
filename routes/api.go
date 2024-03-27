package routes

import (
	"goGram/controllers"
	"goGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
		userGroup.Use(middlewares.Authentication())
		userGroup.PUT("/:userId", controllers.UpdateUser)
		userGroup.DELETE("/delete", controllers.DeleteUser)
	}

	photoGroup := r.Group("/photos")
	{
		photoGroup.Use(middlewares.Authentication())
		photoGroup.POST("/", controllers.CreatePhoto)
		photoGroup.GET("/", controllers.GetAllPhoto)
		photoGroup.GET("/:photoId", middlewares.PhotoAuthorization(), controllers.GetOnePhoto)
		photoGroup.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoGroup.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentGroup := r.Group("/comments")
	{
		commentGroup.Use(middlewares.Authentication())
		commentGroup.POST("/", controllers.CreateComment)
		commentGroup.GET("/", controllers.GetAllComment)
		commentGroup.GET("/:commentId", middlewares.CommentAuthorization(), controllers.GetOneComment)
		commentGroup.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentGroup.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaGroup := r.Group("/socialmedias")
	{
		socialMediaGroup.Use(middlewares.Authentication())
		socialMediaGroup.POST("/", controllers.CreateSocialMedia)
		socialMediaGroup.GET("/", controllers.GetAllSocialMedia)
		socialMediaGroup.GET("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.GetOneSocialMedia)
		socialMediaGroup.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaGroup.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
