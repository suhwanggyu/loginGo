package model

import (
	"fmt"
	"log"
)

type User struct {
	email    string
	password string
}

func CheckAuth(email string, password string) bool {
	db := connectDB()
	defer db.Close()
	rows, err := db.Query("Select email from user where password = ?", password)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	var _email string
	for rows.Next() {
		err := rows.Scan(&_email)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _email == email {
		fmt.Println(_email, email)
		return true
	}
	return false
}
