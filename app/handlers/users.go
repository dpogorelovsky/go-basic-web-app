package handlers

import (
	"log"
	"net/http"

	"github.com/dpogorelovsky/go-basic-web-app/app/storage/astorage"
	R "github.com/dpogorelovsky/go-basic-web-app/app/utils/responder"
)

// User - users handler/controller
type User struct {
	s astorage.UserStorage
}

// List - returning list of users
func (h *User) List(w http.ResponseWriter, r *http.Request) {

	resp, err := h.s.GetUserList(0, 0)
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
