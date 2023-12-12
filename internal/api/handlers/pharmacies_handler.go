package handlers

import (
	"drugstore/pkg/repositories"
	"encoding/json"
	"log"
	"net/http"
)

func PharmaciesHandler(dataSource repositories.PharmacyRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve data from the data source
		pharmacies, err := dataSource.GetPharmacies()
		log.Println("Pharmacies: ", pharmacies)

		if err != nil {
			http.Error(w, "Error fetching pharmacies", http.StatusInternalServerError)
			return
		}

		// Serialize the pharmacies slice to JSON
		responseJSON, err := json.Marshal(pharmacies)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to JSON
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response
		w.Write(responseJSON)
	}
}
