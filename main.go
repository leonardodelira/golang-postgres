package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leonardodelira/go-postgres/configs"
	"github.com/leonardodelira/go-postgres/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	
	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.List)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}