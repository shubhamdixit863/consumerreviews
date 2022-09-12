package routes

import (
	"consumerreviews/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/home")

	// Routes for GET method:
	route.Get("/", controllers.BasePage) // get list of all books

}
