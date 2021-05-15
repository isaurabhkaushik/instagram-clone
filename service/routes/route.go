package routes

import (
	"github.com/gin-gonic/gin"
	Controllers "github.com/isaurabhkaushik/hp/service/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		//LIKE routes
		v1.POST("like", Controllers.CreateALike)
		v1.DELETE("like/:id", Controllers.DeleteALike)
		v1.GET("like/:id", Controllers.GetALike)

		// COMMENT routes
		v1.POST("comment", Controllers.CreateAComment)
		v1.DELETE("comment/:id", Controllers.DeleteAComment)
		v1.GET("comment/:id", Controllers.GetAComment)
		v1.PUT("comment", Controllers.UpdateAComment)

		// POST routes
		v1.POST("post", Controllers.CreateAPost)
		v1.DELETE("post/:id", Controllers.DeleteAPost)
		v1.GET("post/:id", Controllers.GetAPost)
		v1.GET("posts/:id", Controllers.GetAllPost)
		v1.PUT("post", Controllers.UpdateAPost)

		// USER_PROFILE routes
		v1.POST("user-profile", Controllers.CreateAUserProfile)
		v1.DELETE("user-profile", Controllers.DeleteAUserProfile)
		v1.GET("user-profile", Controllers.GetAUserProfile)
		v1.PUT("user-profile", Controllers.UpdateAUserProfile)
	}

	return r
}
