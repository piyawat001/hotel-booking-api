package routes

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yourusername/hotel-booking-api/controllers"
)

func BookingRoutes(r *mux.Router, client *mongo.Client) {
	r.HandleFunc("/booking", controllers.CreateBooking(client)).Methods("POST")
	r.HandleFunc("/review", controllers.CreateReview(client)).Methods("POST")

}
