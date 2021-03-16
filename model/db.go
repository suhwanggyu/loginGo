package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func ViperEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal("Type error")
	}
	return value
}

// @dev please close db
func connectDB() *sql.DB {
	host := ViperEnv("DBHOST")
	user := ViperEnv("USER")
	pwd := ViperEnv("PASSWORD")
	config := user + ":" + pwd + "@tcp(" + host + ")/user"
	db, err := sql.Open("mysql", config)
	if err != nil {
		panic("Can't access database")
		return nil
	}
	return db
}
