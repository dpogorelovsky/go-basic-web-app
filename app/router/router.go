package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetRouter - routes
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", health)

	return r
}

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
