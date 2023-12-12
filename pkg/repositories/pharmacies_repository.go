package repositories

import "drugstore/internal/models"

type PharmacyRepository interface {
	GetPharmacies() ([]models.Pharmacy, error)
}
