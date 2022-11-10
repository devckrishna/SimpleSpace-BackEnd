package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	database "github.com/devckrishna/SimpleSpace/Database"
	models "github.com/devckrishna/SimpleSpace/Models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var teamCollection *mongo.Collection = database.OpenCollection(database.Client, "team")

func CreateTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var team models.Team
		defer cancel()

		hackathonId := c.Param("id")
		if err := c.BindJSON(&team); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(&team); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		newTeam := models.Team{
			ID:          primitive.NewObjectID(),
			TeamName:    team.TeamName,
			TeamMembers: []string{team.Leader},
			Leader:      team.Leader,
			HackathonID: &hackathonId,
		}
		result, err := teamCollection.InsertOne(ctx, newTeam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, result)
	}
}

func AddMemeber() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var team models.Team
		defer cancel()

		id := c.Param("id")
		user_id := c.Param("user_id")
		teamId, _ := primitive.ObjectIDFromHex(id)
		err := teamCollection.FindOne(ctx, bson.M{"_id": teamId}).Decode(&team)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := teamCollection.DeleteOne(ctx, bson.M{"_id": teamId})
		fmt.Println(res)
		newTeam := models.Team{
			ID:          team.ID,
			TeamName:    team.TeamName,
			TeamMembers: append(team.TeamMembers, user_id),
			Leader:      team.Leader,
			HackathonID: team.HackathonID,
		}
		result, err := teamCollection.InsertOne(ctx, newTeam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, result)
	}
}

func GetAllTeams() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var team []bson.M
		result, err := teamCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err = result.All(ctx, &team); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, team)
	}
}

func GetTeamById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var team models.Team
		defer cancel()

		id := c.Param("id")
		teamId, _ := primitive.ObjectIDFromHex(id)
		err := teamCollection.FindOne(ctx, bson.M{"_id": teamId}).Decode(&team)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, team)

	}
}
