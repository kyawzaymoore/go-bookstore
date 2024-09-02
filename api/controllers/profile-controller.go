package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	common "github.com/kzm/go-bookstore/api/common"
	errorcode "github.com/kzm/go-bookstore/api/helper"
	"github.com/kzm/go-bookstore/api/models/response"
	"github.com/kzm/go-bookstore/api/repository"
)

type ProfileController struct {
	UserRepo repository.UserRepository
}

func NewProfileController(userRepo repository.UserRepository) *ProfileController {
	return &ProfileController{UserRepo: userRepo}
}

// Profile godoc
// @Summary Get Profile
// @Description Get Profile
// @Tags Profile
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ResponseModel[response.ProfileResponse]
// @Security BearerAuth
// @Router /profile [get]
func (c *ProfileController) GetProfile(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]

	claims, err := common.GetClaimsFromToken(tokenString)
	if err != nil {
		fmt.Println("Error parsing token: %v", err)
	}

	// Access specific claim values
	userid := claims["userid"].(string)
	_, user := c.UserRepo.GetUserByID(userid)

	profile := response.ProfileResponse{
		Id: user.ID,
		Username: user.Username,
		Name: user.Name,
	}

	res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, profile)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
