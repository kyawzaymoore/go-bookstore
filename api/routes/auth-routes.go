package routes

import (
	"github.com/gorilla/mux"
	"github.com/kzm/go-bookstore/api/controllers"
	"github.com/kzm/go-bookstore/api/middlewares"
	"github.com/kzm/go-bookstore/api/repository"
)

var AuthRoutes = func (router *mux.Router)  {
	// Initialize repository and controller
	userRepo := repository.NewUserRepository()
	authController := controllers.NewAuthController(userRepo)

	api := router.PathPrefix("/auth").Subrouter()
	api.Use(middlewares.BasicAuthMiddleware)

	// @Summary SignIn user
	// @Description Login user with credentials
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param login body request.LoginRequest true "Login credentials"
	// @Success 200 {object} response.TokenResponse
	// @Failure 400 {object} response.ErrorResponse
	// @Router /auth/signin [post]
	api.HandleFunc("/signin", authController.SignIn).Methods("POST")
	api.HandleFunc("/signup", authController.SignUp).Methods("POST")
}
