package routes

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yourusername/hotel-booking-api/controllers"
)

func AuthRoutes(r *mux.Router, client *mongo.Client) {
	r.HandleFunc("/register", controllers.Register(client)).Methods("POST")
	r.HandleFunc("/login", controllers.Login(client)).Methods("POST")
}
