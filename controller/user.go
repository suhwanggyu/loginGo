package controller

import (
	"encoding/json"
	"net/http"

	"github.com/suhwanggyu/loginGo/model"
)

func ControlUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	json.NewDecoder(r.Body).Decode(&user)
	fin := model.MakeUser(user.Email, user.Password)
	if fin == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
