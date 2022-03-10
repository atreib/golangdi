package main

import (
	"encoding/json"
	"fmt"
	"log"

	factories "github.com/atreib/golangdi/cmd/factories/category"
	"github.com/atreib/golangdi/presentation"
)

func main() {
	endpoint := factories.MakeList()
	result := endpoint.Handle(presentation.IRequest{})
	jsonData, err := json.Marshal(&result.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", result.Status)
	fmt.Println("Categories:")
	fmt.Println(string(jsonData))
}
