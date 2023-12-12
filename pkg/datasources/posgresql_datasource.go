package datasources

// import (
// 	"database/sql"
// 	"drugstore/internal/models"
// 	"log"
// )

// type DBDataSource struct {
// 	DB *sql.DB
// }

// func NewDBDataSource(db *sql.DB) *DBDataSource {
// 	return &DBDataSource{
// 		DB: db,
// 	}
// }

// func (ds *DBDataSource) GetPharmacies() ([]models.Pharmacy, error) {
// 	// Implement logic to query pharmacies from the database
// 	// Example: rows, err := ds.DB.Query("SELECT * FROM pharmacies")

// 	// Handle errors
// 	if err != nil {
// 		log.Println("Error querying database:", err)
// 		return nil, err
// 	}

// 	// Process the rows and return pharmacy data
// 	var pharmacies []models.Pharmacy
// 	for rows.Next() {
// 		var pharmacy models.Pharmacy
// 		// Scan rows into the pharmacy struct
// 		// Example: rows.Scan(&pharmacy.ID, &pharmacy.Name, ...)
// 		pharmacies = append(pharmacies, pharmacy)
// 	}

// 	return pharmacies, nil
// }
