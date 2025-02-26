package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hermanowicz/todos/defs"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	// db conn setup
	db, err := sqlx.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalln("Error whene opening conn to sqlite:", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("error whene pinging db")
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(defs.DefaultResponse{
			ResponseBody: "Hello from api todos app, writen in go using chi freamwork.",
			Status:       http.StatusOK,
			YourIP:       r.RemoteAddr,
		})

		if err != nil {
			http.Error(w, "Json response faild with error loged for dev/ops team, sorry", http.StatusInternalServerError)
		}
	})

	r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not implemented.", http.StatusNotImplemented)
	})

	r.Get("/todo-lists", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not implemented.", http.StatusNotImplemented)

	})

	r.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not implemented.", http.StatusNotImplemented)
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
