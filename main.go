package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// TODO: put this in environment variables!!!!
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Chicago"
	db, db_err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if db_err != nil {
		log.Fatal(db_err)
	}

	println(fmt.Sprintf("Database connected: %v", db))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API"))
	})
	println("Go-Fast is running on port 8888")
	err := http.ListenAndServe(":8888", router)
	if err != nil {
		log.Fatal(err)
	}
}
