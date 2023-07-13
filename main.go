package main

import (
	"BookStore/configs"
	"BookStore/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	// Home
	r.Get("/", handlers.HomeGet)
	// Books Routers
	r.Post("/books", handlers.BookCreate)
	r.Put("/books/{id}", handlers.BookUpdate)
	r.Delete("/books/{id}", handlers.BookDelete)
	r.Get("/books", handlers.BookList)
	r.Get("/books/{id}", handlers.BookGet)

	// Authors Routers
	r.Get("/authors", handlers.AuthorList)
	r.Get("/authors/{id}", handlers.AuthorGet)
	r.Get("/authors/name/{authorname}", handlers.AuthorGetByName)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
