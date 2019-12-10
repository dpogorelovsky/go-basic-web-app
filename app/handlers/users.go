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

// User - users handler/controller
type User struct {
	Storage astorage.UserStorage
}

// List - returning list of users
func (h *User) List(w http.ResponseWriter, r *http.Request) {
	limit, offset := utils.GetPaginationParams(r.URL.Query())
	resp, err := h.Storage.GetUserList(limit, offset)
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

// Create - creates new user
func (h *User) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	user, err := validator.UserCreate(body)
	if err != nil {
		log.Println(err)
		R.JSON400(w)
		return
	}

	err = h.Storage.CreateUser(user)
	// @todo this might be also 400 response since email can be a duplicate
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	R.JSON200OK(w)
}

// Update - updates pointed user
func (h *User) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	// @todo we might want extra check that /users/id equals to user.ID received in body
	user, err := validator.UserCreate(body)
	if err != nil || user.ID == 0 {
		log.Println(err)
		R.JSON400(w)
		return
	}

	err = h.Storage.UpdateUser(user)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	R.JSON200OK(w)
}

// Delete - removes user
func (h *User) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		R.JSON400(w)
		return
	}

	err = h.Storage.DeleteUser(intID)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	R.JSON200OK(w)
}

// ID - get user by ID
func (h *User) ID(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		R.JSON400(w)
		return
	}

	user, err := h.Storage.GetUser(intID)
	if err != nil {
		log.Println(err)
		R.JSON500(w)
		return
	}

	if user == nil {
		R.JSON404(w)
		return
	}

	R.JSON200(w, user)
}
