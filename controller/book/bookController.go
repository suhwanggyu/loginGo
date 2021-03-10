package book

import "net/http"

func HandleBook(w http.ResponseWriter, r *http.Request) {
	//프론트가 고정 IP 이용하게 되면 변경 예정
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	switch r.Method {
	case "OPTIONS":
	case "POST":
	case "PUT":
	case "DELETE":
	}
}
