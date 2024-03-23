package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func fetchCharacterHandler(c echo.Context) error {
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

func fetchCharactersHandler(c echo.Context) error {
	ctx := c.Request().Context()
	characters, err := getCharacters(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch characters: "+err.Error())
	}
	return c.JSON(http.StatusOK, characters)
}

func fetchOwnedCharactersListHandler(c echo.Context) error {
	genshin_uuid, err := strconv.Atoi(c.Param("genshin_uuid"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid genshin_uuid")
	}
	enkaUserInfo, err := fetchUserInfo(genshin_uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch user info from Enka: "+err.Error())
	}

	characterStateses := enkaUserInfo.ExtractCharacterStatus()
	return c.JSON(http.StatusOK, characterStateses)
}
