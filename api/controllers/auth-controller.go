package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/kzm/go-bookstore/api/config"
	errorcode "github.com/kzm/go-bookstore/api/helper"
	"github.com/kzm/go-bookstore/api/models/request"
	"github.com/kzm/go-bookstore/api/models/response"
	"github.com/kzm/go-bookstore/api/repository"
	"github.com/kzm/go-bookstore/api/utils"
)

type AuthController struct {
	UserRepo repository.UserRepository
}

func NewAuthController(userRepo repository.UserRepository) *AuthController {
	return &AuthController{UserRepo: userRepo}
}

func createToken(userid string, username string, sessionToken string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userid":   userid,
			"username": username,
			"uuid":     sessionToken,
			"exp":      time.Now().Add(time.Minute * 10).Unix(),
		})
	secretKey := []byte(config.AppConfig.App.SecretKey)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// SignUp godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body request.UserRequest true "User registration details"
// @Success 200 {object} response.ResponseModel[models.User]
// @Security BasicAuth
// @Router /auth/signup [post]
func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	userRequest := &request.UserRequest{}
	utils.ParseBody(r, userRequest)

	if c.UserRepo.CheckUsernameExists(userRequest.Username) {
		res, _ := json.Marshal((response.GetResponseModel("E101", "")))
		w.Write(res)
		return
	}

	user, err := c.UserRepo.CreateUser(userRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(response.GetResponseModel(errorcode.E500, ""))
		w.Write(res)
		return
	}

	res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, user)))
	w.Write(res)
}

// SignIn godoc
// @Summary Login an existing user
// @Description Log in a user with username and password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body request.LoginRequest true "User login details"
// @Success 200 {object} response.ResponseModel[response.TokenResponse]
// @Security BasicAuth
// @Router /auth/signin [post]
func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	loginRequest := &request.LoginRequest{}
	utils.ParseBody(r, loginRequest)

	result, user := c.UserRepo.SignIn(loginRequest)
	if result {
		sessionToken := uuid.New().String()
		tokenString, err := createToken(user.ID, user.Username, sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		c.UserRepo.SaveUserSession(user, sessionToken)

		tokenResponse := response.TokenResponse{
			AccessToken: tokenString,
		}
		res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, tokenResponse)))
		w.Write(res)
		return
	}

	res, _ := json.Marshal((response.GetResponseModel(errorcode.E102, "")))
	w.Write(res)
}
