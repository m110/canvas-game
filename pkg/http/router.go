package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	r.Use(cors.Handler)

	return r
}
