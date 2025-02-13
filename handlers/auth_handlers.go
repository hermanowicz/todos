package handlers

import "github.com/gofiber/fiber/v2"

func GetRegisterPage(ctx *fiber.Ctx) error {
	// todo: setup csrf for main app
	// todo: add logic to return register page template with csrf token embedded
	return ctx.Render("views/register", fiber.Map{
		"Title":       "Todos | Register",
		"Description": "Register for my todo app and save Your tasks for later.",
	}, "views/layout")
}

func GetLoginPage(ctx *fiber.Ctx) error {
	return ctx.Render("views/login", fiber.Map{
		"Title": "Todos | Login",
	}, "views/layout")
}

func CreateSession(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"Path":   ctx.Path(),
		"Origin": ctx.IP(),
		"Token":  ctx.Params("token", string("anon")),
	})
}
