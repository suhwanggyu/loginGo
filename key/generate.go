package key

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Generate() {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic("Key failed")
	}
	file, err := os.Create("key.gen")
	defer file.Close()
	if err != nil {
		panic("file create failed")
	}
	data, err := json.Marshal(privatekey)
	fmt.Fprint(file, string(data))
}

func Recovery() *rsa.PrivateKey {
	jsonfile, err := os.Open("key.gen")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonfile.Close()
	bytes, _ := ioutil.ReadAll(jsonfile)
	var private rsa.PrivateKey
	json.Unmarshal(bytes, &private)
	return &private
}
