package util

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDB() error {
	var err error
	dbSource := os.Getenv("DB_SOURCE")
	Db, err = sqlx.Connect("mysql", dbSource)
	if err != nil {
		fmt.Println("DB connection failed: ", err)
		return err
	}
	fmt.Println("DB connection succeeded")

	return nil
}
