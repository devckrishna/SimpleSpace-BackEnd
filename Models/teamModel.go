package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//todo
type Team struct {
	Name string `json:"teamName"`
	ID   primitive.ObjectID `bson:"_id"`
	TeamID string `json:"teamID"`
	Members []string `json:"teamMembers"`
}