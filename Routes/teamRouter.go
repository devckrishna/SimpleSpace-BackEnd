package routes

import (
	controllers "github.com/devckrishna/SimpleSpace/Controllers"
	"github.com/gin-gonic/gin"
)

func TeamRoute(ctx *gin.Engine) {
	ctx.POST("/team/:id", controllers.CreateTeam())
	ctx.PUT("/team/add/:id/:user_id", controllers.AddMemeber())
	ctx.GET("/team", controllers.GetAllTeams())
	ctx.GET("/team/:id", controllers.GetTeamById())
}
