package service

import (
	"fmt"
	"ai-pdf-chat/db/model"
	"ai-pdf-chat/internal/repository"
	"ai-pdf-chat/pkg/dto"
)


func GetFilesForUser(userID int) (*dto.GenericResponse[[]model.File], error) {
	files, err := repository.GetFileByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch files")
	}
	return dto.OK(&files), nil
}

func UploadFile(userID int, fileName string) (*dto.GenericResponse[model.File], error) {
	file := &model.File{
		UserID:   userID,
		Filename: fileName,
		IsProcessed: 0,
	}

	err := repository.CreateFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file")
	}

	return dto.OK(file), nil
}

func ProcessFile(FileName string, isProcessed int) (*dto.GenericResponse[string], error) {
	result := repository.UpdateFileProcessingStatus(FileName, isProcessed)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to process file")
	}

	if result.RowsAffected == 0 {
		return dto.OKWithMessage(&FileName, "No record found with filename", ), nil
	}
	return dto.OKWithMessage(&FileName, "File processing status updated successfully"), nil
}
