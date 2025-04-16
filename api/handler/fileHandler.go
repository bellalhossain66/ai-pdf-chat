package handler

import (
	"fmt"
	"ai-pdf-chat/api/request"
	"ai-pdf-chat/internal/service"
	"ai-pdf-chat/middleware"
	"github.com/gofiber/fiber/v2"
	"ai-pdf-chat/pkg/dto"
	"os"
)

func ListFiles(c *fiber.Ctx) error {
	userID, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.Status(401).JSON(dto.ErrorWithMessage[string](nil, err.Error()))
	}
	
	resp, err := service.GetFilesForUser(userID)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorWithMessage[string](nil, err.Error()))
	}
	return c.JSON(resp)
}

func UploadFile(c *fiber.Ctx) error {
	var body request.FileUploadRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(dto.ErrorWithMessage[string](nil, "Invalid input"))
	}

	userID, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.Status(401).JSON(dto.ErrorWithMessage[string](nil, err.Error()))
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(dto.ErrorWithMessage[string](nil, "No file provided"))
	}

	if fileHeader.Header.Get("Content-Type") != "application/pdf" {
		return c.Status(400).JSON(dto.ErrorWithMessage[string](nil, "Only PDFs are allowed"))
	}

	savePath := fmt.Sprintf("./uploads/%s", fileHeader.Filename)
	// fmt.Println(savePath)
	if err := c.SaveFile(fileHeader, savePath); err != nil {
		return c.Status(500).JSON(dto.ErrorWithMessage[string](nil, "Failed to save file"))
	}

	resp, err := service.UploadFile(userID, fileHeader.Filename)
	if err != nil {
		os.Remove(savePath)
		return c.Status(500).JSON(dto.ErrorWithMessage[string](nil, err.Error()))
	}

	return c.JSON(resp)
}

func ProcessFile(c *fiber.Ctx) error {
	var body request.FileProcessRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(dto.ErrorWithMessage[string](nil, "Invalid input"))
	}

	isProcessed := 0
	if body.IsProcessed {
		isProcessed = 1
	}

	// Call service to process file
	resp, err := service.ProcessFile(body.FileName, isProcessed)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorWithMessage[string](nil, err.Error()))
	}

	return c.JSON(resp)
}
