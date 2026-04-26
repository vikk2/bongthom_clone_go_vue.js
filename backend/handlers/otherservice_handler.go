package handlers

import (
	"backend/data"
	"github.com/gofiber/fiber/v3"
)

func OtherServiceHandler(c fiber.Ctx) error{
	data := data.InitOtherServiceData()
	return c.JSON(data.GetAll())
}