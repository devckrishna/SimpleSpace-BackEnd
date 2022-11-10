package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Submission struct {
	ID               primitive.ObjectID `bson:"_id"`
	Title            string             `json:"title" validate:"required"`
	Description      string             `json:"description" validate:"required"`
	TechnologiesUsed []string           `json:"tech_used" validate:"required"`
	ProjectLink      string             `json:"project_link" validate:"required"`
	VideoLink        string             `json:"video_link" validate:"required"`
	HackathonID      string             `json:"hackathon_id" validate:"required"`
	TeamID           string             `json:"team_id" validate:"required"`
}
