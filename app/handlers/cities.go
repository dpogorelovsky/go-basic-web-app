package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/dpogorelovsky/go-basic-web-app/app/storage/astorage"
	R "github.com/dpogorelovsky/go-basic-web-app/app/utils/responder"
	"github.com/dpogorelovsky/go-basic-web-app/app/utils/utils"
	"github.com/dpogorelovsky/go-basic-web-app/app/utils/validator"
	"github.com/gorilla/mux"
)

// City - cities handler
type City struct {
	Storage astorage.CityStorage
}

// List - returning list of cities
func (h *City) List(w http.ResponseWriter, r *http.Request) {
	limit, offset := utils.GetPaginationParams(r.URL.Query())
	resp, err := h.Storage.GetCityList(limit, offset)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	if len(resp) < 1 {
		R.JSON404(w)
		return
	}

	R.JSON200(w, resp)
}

// Create - creates new city
func (h *City) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	city, err := validator.CityCreate(body)
	if err != nil {
		log.Println(err)
		R.JSON400(w)
		return
	}

	err = h.Storage.CreateCity(city)
	// @todo this might be also 400 response since email can be a duplicate
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	R.JSON200OK(w)
}

// Update - updates city
func (h *City) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	city, err := validator.CityCreate(body)
	if err != nil || city.ID == 0 {
		log.Println(err)
		R.JSON400(w)
		return
	}

	err = h.Storage.UpdateCity(city)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	R.JSON200OK(w)
}

// Delete - removes city
func (h *City) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		R.JSON400(w)
		return
	}

	err = h.Storage.DeleteCity(intID)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	R.JSON200OK(w)
}

// ID - get city by ID
func (h *City) ID(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		R.JSON400(w)
		return
	}

	city, err := h.Storage.GetCity(intID)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	if city == nil {
		R.JSON404(w)
		return
	}

	R.JSON200(w, city)
}
