package handler

import (
	"ai-pdf-chat/api/request"
	"ai-pdf-chat/internal/service"
	"ai-pdf-chat/middleware"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"ai-pdf-chat/pkg/dto"
	// "ai-pdf-chat/pkg/utils"
	// "net/http"
	"fmt"
)

func GetChats(c *fiber.Ctx) error {
	fileId, _ := strconv.Atoi(c.Query("file_id"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	userID, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.Status(401).JSON(dto.ErrorWithMessage[string](nil, err.Error()))
	}
	// fmt.Println(fileId, userID, page, limit)

	chats, err := service.GetChatList(fileId, userID, page, limit)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorWithMessage[string](nil, "Failed to fetch chat history"))
	}

	return c.JSON(chats)
}

func AskQuestion(c *fiber.Ctx) error {
	var req request.AskQuestionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(dto.ErrorWithMessage[string](nil, "Invalid request"))
	}

	userID, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.Status(401).JSON(dto.ErrorWithMessage[string](nil, err.Error()))
	}

	// var client = &http.Client{
	// 	Timeout: 0 * time.Second,
	// }

	// var baseURL = fmt.Sprintf("%s/%s", utils.PythonUrl, "question")

	// var params = url.Values{}
	// params.Add("q", req.Question)
	// params.Add("f", strconv.Itoa(req.FileId))

	// var fullURL = fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// if request, err = http.NewRequest("GET", fullURL, nil); err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	// if response, err = client.Do(request); err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	// defer response.Body.Close()
	// if response.StatusCode == http.StatusOK {
	// 	var chatResponse request.ChatResponse

	// 	_ = json.NewDecoder(response.Body).Decode(&chatResponse)

	// 	resp, err := service.AskQuestion(req.FileId, req.Question, chatResponse.response, userID)
	// 	if err != nil {
	// 		return c.Status(500).JSON(dto.ErrorWithMessage[string](nil, "failed to ask question"))
	// 	}

	// 	return c.JSON(resp)
	// } else {
	// 	return nil, fmt.Errorf("failed to generate answer, status code: %d", response.StatusCode)
	// }

	answer := fmt.Sprintf("Pretend AI answer to: \"%s\"", req.Question)
	resp, err := service.AskQuestion(req.FileId, req.Question, answer, userID)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorWithMessage[string](nil, "failed to ask question"))
	}

	return c.JSON(resp)
}

