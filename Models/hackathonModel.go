package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hackathon struct {
	ID               primitive.ObjectID `bson:"_id"`
	Theme            string             `json:"theme" validate:"required"`
	Title            string             `json:"title" validate:"required"`
	StartDate        string             `json:"start_date" validate:"required"`
	EndDate          string             `json:"end_date" validate:"required"`
	CreatedAt        time.Time          `bson:"created_at"`
	Participants     []string           `json:"participants"`
	Mentors          []string           `json:"mentors"`
	Description      string             `json:"description" validate:"required"`
	MaxTeamSize      int                `json:"max_team_size" validate:"required"`
	SocialMediaLinks []string           `json:"social_media_links" validate:"required"`
	Submissions      []string           `json:"submissions"`
	Creater          string             `json:"creater" validate:"required"`
}
