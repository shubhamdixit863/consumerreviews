package controllers

import "github.com/gofiber/fiber/v2"

func BasePage(c *fiber.Ctx) error {
	return c.Render("index", nil)
}
