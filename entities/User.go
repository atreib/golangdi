package entities

type User struct {
	Id    int
	Login string
	Email string
}

type Authentication struct {
	User User
	Jwt  string
}
