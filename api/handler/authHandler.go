package handler

import(
	"github.com/gofiber/fiber/v2"
	"ai-pdf-chat/api/request"
	"ai-pdf-chat/internal/service"
	"ai-pdf-chat/pkg/dto"
)

func LoginUser(c *fiber.Ctx) error {
	var body request.LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(dto.ErrorWithMessage[string](nil, "Invalid request"))
	}

	token, err := service.LoginUser(body.Username, body.Password);
	if err != nil {
		return c.Status(401).JSON(dto.ErrorWithMessage[string](nil, err.Error()))
	}

	return c.JSON(dto.OKWithMessage[string](&token, "Login successful"))
}