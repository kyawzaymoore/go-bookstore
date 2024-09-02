package routes

import (
	"github.com/gorilla/mux"
	"github.com/kzm/go-bookstore/api/controllers"
	"github.com/kzm/go-bookstore/api/middlewares"
	"github.com/kzm/go-bookstore/api/repository"
)

var ProfileRoutes = func(router *mux.Router) {
	// Initialize repository and controller
	userRepo := repository.NewUserRepository()
	profileController := controllers.NewProfileController(userRepo)

	api := router.PathPrefix("/profile").Subrouter()
	api.Use(middlewares.BearerTokenMiddleware(userRepo))

	api.HandleFunc("", profileController.GetProfile).Methods("GET")
}
