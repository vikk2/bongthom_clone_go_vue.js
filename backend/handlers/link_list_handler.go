package handlers

import (
	"backend/data"
	"github.com/gofiber/fiber/v3"
)

func LinkListHandler(c fiber.Ctx) error{
	return c.JSON(data.LinkListItems)
}
