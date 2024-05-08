package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/subinoybiswas/devchat/routes"
)

func main() {
	fmt.Println("Gopher is Running!")
	routes.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
