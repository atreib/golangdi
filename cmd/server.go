package main

import (
	"fmt"
	"log"

	"github.com/atreib/golangdi/infra/libs"
	"github.com/atreib/golangdi/infra/repositories"
	"github.com/atreib/golangdi/services"
)

func main() {
	jwtEncrypter := &libs.JwtEncrypter{}
	userRepository := &repositories.UserRepository{}
	authenticationService := &services.AuthenticateService{
		UserRepository: userRepository,
		JwtEncrypter:   jwtEncrypter,
	}

	username := "admin"
	//password := "123456" // correct password
	password := "123" // wrong password
	user, err := authenticationService.Authenticate(username, password)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Authentication succeded:", user.User.Login)
}
