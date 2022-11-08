package routes

import (
	controller "github.com/devckrishna/SimpleSpace/Controllers"
	middleware "github.com/devckrishna/SimpleSpace/Middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(ctx *gin.Engine) {
	ctx.Use(middleware.Authenticate())
	ctx.GET("/users", controller.GetUsers())
	ctx.GET("/users/:user_id", controller.GetUser())
}
