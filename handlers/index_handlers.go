package handlers

import "github.com/gofiber/fiber/v2"

func GetIndex(ctx *fiber.Ctx) error {
	return ctx.Render(
		"views/index", fiber.Map{
			"Title": "Todos | Welcome",
		}, "views/layout",
	)
}
