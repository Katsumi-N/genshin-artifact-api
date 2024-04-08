package main

import (
	"net/http"

	"github.com/Katsumi-N/genshin-artifact-api/infrastructure/handler"
	"github.com/Katsumi-N/genshin-artifact-api/infrastructure/repository"
	"github.com/Katsumi-N/genshin-artifact-api/usecase"
	"github.com/Katsumi-N/genshin-artifact-api/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	err := util.InitDB()
	if err != nil {
		return
	}
	defer util.Db.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! with air")
	})

	repository := repository.NewCharacterRepository(util.Db)
	service := usecase.NewCharacterService(repository)
	characterHandler := handler.NewCharacterHandler(service)

	e.GET("/characters", characterHandler.FetchCharacters)
	e.GET("/characters/:id", characterHandler.FetchCharacter)
	// e.GET("/genshin/:genshin_uuid/characters", fetchOwnedCharactersListHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
