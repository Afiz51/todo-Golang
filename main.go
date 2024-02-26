package main

import (
	"github.com/Afiz51/TodoGoRest/db"
	"github.com/Afiz51/TodoGoRest/middleware"
	"github.com/Afiz51/TodoGoRest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware())
	db.ConnectToMongoDB("TodoApp")

	//db.PrintCollection("TodoApp", "users")

	users := r.Group("/user")

	routes.UserRoutesSetup(users)

	r.Run(":9090") // listen and serve on 0.0.0.0:8080
}
