package main

import (
	"fmt"
	"net/http"

	"github.com/tthawks/calculator/controllers"
)

func main() {
	controllers.RegisterControllers()

	fmt.Println("Running localhost server on port 3000")
	http.ListenAndServe(":3000", nil)
}
