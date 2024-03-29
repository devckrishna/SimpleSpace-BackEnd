package routes

import (
	controller "github.com/devckrishna/SimpleSpace/Controllers"
	"github.com/gin-gonic/gin"
)

func HackathonRoute(ctx *gin.Engine) {
	ctx.POST("/hackathon", controller.CreateHackathon())
	ctx.GET("/hackathon", controller.GetHackathons())
	ctx.GET("/hackathon/:id", controller.GetHackathonById())
	ctx.DELETE("/hackathon/:id", controller.DeleteHackaThon())
	ctx.PUT("/hackathon/:id/:user_id", controller.AddParticipant())
	ctx.PUT("/hackathon/mentor/:id/:user_id", controller.AddMentor())
}
