package controller

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

func CheckTokenExpired(pubkey rsa.PublicKey, token TokenExpired) bool{
	if token.Expired.Unix() > time.Now().Unix() {
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