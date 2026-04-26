package handlers

import (
	"backend/data"

	"github.com/gofiber/fiber/v3"
)

func AnnouncementHandler(c fiber.Ctx) error{
	data := data.InitAnnouncementData()
	return c.JSON(data.GetAll())
}

