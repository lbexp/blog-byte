package jwt_utils

import (
	"blog-byte/app/entity"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Id    int
	Email string
	jwt.MapClaims
}

var tokenExpiration = time.Now().Add(time.Hour * 72).Unix()

func GenerateToken(user entity.User) (string, error) {
	key := []byte(os.Getenv("JWT_KEY"))

	newJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   tokenExpiration,
	})

	token, err := newJwt.SignedString(key)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetJwtClaims(ctx *fiber.Ctx) (JwtClaims, error) {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(JwtClaims)

	return claims, nil
}
