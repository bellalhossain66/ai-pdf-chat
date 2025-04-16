package service

import (
	"fmt"
	"ai-pdf-chat/internal/repository"
	"ai-pdf-chat/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:username`
	jwt.StandardClaims
}

func LoginUser(username, password string)(string, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return "", fmt.Errorf("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("Incorrect password")
	}

	token, err := utils.GenerateJWT(&user)
	if err != nil {
		return "", err
	}

	return token, nil

}
