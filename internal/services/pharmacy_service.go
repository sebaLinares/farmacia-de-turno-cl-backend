package services

import (
	"drugstore/internal/models"
	"drugstore/pkg/repositories"
)

// PharmaciesService is the struct that handles the business logic
type PharmaciesService struct {
	repo repositories.PharmacyRepository
}

// NewFarmacyService creates a new instance of PharmaciesService
func NewPharmacyService(repo repositories.PharmacyRepository) *PharmaciesService {
	return &PharmaciesService{repo: repo}
}

// GetPharmacies returns a list of pharmacies
func (s *PharmaciesService) GetPharmacies() ([]models.Pharmacy, error) {
	// pharmacies, err := s.repo.GetPharmacies()
	// if err != nil {
	// 	return nil, err
	// }

	// return pharmacies, nil
	return s.repo.GetPharmacies()
}
