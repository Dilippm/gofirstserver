package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/myproject/routes"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Setup routes
	routes.SetupRoutes()
	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
