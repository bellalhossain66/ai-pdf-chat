package service

import (
	"fmt"
	// "ai-pdf-chat/config"
	"ai-pdf-chat/db/model"
	"ai-pdf-chat/internal/repository"
	"ai-pdf-chat/pkg/dto"
	// "ai-pdf-chat/api/request"
)

func GetChatList(fileId int, userId int, page int, limit int) (*dto.GenericResponse[[]model.Chat], error) {
	chats, err := repository.GetChatsByFile(fileId, userId, page, limit)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch chat")
	}
	// fmt.Println(chats)
	return dto.OK(&chats), nil
}

func AskQuestion(fileId int, question string, answer string, userId int) (*dto.GenericResponse[model.Chat], error) {
	chat := &model.Chat{
		UserId:   userId,
		FileId:   fileId,
		Question: question,
		Answer:   answer,
	}

	if err := repository.SaveChat(chat); err != nil {
		return nil, err
	}

	return dto.OK(chat), nil
}
