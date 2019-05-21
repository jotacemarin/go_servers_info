package analyzecontroller

import (
	"encoding/json"
	"net/http"

	"../models"
)

// Index : funct
func Index(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	response, err := json.Marshal(models.Server{"Hola", host, "como", "estas"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
