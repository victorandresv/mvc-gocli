package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"mvcgo/routes"
	"net/http"
)

func main() {
	port := "8900"
	router := mux.NewRouter()

	fmt.Println("Server running on port: " + port)
	err := http.ListenAndServe(":"+port, routes.GetRoutes(router))
	if err != nil {
		panic(err.Error())
	}
}
