package routes

import (
	controllers "github.com/devckrishna/SimpleSpace/Controllers"
	"github.com/gin-gonic/gin"
)

func SubmissionRoute(ctx *gin.Engine) {
	ctx.POST("/submission", controllers.CreateSubmission())
	ctx.GET("/submission", controllers.GetAllSubmissions())
	ctx.GET("/submission/:id", controllers.GetSubmissionById())
}
