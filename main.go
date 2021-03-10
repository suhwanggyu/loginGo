package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suhwanggyu/loginGo/key"
	"github.com/suhwanggyu/loginGo/router"
)

func main() {
	key.Generate()
	key.Recovery()
	mux := http.NewServeMux()
	fmt.Println("Login server open")
	mux.HandleFunc("/", router.Handler)
	log.Fatal(http.ListenAndServe(":15555", mux))
}
