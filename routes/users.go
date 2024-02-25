package routes

import (
	"github.com/Afiz51/TodoGoRest/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutesSetup(users *gin.RouterGroup) {

	users.POST("/creates", handlers.CreateUserHandler)

}
