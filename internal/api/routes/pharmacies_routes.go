package routes

import (
	"drugstore/internal/api/handlers"
	"drugstore/pkg/repositories"

	"github.com/gorilla/mux"
)

func NewRouter(shiftDataSource repositories.PharmacyRepository, nonShiftDataSource repositories.PharmacyRepository) *mux.Router {
	r := mux.NewRouter()

	// Define routes for shift and non-shift pharmacies
	r.HandleFunc("/shift-pharmacies", handlers.PharmaciesHandler(shiftDataSource)).Methods("GET")
	r.HandleFunc("/non-shift-pharmacies", handlers.PharmaciesHandler(nonShiftDataSource)).Methods("GET")

	return r
}

// func SetupPharmaciesRoutes(router *mux.Router) {
// 	router.HandleFunc("/api/pharmacies", handlers.HandleGetPharmacies).
// 		Methods("GET")
// }
