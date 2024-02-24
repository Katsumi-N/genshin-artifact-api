package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var db *sqlx.DB

func initDB() error {
	var err error
	dbSource := os.Getenv("DB_SOURCE")
	db, err = sqlx.Connect("mysql", dbSource)
	if err != nil {
		fmt.Println("DB connection failed: ", err)
		return err
	}
	fmt.Println("DB connection succeeded")

	return nil
}

func main() {
	err := initDB()
	if err != nil {
		return
	}
	defer db.Close()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! with air")
	})
	e.GET("/characters", fetchCharacters)
	e.Logger.Fatal(e.Start(":8080"))
}
