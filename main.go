package main

import (
	configs2 "consumerreviews/pkg/configs"
	middleware "consumerreviews/pkg/middlewares"
	"consumerreviews/pkg/routes"
	"consumerreviews/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	// Define Fiber config.
	config := configs2.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.

	routes.PublicRoutes(app) // Register a public routes for app.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
