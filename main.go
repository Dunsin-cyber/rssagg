package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main () {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	serve := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server started on port %v 🎉🐱‍🏍", portString)
	err := serve.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

}