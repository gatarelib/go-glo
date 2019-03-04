package main

import (
	"log"
	"os"

	"github.com/jackmcguire1/go-glo"
)

var token string

func init() {
	token = os.Getenv("TOKEN")
}

func main() {
	client := glo.NewClient(token)
	user, err := client.GetUser()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user.ID, user.Name, user.Username, user.Email)
}
