package repositories

import "github.com/atreib/golangdi/entities"

type ICategoryRepository interface {
	GetAll() ([]*entities.Category, error)
}
