package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Character struct {
	ID        int64     `json:"id" db:"id"`
	EnkaId    string    `json:"enka_id" db:"enka_id"`
	Name      string    `json:"name" db:"name"`
	ImageUrl  string    `json:"image_url" db:"image_url"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}

func fetchCharacter(c echo.Context) error {
	enkaId := c.Param("id")
	ctx := c.Request().Context()
	character, err := getCharacter(enkaId, ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "character not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch character: "+err.Error())
	}
	return c.JSON(http.StatusOK, character)
}

func fetchCharacters(c echo.Context) error {
	ctx := c.Request().Context()
	characters, err := getCharacters(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch characters: "+err.Error())
	}
	return c.JSON(http.StatusOK, characters)
}

func getCharacter(enkaId string, ctx context.Context) (*Character, error) {
	var character Character
	err := db.GetContext(ctx, &character, "SELECT * FROM characters WHERE enka_id = ?", enkaId)
	if err != nil {
		return nil, err
	}
	return &character, nil
}

func getCharacters(ctx context.Context) ([]Character, error) {
	var characters []Character
	err := db.SelectContext(ctx, &characters, "SELECT * FROM characters")
	if err != nil {
		return nil, err
	}
	return characters, nil
}
