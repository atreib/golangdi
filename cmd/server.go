package main

import (
	"fmt"
	"log"

	"github.com/atreib/golangdi/infra/repositories"
	services "github.com/atreib/golangdi/services/category"
)

func main() {
	categoriesRepo := &repositories.InMemoryCategoryRepository{}
	listCategoriesService := &services.ListCategoriesService{
		CategoryRepository: categoriesRepo,
	}

	categories, err := listCategoriesService.List()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Categories:")
	for _, v := range categories {
		fmt.Println("-", v.Name)
	}
}
