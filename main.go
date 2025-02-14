package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/hermanowicz/todos/handlers"
	"github.com/hermanowicz/todos/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	port = "8080"
)

//go:embed views/*
var viewsfs embed.FS

func main() {

	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error when opening sqlite", err.Error())
	}

	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	if err != nil {
		log.Fatal("Error when opening sqlite", err.Error())
	}

	db.AutoMigrate(models.Todo{})

	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
		Prefork:               false,
		CaseSensitive:         false,
		Views:                 engine,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", handlers.GetIndex)
	app.Get("/about", handlers.GetAboutPage)
	app.Get("/privacy", handlers.GetPrivacyPolicy)
	app.Get("/newsletter-policy", handlers.GetNewsletterPolicy)
	app.Get("/register", handlers.GetRegisterPage)
	app.Get("/login", handlers.GetLoginPage)
	app.Get("/new-session/:token", handlers.CreateSession)

	// main todos logic
	app.Get("/todos", handlers.GetTodos)

	err = app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Faild to start application with error: \n", err.Error())
	}
}
