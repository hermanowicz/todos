package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hermanowicz/todos/models"
	"gorm.io/gorm"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	// starting db_conn & migrating models
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error when connecting to Sqlite, with error:", err.Error())
	}

	db.AutoMigrate(&models.Todos{}, &models.TodoLists{}, &models.Users{})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(models.DefaultResponse{
			ResponseBody: "Hello from api todos app, writen in go using chi freamwork.",
			Status:       http.StatusOK,
			YourIP:       r.RemoteAddr,
		})

		if err != nil {
			http.Error(w, "Json response faild with error loged for dev/ops team, sorry", http.StatusInternalServerError)
		}
	})

	r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		results := db.Find(&models.Todos{})
		fmt.Println(results)
	})

	r.Get("/todo-lists", func(w http.ResponseWriter, r *http.Request) {
		results := db.Find(&models.TodoLists{})
		fmt.Println(results)
	})

	r.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		var todo models.Todos
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "error while decoding req body", http.StatusInternalServerError)
		}

		result := db.Create(&todo)
		json.NewEncoder(w).Encode(model.DefaultResponse{
			Status
		})
	})

	s := &http.Server{
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      r,
		Addr:         ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
