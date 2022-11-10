package controllers

import (
	"context"
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

var submissionCollection *mongo.Collection = database.OpenCollection(database.Client, "submission")

func CreateSubmission() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var submission models.Submission
		defer cancel()

		if err := c.BindJSON(&submission); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(&submission); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		newSubmission := models.Submission{
			ID:               primitive.NewObjectID(),
			Title:            submission.Title,
			Description:      submission.Description,
			TechnologiesUsed: submission.TechnologiesUsed,
			ProjectLink:      submission.ProjectLink,
			VideoLink:        submission.VideoLink,
			HackathonID:      submission.HackathonID,
			TeamID:           submission.TeamID,
		}
		result, err := submissionCollection.InsertOne(ctx, newSubmission)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, result)
	}
}

func GetAllSubmissions() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var submissions []bson.M
		result, err := submissionCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err = result.All(ctx, &submissions); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, submissions)
	}
}

func GetSubmissionById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var submission models.Submission
		defer cancel()

		id := c.Param("id")
		submissionId, _ := primitive.ObjectIDFromHex(id)
		err := submissionCollection.FindOne(ctx, bson.M{"_id": submissionId}).Decode(&submission)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, submission)
	}
}
