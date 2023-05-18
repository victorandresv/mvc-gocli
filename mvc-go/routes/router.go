package routes

import (
	"github.com/gorilla/mux"
	"mvcgo/controllers"
)

func GetRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.HomeController)

	return router
}
