package routes

import (
	"github.com/gorilla/mux"
	"github.com/kzm/go-bookstore/api/controllers"
	"github.com/kzm/go-bookstore/api/middlewares"
	"github.com/kzm/go-bookstore/api/repository"
)

var BookRoutes = func (router *mux.Router)  {
	// Initialize repository and controller
	bookRepo := repository.NewBookRepository()
	bookController := controllers.NewBookController(bookRepo)

	userRepo := repository.NewUserRepository()

	api := router.PathPrefix("/book").Subrouter()
	api.Use(middlewares.BearerTokenMiddleware(userRepo))

	
	api.HandleFunc("", bookController.GetBook).Methods("GET")
	api.HandleFunc("", bookController.CreateBook).Methods("POST")
	api.HandleFunc("/{bookId}", bookController.GetBookById).Methods("GET")
	api.HandleFunc("/{bookId}", bookController.UpdateBook).Methods("PUT")
	api.HandleFunc("/{bookId}", bookController.DeleteBook).Methods("DELETE")

	// Add Basic Auth middleware to a route
    //router.Handle("/book", middlewares.BasicAuthMiddleware(http.HandlerFunc(controllers.GetBook))).Methods("GET")
    
}
