package repositories

import (
	"github.com/atreib/golangdi/entities"
)

type InMemoryCategoryRepository struct {
}

func (*InMemoryCategoryRepository) GetAll() ([]*entities.Category, error) {
	var categories []*entities.Category

	categories = append(
		categories,
		&entities.Category{
			Id:             "1",
			NewCategoryDto: entities.NewCategoryDto{Name: "Games"},
		},
		&entities.Category{
			Id:             "2",
			NewCategoryDto: entities.NewCategoryDto{Name: "Books"},
		},
	)

	return categories, nil
}
