package mappings

import (
	"money-game-2/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api")
	{
		//user
		v1.GET("user", controllers.GetUsers)
		v1.POST("user", controllers.CreateUser)
		v1.GET("user/:id", controllers.GetUserByID)
		v1.PUT("user/:id", controllers.UpdateUser)
		v1.DELETE("user/:id", controllers.DeleteUser)

		//activities
		v1.GET("activities", controllers.GetActivities)

		//activities
		v1.GET("categories", controllers.GetCategories)
	}
	return r
}
