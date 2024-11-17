package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/users", getUsers)    // GET, POST, PUT, PATCH, DELETE
	server.GET("/users/:id", getUser) // /events/1, /events/5
	server.POST("/users", createUser)
	server.PUT("/users/:id", updateUser)
	server.DELETE("/users/:id", deleteUser)
}
