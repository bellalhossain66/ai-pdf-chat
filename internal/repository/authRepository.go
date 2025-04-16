package repository

import(
	"ai-pdf-chat/db/model"
	"ai-pdf-chat/config"
)

func GetUserByUsername(username string)(model.User, error) {
	var user model.User
	result := config.DB.Where("username = ?",username).First(&user)
	return user, result.Error
}