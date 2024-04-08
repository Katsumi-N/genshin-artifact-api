package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Katsumi-N/genshin-artifact-api/infrastructure/handler"
	"github.com/Katsumi-N/genshin-artifact-api/infrastructure/repository"
	"github.com/Katsumi-N/genshin-artifact-api/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	dbSource := os.Getenv("DB_SOURCE")
	db, err := sqlx.Connect("mysql", dbSource)
	if err != nil {
		fmt.Println("DB connection failed: ", err)
		return
	}
	defer db.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! with air")
	})

	repository := repository.NewCharacterRepository(db)
	service := usecase.NewCharacterService(repository)
	characterHandler := handler.NewCharacterHandler(service)

	e.GET("/characters", characterHandler.FetchCharacters)
	e.GET("/characters/:id", characterHandler.FetchCharacter)
	// e.GET("/genshin/:genshin_uuid/characters", fetchOwnedCharactersListHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
