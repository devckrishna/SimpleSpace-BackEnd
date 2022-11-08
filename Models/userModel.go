package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `json:"name" validate:"required,min=2,max=100"`
	Password  *string            `json:"password" validate:"required,min=6"`
	Email     *string            `json:"email" validate:"email,required"`
	Token     *string            `json:"token"`
	User_type *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER|eq=MENTOR"`
	User_id   string             `json:"user_id"`
}
