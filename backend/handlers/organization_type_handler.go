package handlers

import (
	"backend/data"
	"github.com/gofiber/fiber/v3"
)

func OrganizationTypeHandler(c fiber.Ctx) error {
	data := data.InitOrganizationTypeData()
	return c.JSON(data.GetAll())
}
