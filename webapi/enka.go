package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func fetchUserInfo(c echo.Context) error {
	resBody, statusCode, err := sendRequest("GET", "https://enka.network", "/api/uid/814178891", "", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch user info. returned: "+strconv.Itoa(statusCode)+err.Error())
	}

	var enkaUserInfo EnkaUserInfo
	if err := json.Unmarshal(resBody, &enkaUserInfo); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to unmarshal response body: "+err.Error())
	}

	return c.JSON(http.StatusOK, enkaUserInfo)
}
