package services

import (
	"testing"

	"github.com/atreib/golangdi/infra/repositories"
	mocks "github.com/atreib/golangdi/infra/repositories/__mocks__"
)

func makeSut() (*ListCategoriesService, repositories.ICategoryRepository) {
	return &ListCategoriesService{}, mocks.MakeCategoriesRepoStub()
}

func ListCategories_HappyPath(t *testing.T) {
	t.Run("Should return the same result as the repository", func(t *testing.T) {
		sut, categoriesRepo := makeSut()
		got, _ := sut.List()
		want, _ := categoriesRepo.GetAll()
		if &got != &want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
