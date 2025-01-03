package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/yourusername/football-chat-backend/config"
	"github.com/yourusername/football-chat-backend/models"
)

// CreateUser registers a new user
func CreateUser(email, password string) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
		// Nickname can be defaulted, or you can ask for it in signup
		Nickname: email, // just as a placeholder
	}

	// Save to DB
	result := config.Config.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

// LoginUser validates credentials and returns a JWT token
func LoginUser(email, password string) (string, error) {
	var user models.User
	result := config.Config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return "", errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Config.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
