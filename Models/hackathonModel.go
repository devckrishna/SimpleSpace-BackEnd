package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Hackathon struct{
	ID  primitive.ObjectID `bson:"_id"`
	HackathonID string `json:"hackathonid"`
	Theme string          `json:"theme"`
	Title string             `json:"title"`
	StartDate time.Time     `json:"startDate"`
	EndDate time.Time         `json:"endDate"`
	CreatedAt time.Time   `json:"createdAt"`
	Participants []string    `json:"participants"`
	Mentors     []string    `json:"mentors"`
	Description string `json:"description"`
	MaxTeamSize int `json:"maxTeamSize"`
	RegistrationCost int `json:"registrationCost"`
	SocialMediaLinks []string `json:"socialMediaLinks"`
	Submissions []string `json:"submissions"`
}