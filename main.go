package main

import (
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/hermanowicz/todos/handlers"
	"log"
	"net/http"
)

var (
	port = "8080"
)

//go:embed views/*
var viewsfs embed.FS

func main() {
	log.Println("Starting todos app on port:", port)

	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               false,
		CaseSensitive:         false,
		Views:                 engine,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", handlers.GetIndex)

	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Faild to start application with error: \n", err.Error())
	}
}
