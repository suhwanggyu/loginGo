package controller

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/suhwanggyu/loginGo/key"
	"net/http"
	"time"

	"github.com/suhwanggyu/loginGo/model"
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

var privateKey *rsa.PrivateKey

func (token *TokenExpired) sig(email string) {
	hash := sha256.New()
	token.Expired = time.Now().Add(1 * time.Hour)
	token.Email = email
	bin := make([]byte, 8)
	binary.LittleEndian.PutUint64(bin, uint64(token.Expired.Unix()))
	x := append([]byte(email), bin...)
	hash.Write(x)
	dig := hash.Sum(nil)
	data, _ := privateKey.Sign(rand.Reader, dig, crypto.SHA256)
	token.Sigdata = data

	err := rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, dig, data)
	if err == nil {
		fmt.Println("Success publish a token")
	}
}


func ControlToken(w http.ResponseWriter, r *http.Request) {
	if privateKey == nil {
		privateKey = key.Recovery()
	}
	req := Request{}
	json.NewDecoder(r.Body).Decode(&req)
	check := model.CheckAuth(req.Email, req.Password)
	if check == false {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w)
		return
	}
	var token Token
	switch req.Method {
	case "Expired":
		token = &TokenExpired{}
		// @dev Will make more option
	default:
		return
	}
	token.sig(req.Email)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}

func ControlPubkey(w http.ResponseWriter, r *http.Request) {
	if privateKey == nil {
		privateKey = key.Recovery()
	}
	pub := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	w.WriteHeader(http.StatusAccepted)
	w.Write(pub)
}