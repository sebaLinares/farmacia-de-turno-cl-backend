package datasources

import (
	"drugstore/internal/models"
	"encoding/json"
	"os"
)

type JSONDataSource struct {
	FilePath string
}

func NewJSONDataSource(filePath string) *JSONDataSource {
	return &JSONDataSource{
		FilePath: filePath,
	}
}

func (ds *JSONDataSource) GetPharmacies() ([]models.Pharmacy, error) {
	file, err := os.Open(ds.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pharmacies []models.Pharmacy
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pharmacies); err != nil {
		return nil, err
	}

	return pharmacies, nil
}
