package main

import (
	"fmt"

	"github.com/alephmelo/go-web/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Aleph",
		LastName:  "Melo",
	}
	fmt.Println(u)
}
