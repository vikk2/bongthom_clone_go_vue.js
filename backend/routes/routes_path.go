package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v3"
)

func Linklistroute(app *fiber.App) {
	app.Get("/api/link-list", handlers.LinkListHandler)
}

func AnnouncementRoute(app *fiber.App) {
	app.Get("/api/announcement-list", handlers.AnnouncementHandler)
}

func ClassifiedRoute(app *fiber.App) {
	app.Get("/api/classified-list", handlers.ClassifiedHandler)
}

func JobRoute(app *fiber.App) {
	app.Get("/api/job-link", handlers.JobHandler)
}

func OrganizationTypeRoute(app *fiber.App) {
	app.Get("/api/organization-type", handlers.OrganizationTypeHandler)
}

func OtherServiceRoute(app *fiber.App) {
	app.Get("/api/other-service-link", handlers.OtherServiceHandler)
}

func RssFilesRoute(app *fiber.App) {
	app.Get("/api/rssfiles-link", handlers.RssFileHandler)
}

func SpecialScheduleRoute(app *fiber.App){
	app.Get("/api/specialSched-link", handlers.SpecialScheduleHandler)
}