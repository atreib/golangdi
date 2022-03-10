package factories

import (
	repositories "github.com/atreib/golangdi/infra/repositories/category"
	"github.com/atreib/golangdi/presentation"
	endpoints "github.com/atreib/golangdi/presentation/endpoints/category"
	services "github.com/atreib/golangdi/services/category"
)

func MakeList() presentation.IEndpoint {
	repository := &repositories.InMemoryCategoryRepository{}
	service := &services.ListCategoriesService{
		CategoryRepository: repository,
	}
	endpoint := &endpoints.ListCategoriesEndpoint{
		ListCategoriesService: service,
	}
	return endpoint
}
