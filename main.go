package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/suhwanggyu/loginGo/key"
	"github.com/suhwanggyu/loginGo/router"
)

func main() {
	gen := flag.Bool("gen", false,"a boolean")
	if *gen {
		key.Generate()
	}
	mux := http.NewServeMux()
	fmt.Println("Login server open")
	mux.HandleFunc("/", router.Handler)
	log.Fatal(http.ListenAndServe(":15555", mux))
}
