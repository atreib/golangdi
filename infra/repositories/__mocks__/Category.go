package mocks

import (
	"github.com/atreib/golangdi/entities"
	"github.com/atreib/golangdi/infra/repositories"
)

var mockCategories = []*entities.Category{
	{
		Id:             "1",
		NewCategoryDto: entities.NewCategoryDto{Name: "Games"},
	},
	{
		Id:             "2",
		NewCategoryDto: entities.NewCategoryDto{Name: "Books"},
	},
}

type CategoryRepositoryStub struct{}

func (*CategoryRepositoryStub) GetAll() ([]*entities.Category, error) {
	return mockCategories, nil
}

func MakeCategoriesRepoStub() repositories.ICategoryRepository {
	return &CategoryRepositoryStub{}
}
