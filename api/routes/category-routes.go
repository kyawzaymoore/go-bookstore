package routes

import (
	"github.com/gorilla/mux"
	"github.com/kzm/go-bookstore/api/controllers"
	"github.com/kzm/go-bookstore/api/middlewares"
	"github.com/kzm/go-bookstore/api/repository"
)

var CategoryRoutes = func (router *mux.Router)  {
	// Initialize repository and controller
	categoryRepo := repository.NewCategoryRepository()
	categoryController := controllers.NewCategoryController(categoryRepo)

	userRepo := repository.NewUserRepository()

	api := router.PathPrefix("/category").Subrouter()
	api.Use(middlewares.BearerTokenMiddleware(userRepo))

	
	api.HandleFunc("", categoryController.GetCategory).Methods("GET")
	api.HandleFunc("", categoryController.CreateCategory).Methods("POST")
	api.HandleFunc("/{categoryId}", categoryController.GetCategoryById).Methods("GET")
	api.HandleFunc("/{categoryId}", categoryController.UpdateCategory).Methods("PUT")
	api.HandleFunc("/{categoryId}", categoryController.DeleteCategory).Methods("DELETE")

	// Add Basic Auth middleware to a route
    //router.Handle("/category", middlewares.BasicAuthMiddleware(http.HandlerFunc(controllers.GetCategory))).Methods("GET")
    
}
