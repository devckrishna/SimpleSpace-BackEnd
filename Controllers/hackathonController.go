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

var hackathonCollection *mongo.Collection = database.OpenCollection(database.Client, "hackathon")

func CreateHackathon() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var hackathon models.Hackathon
		defer cancel()

		if err := c.BindJSON(&hackathon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(&hackathon); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		newHackathon := models.Hackathon{
			ID:               primitive.NewObjectID(),
			Theme:            hackathon.Theme,
			Title:            hackathon.Title,
			StartDate:        hackathon.StartDate,
			EndDate:          hackathon.EndDate,
			CreatedAt:        time.Now(),
			Participants:     []string{},
			Mentors:          []string{},
			Description:      hackathon.Description,
			MaxTeamSize:      hackathon.MaxTeamSize,
			SocialMediaLinks: hackathon.SocialMediaLinks,
			Submissions:      []string{},
			Creater:          hackathon.Creater,
		}
		result, err := hackathonCollection.InsertOne(ctx, newHackathon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, result)
	}
}

func AddParticipant() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var hackathon models.Hackathon
		defer cancel()

		id := c.Param("id")
		hackathonId, _ := primitive.ObjectIDFromHex(id)
		userId := c.Param("user_id")

		err := hackathonCollection.FindOne(ctx, bson.M{"_id": hackathonId}).Decode(&hackathon)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := hackathonCollection.DeleteOne(ctx, bson.M{"_id": hackathonId})
		fmt.Println(res)

		newHackathon := models.Hackathon{
			ID:               hackathon.ID,
			Theme:            hackathon.Theme,
			Title:            hackathon.Title,
			StartDate:        hackathon.StartDate,
			EndDate:          hackathon.EndDate,
			CreatedAt:        time.Now(),
			Participants:     append(hackathon.Participants, userId),
			Mentors:          []string{},
			Description:      hackathon.Description,
			MaxTeamSize:      hackathon.MaxTeamSize,
			SocialMediaLinks: hackathon.SocialMediaLinks,
			Submissions:      []string{},
			Creater:          hackathon.Creater,
		}
		result, err := hackathonCollection.InsertOne(ctx, newHackathon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, result)
	}
}

func AddMentor() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var hackathon models.Hackathon
		defer cancel()

		id := c.Param("id")
		hackathonId, _ := primitive.ObjectIDFromHex(id)
		userId := c.Param("user_id")

		err := hackathonCollection.FindOne(ctx, bson.M{"_id": hackathonId}).Decode(&hackathon)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := hackathonCollection.DeleteOne(ctx, bson.M{"_id": hackathonId})
		fmt.Println(res)

		newHackathon := models.Hackathon{
			ID:               hackathon.ID,
			Theme:            hackathon.Theme,
			Title:            hackathon.Title,
			StartDate:        hackathon.StartDate,
			EndDate:          hackathon.EndDate,
			CreatedAt:        time.Now(),
			Participants:     hackathon.Participants,
			Mentors:          append(hackathon.Mentors, userId),
			Description:      hackathon.Description,
			MaxTeamSize:      hackathon.MaxTeamSize,
			SocialMediaLinks: hackathon.SocialMediaLinks,
			Submissions:      []string{},
			Creater:          hackathon.Creater,
		}
		result, err := hackathonCollection.InsertOne(ctx, newHackathon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, result)
	}
}

func GetHackathons() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var hackathons []bson.M
		result, err := hackathonCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err = result.All(ctx, &hackathons); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, hackathons)
	}
}

func GetHackathonById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var hackathon models.Hackathon
		defer cancel()

		id := c.Param("id")
		hackathonId, _ := primitive.ObjectIDFromHex(id)
		err := hackathonCollection.FindOne(ctx, bson.M{"_id": hackathonId}).Decode(&hackathon)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, hackathon)
	}
}

func DeleteHackaThon() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id := c.Param("id")
		hackathonId, _ := primitive.ObjectIDFromHex(id)
		result, err := hackathonCollection.DeleteOne(ctx, bson.M{"_id": hackathonId})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "no document with this id foung"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	}
}
