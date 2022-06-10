package domain

import "mux-api-boilerplate/internal/app/models"

// Service is a port that defines the interaction between the Handlers & code itself.
type Service interface {
	GetById(id string) (models.SampleModel, error)
	GetAll() ([]models.SampleModel, error)
	Create(rawObj models.SampleModel) (models.SampleModel, error)
	Update(id string, rawObj models.SampleModel) error
	UpdateAttribute(id string, attribute string, value any) error
	Remove(id string) error
}
