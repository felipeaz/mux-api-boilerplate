package service

import (
	"mux-api-boilerplate/internal/app/models"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s Service) GetById(id string) (models.SampleModel, error) {
	return models.NewSample(), nil
}

func (s Service) GetAll() ([]models.SampleModel, error) {
	return models.NewSampleArray(), nil
}

func (s Service) Create(rawObj models.SampleModel) (models.SampleModel, error) {
	return models.NewSample(), nil
}

func (s Service) Update(id string, rawObj models.SampleModel) error {
	return nil
}

func (s Service) UpdateAttribute(id string, attribute string, value any) error {
	return nil
}

func (s Service) Remove(id string) error {
	return nil
}
