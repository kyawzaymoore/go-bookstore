package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/kzm/go-bookstore/api/config"
	"github.com/kzm/go-bookstore/api/models"
	"github.com/kzm/go-bookstore/api/models/request"
	"github.com/kzm/go-bookstore/api/utils"
)

type UserRepository interface {
	CheckUsernameExists(username string) bool
	CreateUser(data *request.UserRequest) (*models.User, error)
	SignIn(data *request.LoginRequest) (bool, *models.User)
	SaveUserSession(data *models.User, sessionToken string) (*models.UserSession, error)
	GetUserByID(userid string) (bool, *models.User)
	CheckSession(userid string , sessionToken string) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{db: config.GetDB()}
}

func (r *userRepository) CheckUsernameExists(username string) bool {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	return result.Error == nil
}

func (r *userRepository) CreateUser(data *request.UserRequest) (*models.User, error) {
	hashedPassword := utils.HashString(data.Password)
	user := models.User{
		Username: data.Username,
		Password: hashedPassword,
		Name:     data.Name,
	}
	user.ID = uuid.New().String()
	user.RegisterDate = time.Now()

	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) SignIn(data *request.LoginRequest) (bool, *models.User) {
	var user models.User
	result := r.db.Where("username = ?", data.Username).First(&user)

	if result.Error == nil {
		userPassword := utils.HashString(data.Password)

		if user.Password == userPassword {
			return true, &user
		}
	}

	return false, nil
}

func (r *userRepository) SaveUserSession(data *models.User, sessionToken string) (*models.UserSession, error) {
	userToken := models.UserSession{
		Token:  sessionToken,
		UserID: data.ID,
	}
	//Delete Existing Session for Single Device Login
	r.db.Where("user_id = ?", data.ID).Delete(&models.UserSession{})

	userToken.ID = uuid.New().String()
	userToken.CreatedDate = time.Now()
	
	result := r.db.Create(userToken)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userToken, nil
}

func (r *userRepository) GetUserByID(userid string) (bool, *models.User) {
	var user models.User
	result := r.db.Where("id = ?", userid).First(&user)
	return result.Error == nil, &user
}

func (r *userRepository) CheckSession(userid string , sessionToken string) bool {
	var userSession models.UserSession
	result := r.db.Where("user_id = ?", userid).First(&userSession)
	if result.Error != nil {
		// If no record is found, return false
		return false
	}
	if userSession.Token == sessionToken {
		return true
	}
	return false
}
