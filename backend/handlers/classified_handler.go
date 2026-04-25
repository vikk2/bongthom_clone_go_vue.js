package handlers

import (
	"backend/data"
	"github.com/gofiber/fiber/v3"
)

func ClassifiedHandler(c fiber.Ctx) error{
	return c.JSON(data.ClassifiedItems)
}