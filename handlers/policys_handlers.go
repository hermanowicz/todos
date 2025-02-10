package handlers

import "github.com/gofiber/fiber/v2"

func GetPrivacyPolicy(ctx *fiber.Ctx) error {
	// temp solution todo: add page handler returning html template
	return ctx.Render("views/privacy_policy", fiber.Map{"Title": "Todos | Privacy Policy"}, "views/layout")
}

func GetNewsletterPolicy(ctx *fiber.Ctx) error {
	return ctx.Render("views/newsletter_policy", fiber.Map{"Title": "Todos | Newsletter Policy"}, "views/layout")
}
