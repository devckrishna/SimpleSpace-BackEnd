package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID          primitive.ObjectID `bson:"_id"`
	TeamName    string             `json:"team_name" validate:"required"`
	TeamMembers []string           `json:"team_members"`
	Leader      string             `json:"leader" validate:"required"`
	HackathonID *string            `jsom:"hackathon_id,omitempty"`
}
