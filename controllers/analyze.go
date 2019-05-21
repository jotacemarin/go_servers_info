package analyzecontroller

import (
	"encoding/json"
	"net/http"

	"../models"
)

// Index : funct
func Index(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(models.Server{"Hola", "mundo", "como", "estas"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
