package handlers

import "github.com/gofiber/fiber/v2"

func GetRegisterPage(ctx *fiber.Ctx) error {
	// todo: setup csrf for main app
	// todo: add logic to return register page template with csrf token embedded
	return ctx.JSON(fiber.Map{"Page": "Register Page", "Status": "OK"})
}
