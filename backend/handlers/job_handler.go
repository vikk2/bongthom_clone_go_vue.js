package handlers

import (
	"backend/data"
	"github.com/gofiber/fiber/v3"
)

func JobHandler(c fiber.Ctx) error{
	data := data.InitJobData()
	return c.JSON(data.GetAll())
}