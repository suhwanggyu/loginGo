package controller

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/suhwanggyu/loginGo/key"
)

type Token interface {
	sig(string)
}

type TokenExpired struct {
	Expired time.Time
	Email   string
	Sigdata []byte
}

type Request struct {
	Method   string `json:"method"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (token *TokenExpired) sig(email string) {
	privatekey := key.Recovery()
	hash := sha256.New()
	token.Expired = time.Now().Add(1 * time.Hour)
	token.Email = email
	bin := make([]byte, 8)
	binary.LittleEndian.PutUint64(bin, uint64(token.Expired.Unix()))
	x := append([]byte(email), bin...)
	fmt.Println(x)
	hash.Write(x)
	dig := hash.Sum(nil)
	data, _ := privatekey.Sign(rand.Reader, dig, crypto.SHA256)
	token.Sigdata = data

	err := rsa.VerifyPKCS1v15(&privatekey.PublicKey, crypto.SHA256, dig, data)
	if err == nil {
		fmt.Println("Success publish a token")
	}
}

func ControlToken(w http.ResponseWriter, r *http.Request) {
	req := Request{}
	json.NewDecoder(r.Body).Decode(&req)
	var token Token
	switch req.Method {
	case "Expired":
		token = &TokenExpired{}
		// @dev Will make more option
	}
	token.sig(req.Email)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}
