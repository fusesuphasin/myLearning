package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty" form:"_id"`
	//ID       uint               `bson:"id" form:"id" json:"id" xml:"id"  validate:"required"`
	Name     string             `bson:"name" form:"name" json:"name" xml:"name"  validate:"required"`
	Email    string             `bson:"email" form:"email" json:"email" xml:"email"  validate:"required"`
	Password string             `bson:"password" form:"password" json:"password" xml:"password"  validate:"required"`
	Image    string             `bson:"image" form:"image" json:"image" xml:"image"  `
	Role    string             `bson:"role" form:"role" json:"role" xml:"role"  `
}

