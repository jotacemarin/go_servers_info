package analyzeroutes

import (
	analyzecontroller "../../../controllers"
	"github.com/go-chi/chi"
)

// Routes for analyze endpoint
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/analyze", analyzecontroller.Analyze)
	return router
}
