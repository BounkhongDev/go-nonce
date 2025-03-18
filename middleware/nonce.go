package middleware

import (
	"go-nonce/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Nonce expiration time
const nonceTTL = 20 * time.Second

// Generate a nonce and store it in Redis
func GenerateNonce(c *fiber.Ctx) error {
	nonce := uuid.New().String()

	// Store nonce in Redis
	err := database.RedisClient.Set(database.Ctx, nonce, true, nonceTTL).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate nonce"})
	}

	return c.JSON(fiber.Map{"nonce": nonce})
}

// Middleware to validate nonce
func ValidateNonce(c *fiber.Ctx) error {
	nonce := c.Get("X-Nonce") // Expect nonce in header
	if nonce == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing nonce"})
	}

	// Check if nonce exists in Redis
	exists, err := database.RedisClient.Exists(database.Ctx, nonce).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to validate nonce"})
	}
	if exists == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid nonce"})
	}

	return c.Next()
}
