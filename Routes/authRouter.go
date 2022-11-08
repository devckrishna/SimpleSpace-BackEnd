package routes

import (
	controller "github.com/devckrishna/SimpleSpace/Controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(ctx *gin.Engine) {
	ctx.POST("users/signup", controller.Signup())
	ctx.POST("users/login", controller.Login())
}
