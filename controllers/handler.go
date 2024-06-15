package controllers

import (
	"fmt"
	"net/http"
)

// FormHandler handles requests to the /form route
func FormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {

		fmt.Fprintf(w, "Prseform() err: %v", err)
		return

	}
	fmt.Fprintf(w, "post request successful \n")
	name := r.FormValue("Name")
	address := r.FormValue("Address")

	fmt.Fprintf(w, "name = %s \n", name)
	fmt.Fprintf(w, "address = %s \n", address)
}

// HelloHandler handles requests to the /hello route
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello")
}
