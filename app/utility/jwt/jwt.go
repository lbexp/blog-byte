package jwt_utils

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Sub   int
	Email string
}

func GenerateToken(claims JwtClaims) (string, error) {
	key := []byte(os.Getenv("JWT_KEY"))

	newJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   claims.Sub,
		"email": claims.Email,
	})

	token, err := newJwt.SignedString(key)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetJwtClaims(ctx *fiber.Ctx) (JwtClaims, error) {
	// TODO: Add logic to get jwt claims
	return JwtClaims{
		Sub:   1,
		Email: "dummy",
	}, nil
}
