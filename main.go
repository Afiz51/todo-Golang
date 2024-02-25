package main

import (
	"github.com/Afiz51/TodoGoRest/db"
	"github.com/Afiz51/TodoGoRest/handlers"
	"github.com/Afiz51/TodoGoRest/middleware"
	"github.com/Afiz51/TodoGoRest/routes"
	"github.com/Afiz51/TodoGoRest/validate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(validate.UserStructLevelValidation, handlers.CreateUserRequest{})
		// v.RegisterStructValidation(routes.LoginStructLevelValidation, routes.Login{})
	}

	r.Use(middleware.LoggerMiddleware())
	db.ConnectToMongoDB()

	//db.PrintCollection("local", "startup_log")

	users := r.Group("/user")

	routes.UserRoutesSetup(users)

	r.Run(":9090") // listen and serve on 0.0.0.0:8080
}
