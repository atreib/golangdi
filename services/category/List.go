package services

import (
	"github.com/atreib/golangdi/entities"
	"github.com/atreib/golangdi/infra/repositories"
)

type ListCategoriesService struct {
	CategoryRepository repositories.ICategoryRepository
}

func (s *ListCategoriesService) List() ([]*entities.Category, error) {
	categories, err := s.CategoryRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return categories, nil
}
