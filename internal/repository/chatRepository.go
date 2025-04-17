package repository

import (
	"ai-pdf-chat/db/model"
	"ai-pdf-chat/config"
)

func SaveChat(chat *model.Chat) error {
	return config.DB.Create(chat).Error
}

func GetChatsByFile(fileId int, userId int, page int, limit int) ([]model.Chat, error) {
	var chats []model.Chat
	offset := (page - 1) * limit
	result := config.DB.
		Where("file_id = ? AND user_id = ?", fileId, userId).
		Offset(offset).
		Limit(limit).
		Order("id DESC").
		Find(&chats)
	return chats, result.Error
}
