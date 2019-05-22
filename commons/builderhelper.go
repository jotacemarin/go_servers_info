package commons

import (
	"encoding/json"
	"net/http"
)

// BuilderJSON : builder json
func BuilderJSON(writer http.ResponseWriter, success bool, httpStatusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil || !success {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpStatusCode)
	writer.Write(response)
}
