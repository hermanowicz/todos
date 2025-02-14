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

func CreateTodo(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World! -> Create Todo")
}

func UpdateTodo(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World! -> Update Todo")
}

func DeleteTodo(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World! -> Delete Todo")
}
