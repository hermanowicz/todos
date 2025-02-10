package handlers

import "github.com/gofiber/fiber/v2"

func GetAboutPage(ctx *fiber.Ctx) error {
	// todo: add logic to return about page template
	return ctx.Render("views/about", fiber.Map{"Title": "Todos | About Page", "Description": "About page for Todos app by TTH"}, "views/layout")
}
