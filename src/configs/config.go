package configs

import (
	"database/sql"
	"fmt"
)

func GetDB() (db *sql.DB, err error) {

	host := "postgres"
	port := "5432"
	user := "admin"
	password := "pass1234"
	dbname := "theater_seating_database"
	connStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	return
}
