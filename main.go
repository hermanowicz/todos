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

		var allTodos []defs.Todo
		rows, err := db.Query(`select user, todo_title, todo_body, status from todos`)
		if err != nil {
			http.Error(w, "error logged for admins.", http.StatusInternalServerError)
		}
		defer rows.Close()

		for rows.Next() {
			var todo defs.Todo
			rows.Scan(&todo.User, &todo.TodoTitle, &todo.TodoBody, &todo.Status)
			allTodos = append(allTodos, todo)
		}
		err = rows.Err()
		if err != nil {
			http.Error(w, "error logged for admins.", http.StatusInternalServerError)
		}

		_ = json.NewEncoder(w).Encode(allTodos)
	})

	r.Get("/todo-lists", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("some day, maybe."))
	})

	r.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		var newTodo defs.Todo
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(w, "error wile ecoding data", http.StatusInternalServerError)
		}
		_, err = db.Exec(`insert into todos(user, todo_title, todo_body) values (?, ?, ?)`, &newTodo.User, &newTodo.TodoTitle, &newTodo.TodoBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "loclahost:8080/todos", http.StatusTemporaryRedirect)
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
