package routes

import (
	"net/http"

	"example.com/myproject/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/form", controllers.FormHandler)
	http.HandleFunc("/hello", controllers.HelloHandler)
}
