package main

import (
	"fmt"
	"os"

	database "github.com/devckrishna/SimpleSpace/Database"
	helpers "github.com/devckrishna/SimpleSpace/Helpers"
	routes "github.com/devckrishna/SimpleSpace/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	server := gin.New()
	server.Use(helpers.CORSMiddleware())
	server.Use(gin.Logger())
	routes.HackathonRoute(server)
	routes.TeamRoute(server)
	routes.AuthRoutes(server)
	routes.UserRoutes(server)

	server.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, "hye!!!")
	})

	collection := database.OpenCollection(database.Client, "users")

	fmt.Println(collection)

	server.Run(":" + port)
}
