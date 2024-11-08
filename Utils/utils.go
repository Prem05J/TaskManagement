package Utils

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetEnv(key, fallBack string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallBack
}

func WriteErrorJson(c *fiber.Ctx, status int, desc string) error {
	return c.Status(status).JSON(fiber.Map{"error": desc})
}

func WriteJson(c *fiber.Ctx, status int, item interface{}) error {
	return c.Status(status).JSON(item)
}

func WriteFiberMap(c *fiber.Ctx, status int, key string, value any) error {
	return c.Status(status).JSON(fiber.Map{key: value})
}

func WriteMap(c *fiber.Ctx, status int, desc fiber.Map) error {
	return c.Status(status).JSON(desc)
}
