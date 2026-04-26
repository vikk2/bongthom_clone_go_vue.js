package handlers

import (
	"backend/data"
	"github.com/gofiber/fiber/v3"
)

func SpecialScheduleHandler(c fiber.Ctx) error{
	data := data.InitSpecialScheduleData()
	return c.JSON(data.GetAll())
}