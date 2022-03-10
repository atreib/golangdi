package endpoints

import (
	"github.com/atreib/golangdi/presentation"
	services "github.com/atreib/golangdi/services/category"
)

type ListCategoriesEndpoint struct {
	ListCategoriesService services.IListCategoriesService
}

func (e *ListCategoriesEndpoint) Handle(presentation.IRequest) *presentation.IResponse {
	categories, err := e.ListCategoriesService.List()

	if err != nil {
		return &presentation.IResponse{
			Status: 500,
			Body:   err,
		}
	}

	return &presentation.IResponse{
		Status: 200,
		Body:   categories,
	}
}
