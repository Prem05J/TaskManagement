package Middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/taskManagement/Utils"
)

// JWT Middleware to validate JWT token
var jwtSecretKey = []byte(Utils.GetEnv("SECRET_KEY", "your-secret-key"))

func JwtMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return Utils.WriteErrorJson(c, fiber.StatusUnauthorized, "UnAuthorized")
	}

	tokenString := authHeader[len("Bearer "):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecretKey, nil
	})

	if err != nil {
		return Utils.WriteErrorJson(c, fiber.StatusUnauthorized, "Invalid Token")
	}

	c.Locals("users", token)
	return c.Next()
}
