package handlers

import (
	"backend/data"
	"github.com/gofiber/fiber/v3"
)

func JobHandler(c fiber.Ctx) error{
	return c.JSON(data.JobItems)
}