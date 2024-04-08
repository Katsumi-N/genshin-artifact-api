package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
