package reponder

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON200 - success response
func JSON200(w http.ResponseWriter, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not encode response, data: %v", data)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
