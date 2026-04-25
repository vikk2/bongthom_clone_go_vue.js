package main

import (
	"backend/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173"},
	}))

	routes.Linklistroute(app)
	routes.AnnouncementRoute(app)
	routes.ClassifiedRoute(app)
	routes.JobRoute(app)
	routes.OrganizationTypeRoute(app)
	routes.OtherServiceRoute(app)
	routes.RssFilesRoute(app)
	routes.SpecialScheduleRoute(app)

	app.Listen(":8080")
}
