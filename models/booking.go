package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	CheckInDate  time.Time          `bson:"check_in_date" json:"check_in_date"`
	CheckOutDate time.Time          `bson:"check_out_date" json:"check_out_date"`
	GuestCount   int                `bson:"guest_count" json:"guest_count"`
	RoomType     string             `bson:"room_type" json:"room_type"`
	PricePerNight float64           `bson:"price_per_night" json:"price_per_night"`
}
