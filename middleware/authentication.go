package middleware

import (
	"os"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
	"ai-pdf-chat/pkg/dto"
)

func JWTAuthentication(c *fiber.Ctx) error {
	tokenStr := c.Get("Authorization")
	if tokenStr == "" {
		return c.Status(401).JSON(dto.ErrorWithMessage[string](nil, "Missing token"))
	}

	tokenStr = strings.Split(tokenStr, "Bearer ")[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(dto.ErrorWithMessage[string](nil, "Invalid token"))
	}

	// claims, ok := token.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return c.Status(401).JSON(dto.ErrorWithMessage[string](nil, "Invalid token claims"))
	// }

	// userId := int(claims["UserId"].(float64))
	// c.Locals("UserId", userId)

	return c.Next()
}

// func GetUserIDFromJWT(c *fiber.Ctx) int {
// 	return c.Locals("UserId").(int)
// }

func ExtractUserID(c *fiber.Ctx) (int, error) {
	tokenStr := c.Get("Authorization")
	if tokenStr == "" {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Authorization token required")
	}

	tokenStr = strings.Split(tokenStr, "Bearer ")[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["UserId"].(float64))
	return userID, nil
}