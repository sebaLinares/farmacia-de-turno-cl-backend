package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"drugstore/internal/api/routes"
	"drugstore/internal/config"
	"drugstore/internal/services"
	"drugstore/pkg/datasources"
	"drugstore/pkg/repositories"

	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("APP_ENV")

	if env == "dev" {
		if err := godotenv.Load(".env.dev"); err != nil {
			log.Fatalf("Error loading .env.dev file: %v", err)
		}
	} else if env == "staging" {
		if err := godotenv.Load(".env.staging"); err != nil {
			log.Fatalf("Error loading .env.staging file: %v", err)
		}
	} else if env == "prod" {
		if err := godotenv.Load(".env.prod"); err != nil {
			log.Fatalf("Error loading .env.prod file: %v", err)
		}
	} else {
		log.Fatal("Unknown environment. Set APP_ENV environment variables")
	}
	// Define datasource from env
	cfg := config.LoadConfig()
	datasourceType := cfg.DataSource
	log.Println("Datasource type: ", datasourceType)

	var shiftDataSource repositories.PharmacyRepository
	var nonShiftDataSource repositories.PharmacyRepository
	var shiftPharmacyService services.PharmaciesService
	var nonShiftPharmacyService services.PharmaciesService

	if datasourceType == "json" {
		shiftDataSource = datasources.NewJSONDataSource("pkg/data/shift-drugstores.json")
		nonShiftDataSource = datasources.NewJSONDataSource("pkg/data/non-shift-drugstores.json")
		shiftPharmacyService = *services.NewPharmacyService(shiftDataSource)
		nonShiftPharmacyService = *services.NewPharmacyService(nonShiftDataSource)
	} else if datasourceType == "postgresql" {
		// dataSource = datasources.NewPostgreSQLDataSource()
	} else {
		fmt.Println("Invalid datasource type")
		return
	}

	router := routes.NewRouter(&nonShiftPharmacyService, &shiftPharmacyService)
	http.Handle("/", router)

	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
