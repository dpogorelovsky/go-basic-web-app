package reponder

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON200 - success response
func JSON200(w http.ResponseWriter, data interface{}) {
	renderJSON(w, http.StatusOK, data)
}

// JSON200OK - default OK message
func JSON200OK(w http.ResponseWriter) {
	renderJSON(w, http.StatusOK, map[string]string{"result": "ok"})
}

// JSON500 - responds with 500 error
func JSON500(w http.ResponseWriter) {
	data := map[string]string{"error": "internal server error"}
	renderJSON(w, http.StatusInternalServerError, data)
}

// JSON404 - responds with 404 error
func JSON404(w http.ResponseWriter) {
	data := map[string]string{"error": "resource not found"}
	renderJSON(w, http.StatusNotFound, data)
}

// JSON400 - responds with 400 error
func JSON400(w http.ResponseWriter) {
	data := map[string]string{"error": "Bad request"}
	renderJSON(w, http.StatusBadRequest, data)
}

//renderJSON generic json response
func renderJSON(w http.ResponseWriter, status int, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not encode response, data: %v", data)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)
}
