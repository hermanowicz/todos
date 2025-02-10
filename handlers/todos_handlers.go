package handlers

import "github.com/gofiber/fiber/v2"

func GetTodos(ctx *fiber.Ctx) error {
	// todo: response with todos page scoped to user
	return ctx.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"Page": "Your Todos",
		},
	})
}
