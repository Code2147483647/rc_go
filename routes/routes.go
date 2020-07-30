package routes

import "github.com/gin-gonic/gin"

// InitUserRoute init user route
func InitUserRoute(route *gin.Engine) {
	userRouter := route.Group("/user")
	userRouter.GET("/get", GetUser)
	userRouter.POST("/", AddUser)
}
