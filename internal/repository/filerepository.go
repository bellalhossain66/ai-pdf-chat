package repository

import (
	"ai-pdf-chat/config"
	"ai-pdf-chat/db/model"
	"gorm.io/gorm"
)

func GetFileByUserID(userID int) ([]model.File, error) {
	var files []model.File
	result := config.DB.Where("user_id = ?", userID).Find(&files)
	return files, result.Error
}

func CreateFile(file *model.File) error {
	return config.DB.Create(file).Error
}

func UpdateFileProcessingStatus(FileName string, isProcessed int) *gorm.DB {
	return config.DB.Model(&model.File{}).Where("filename = ?", FileName).Update("is_processed", isProcessed)
}
