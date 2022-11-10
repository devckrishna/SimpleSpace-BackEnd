package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Submission struct{
	ID primitive.ObjectID  `bson:"_id"`
	SubmissionID string `json:"submissionID"`
	TechnologiesUsed []string `json:"technologiesUsed"`
    Title1 string `json:"title1"`
	Title2 string `json:"title2"`
	ProjectLinkVideo string `json:"projectLinkVideo"`
    ProjectLinkCode string `json:"projectLinkCode"`
	ProjectSnapShots []string `json:"projectSnapShots"`
	TeamPosition int `json:"teamPosition"`
	TeamID string `json:"teamID"`
}

