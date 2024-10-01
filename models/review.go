package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	Username string             `bson:"username" json:"username"`
	Rating   int                `bson:"rating" json:"rating"`
	Comment  string             `bson:"comment" json:"comment"`
}
