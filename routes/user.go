package routes

import (
	usercontroller "example.com/gin-api/controllers/user"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")

	//get all user
	routerGroup.GET("/", usercontroller.GetAll)

	//get all user
	routerGroup.POST("/register", usercontroller.Register)

	//get all user
	routerGroup.POST("/login", usercontroller.Login)

	//get all user
	routerGroup.GET("/:id", usercontroller.GetById)

	//get all user
	routerGroup.GET("/search", usercontroller.SearchByFullname)

}
