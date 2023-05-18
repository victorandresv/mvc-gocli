package controllers

import "net/http"

func HomeController(response http.ResponseWriter, request *http.Request) {

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Welcome"))
}
