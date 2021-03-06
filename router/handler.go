package router

import (
	"fmt"
	"net/http"

	"github.com/suhwanggyu/loginGo/controller"
)

func optionAndCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func routing(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/token":
		controller.ControlToken(w, r)
	case "/user":
		controller.ControlUser(w, r)
	case "/publickey":
		controller.ControlPubkey(w, r)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request arrived")
	switch r.Method {
	case "OPTIONS":
		optionAndCors(w, r)
		fmt.Fprint(w, nil)
	default:
		optionAndCors(w, r)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		routing(w, r)
	}
}
