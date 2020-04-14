package main

import (
	"fmt"
	"net/http"

	"github.com/alephmelo/go-web/controllers"
)

func main() {
	port := ":3000"
	controllers.RegisterControllers()

	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}
