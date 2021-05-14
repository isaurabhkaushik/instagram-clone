package Routes

import (
	"github.com/gin-gonic/gin"
	Controllers "github.com/isaurabhkaushik/hp/service/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)

		//LIKE Routes
		v1.POST("like", Controllers.CreateALike)
		v1.DELETE("like/:id", Controllers.DeleteALike)
		v1.GET("like/:id", Controllers.GetALike)

		// COMMENT Routes
		v1.POST("comment", Controllers.CreateAComment)
		v1.DELETE("comment/:id", Controllers.DeleteAComment)
		v1.GET("comment/:id", Controllers.GetAComment)
		v1.PUT("comment/:id", Controllers.UpdateAComment)

		// POST Routes
		v1.POST("post", Controllers.CreateAPost)
		v1.DELETE("post/:id", Controllers.DeleteAPost)
		v1.GET("post/:id", Controllers.GetAPost)
		v1.PUT("post/:id", Controllers.UpdateAPost)
	}

	return r
}
