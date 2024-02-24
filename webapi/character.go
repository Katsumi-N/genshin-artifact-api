package main

import (
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

func fetchCharacters(c echo.Context) error {
	ctx := c.Request().Context()
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to begin transaction: "+err.Error())
	}
	defer tx.Rollback()

	var characters []Character
	if err := tx.SelectContext(ctx, &characters, "SELECT * FROM characters"); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "characters not found")
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to select: "+err.Error())
		}
	}
	if err := tx.Commit(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to commit: "+err.Error())
	}

	return c.JSON(http.StatusOK, characters)
}
