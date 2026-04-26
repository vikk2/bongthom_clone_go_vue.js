package handlers

import (
	"backend/data"
	"github.com/gofiber/fiber/v3"
)

func RssFileHandler(c fiber.Ctx) error{
	data := data.InitRssFileData()
	return c.JSON(data.GetAll())
}