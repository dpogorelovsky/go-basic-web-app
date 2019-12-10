package router

import (
	"encoding/json"
	"net/http"

	"github.com/dpogorelovsky/go-basic-web-app/app/handlers"
	"github.com/dpogorelovsky/go-basic-web-app/app/middlewares"
	"github.com/dpogorelovsky/go-basic-web-app/app/storage/astorage"
	"github.com/gorilla/mux"
)

// GetRouter - routes
func GetRouter(db astorage.Storage) *mux.Router {
	r := mux.NewRouter()

	// commmon routes
	r.Handle("/", middlewares.Cacheable(http.HandlerFunc(health), "bugaga"))

	// users routes
	users := handlers.User{
		Storage: db,
	}
	r.HandleFunc("/users", users.Create).Methods("POST")
	r.HandleFunc("/users", users.List).Methods("GET")
	r.HandleFunc("/users/{id}", users.ID).Methods("GET")
	r.HandleFunc("/users/{id}", users.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", users.Delete).Methods("DELETE")

	// cities routes
	cities := handlers.City{
		Storage: db,
	}
	r.HandleFunc("/cities", cities.Create).Methods("POST")
	r.HandleFunc("/cities", cities.List).Methods("GET")
	r.HandleFunc("/cities/{id}", cities.ID).Methods("GET")
	r.HandleFunc("/cities/{id}", cities.Update).Methods("PUT")
	r.HandleFunc("/cities/{id}", cities.Delete).Methods("DELETE")

	return r
}

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
