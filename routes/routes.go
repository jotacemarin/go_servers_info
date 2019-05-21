package routes

import (
	analyzeroutes "./v1"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Router : main of all application
func Router() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/api", func(r chi.Router) {
		r.Mount("/v1/analyze", analyzeroutes.Routes())
	})

	return router
}
