package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	UserID string
	Role string
	ExpiredAt float64 
}

func GetJWTFromContext(c *fiber.Ctx) (*JWT) {
		// Find id from jwt
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return &JWT{
		UserID: claims["id"].(string),
		Role: claims["role"].(string),
		ExpiredAt: claims["exp"].(float64),
	}
}