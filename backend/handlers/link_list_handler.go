package handlers

import (
	"backend/data"
	"github.com/gofiber/fiber/v3"
)

func LinkListHandler(c fiber.Ctx) error{
	data := data.InitLinkListData()
	return c.JSON(data.GetAll())
}
