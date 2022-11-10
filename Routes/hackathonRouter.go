package routes

import (
	controller "github.com/devckrishna/SimpleSpace/Controllers"
	"github.com/gin-gonic/gin"
)

func HackathonRoute(ctx *gin.Engine) {
	ctx.POST("hackathon/create", controller.CreateOneHackathon())
	ctx.GET("hackathons",controller.GetAllHackathons())
	ctx.GET("hackathons/:hackathon_id",controller.GetAHackathon())
	ctx.PUT("hackathons/:hackathon_id",controller.UpdateHackathon())
}
