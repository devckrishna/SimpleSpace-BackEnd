package controllers

import (
	"context"

	// "log"
	"net/http"
	"time"

	database "github.com/devckrishna/SimpleSpace/Database"
	// helper "github.com/devckrishna/SimpleSpace/Helpers"
	models "github.com/devckrishna/SimpleSpace/Models"
	"github.com/gin-gonic/gin"

	// "github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "golang.org/x/crypto/bcrypt"
)

var submissionCollection *mongo.Collection = database.OpenCollection(database.Client,"submissions")


func CreateSubmission() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var submission models.Submission
        defer cancel()
        
        //validate the request body
        if err := c.BindJSON(&submission); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body1"})
            return
        }
        
        //use the validator library to validate required fields
        if validationErr := validate.Struct(&submission); validationErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body2"})
            return
        }
        submission.ID=primitive.NewObjectID()
        submission.SubmissionID=submission.ID.Hex()
        result, err := submissionCollection.InsertOne(ctx, submission)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err})
            return
        }
        c.JSON(http.StatusCreated, result)
    }
}

func GetASubmission() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        submissionID := c.Param("submissionID")
        var submission models.Submission
        defer cancel()
        
        err := submissionCollection.FindOne(ctx, bson.M{"submissionID": submissionID}).Decode(&submission)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error":"error"})
            return
        }
        c.JSON(http.StatusOK, submission)
    }
}

func GetAllSubmissions() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var submissions []models.Submission
        defer cancel()
      
        results, err := submissionCollection.Find(ctx, bson.M{})
      
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error":"Some error1"})
            return
        }
      
        defer results.Close(ctx)
        for results.Next(ctx) {
            var submission models.Submission
            if err = results.Decode(&submission); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error":"Some error2"})
            }
          
            submissions = append(submissions, submission)
        }
        c.JSON(http.StatusOK,submissions)
    }
}

func UpdateSubmission() gin.HandlerFunc{
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        submissionID := c.Param("submissionID")
        var submission models.Submission
        defer cancel()
      
      
        //validate the request body
        if err := c.BindJSON(&submission); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error":"error 1"})
            return
        }
      
        //use the validator library to validate required fields
        if validationErr := validate.Struct(&submission); validationErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error":"error 2"})
            return
        }
        result, err := submissionCollection.UpdateOne(ctx, bson.M{"submissionID": submissionID}, bson.M{"$set": submission})
      
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error":"error 3"})
            return
        }
      
        //get updated user details
        var updatedSubmission models.Submission
        if result.MatchedCount == 1 {
            err := submissionCollection.FindOne(ctx, bson.M{"submissionID": submissionID}).Decode(&updatedSubmission)
            if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error":"error 4"})
                return
            }
        }
        c.JSON(http.StatusOK, updatedSubmission)
    }
}
