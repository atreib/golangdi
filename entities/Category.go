package entities

type NewCategoryDto struct {
	Name string
}

type Category struct {
	Id string
	NewCategoryDto
}
