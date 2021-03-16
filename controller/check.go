package controller

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"github.com/suhwanggyu/loginGo/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CheckTokenExpired(pubkey rsa.PublicKey, token TokenExpired) bool{
	if token.Expired.Unix() < time.Now().Unix() {
		return false
	}
	bin := make([]byte, 8)
	binary.LittleEndian.PutUint64(bin, uint64(token.Expired.Unix()))
	x := append([]byte(token.Email), bin...)
	hash := sha256.New()
	hash.Write(x)
	dig := hash.Sum(nil)
	err := rsa.VerifyPKCS1v15(&pubkey, crypto.SHA256, dig, token.Sigdata)
	if err == nil {
		fmt.Println("Verify")
		return true
	}
	return false
}

func RequestPubKey() *rsa.PublicKey {
	host := model.ViperEnv("LOGINHOST")
	res, err := http.Get(host + "/publickey")
	if err != nil {
		panic("Fail to take public key")
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	pubkey, err := x509.ParsePKCS1PublicKey(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Take pubkey")
	return pubkey
}

