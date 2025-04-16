package utils

import(
	"github.com/dgrijalva/jwt-go"
	"time"
	"os"
	"ai-pdf-chat/db/model"
)

type Claims struct {
	Username string `json:username`
	UserId int `json:id`
	jwt.StandardClaims
}

func GenerateJWT(user *model.User) (string, error) {
	claims := Claims{
		Username: user.Username,
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
			Issuer: "ai-pdf-chat",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}