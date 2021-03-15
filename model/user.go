package model

import (
	"fmt"
	"log"
)

type User struct {
	Email    string
	Password string
}

func CheckAuth(email string, password string) bool {
	db := connectDB()
	defer db.Close()
	rows, err := db.Query("Select password from user where email = ?", email)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	var _password string
	for rows.Next() {
		err := rows.Scan(&_password)
		if err != nil {
			log.Fatal(err)
		}
	}
	if _password == password {
		return true
	}
	return false
}

func MakeUser(email string, password string) bool {
	defer func(){
		r := recover()
		if r != nil{
			log.Fatal("에러복구")
		}
	}()
	db := connectDB()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer tx.Rollback()
	rows, err := tx.Query("Select count(*) from user where email = ?", email)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Panic(err)
		}
	}
	if count == 1 {
		return false
	}
	_, err = tx.Exec("Insert into user values (?, ?)", email, password)
	if err != nil {
		log.Panic(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Success create a user")
	return true
}