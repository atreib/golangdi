package services

import "github.com/atreib/golangdi/entities"

type IListCategoriesService interface {
	List() ([]*entities.Category, error)
}
