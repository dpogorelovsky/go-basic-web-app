package router

import (
	"encoding/json"
	"net/http"

	"github.com/dpogorelovsky/go-basic-web-app/app/handlers"
	"github.com/gorilla/mux"
)

// GetRouter - routes
func GetRouter() *mux.Router {
	r := mux.NewRouter()

	// commmon routes
	r.HandleFunc("/", health)

	// users routes
	users := handlers.User{}
	r.HandleFunc("/users", users.List)

	// cities routes
	cities := handlers.City{}
	r.HandleFunc("/cities", cities.List)

	return r
}

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
