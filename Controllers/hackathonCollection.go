package controllers

import (
	"context"
	"fmt"

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

var hackathonCollection *mongo.Collection = database.OpenCollection(database.Client,"hackathon")


func CreateOneHackathon() gin.HandlerFunc {
    return func(c *gin.Context) {
		fmt.Println("inside create one hacakthon")
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var hackathon models.Hackathon
        defer cancel()
        
        //validate the request body
        if err := c.BindJSON(&hackathon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body1"})
            return
        }
        
        //use the validator library to validate required fields
        if validationErr := validate.Struct(&hackathon); validationErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body2"})
            return
        }
        hackathon.ID=primitive.NewObjectID()
        hackathon.HackathonID=hackathon.ID.Hex()
        result, err := hackathonCollection.InsertOne(ctx, hackathon)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err})
            return
        }
        c.JSON(http.StatusCreated, result)
    }
}

func GetAHackathon() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        hackathonID := c.Param("hackathon_id")
        fmt.Println("ffffffffffffffffffff")
        fmt.Println(hackathonID)
        var hackathon models.Hackathon
        defer cancel()
        
        err := hackathonCollection.FindOne(ctx, bson.M{"hackathonid": hackathonID}).Decode(&hackathon)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error":err})
            return
        }
        c.JSON(http.StatusOK, hackathon)

    }
}

func GetAllHackathons() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var hackathons []models.Hackathon
        defer cancel()
      
        results, err := hackathonCollection.Find(ctx, bson.M{})
      
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error":"Some error1"})
            return
        }
      
        defer results.Close(ctx)
        for results.Next(ctx) {
            var hackathon models.Hackathon
            if err = results.Decode(&hackathon); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error":"Some error2"})
            }
          
            hackathons = append(hackathons, hackathon)
        }
        c.JSON(http.StatusOK,hackathons)
    }
}

func UpdateHackathon() gin.HandlerFunc{
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        hackathonID := c.Param("hackathon_id")
        var hackathon models.Hackathon 
        defer cancel()
    
        //validate the request body
        if err := c.BindJSON(&hackathon); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error":"error 1"})
            return
        }
       
        // //use the validator library to validate required fields
        // if validationErr := validate.Struct(&hackathon); validationErr != nil {
        //     c.JSON(http.StatusBadRequest, gin.H{"error":"error 2"})
        //     return
        // }
        result, err := hackathonCollection.UpdateOne(ctx, bson.M{"hackathonid": hackathonID}, bson.M{"$set": hackathon})
        fmt.Println(hackathon)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error":"error 3"})
            return
        }
      
        //get updated user details
        var updatedHackathon models.Hackathon
        if result.MatchedCount == 1 {
            err := hackathonCollection.FindOne(ctx, bson.M{"hackathonid": hackathonID}).Decode(&updatedHackathon)
            if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error":"error 4"})
                return
            }
        }
        c.JSON(http.StatusOK, updatedHackathon)
    }
}
