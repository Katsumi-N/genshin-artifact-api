package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func fetchUserInfo(genshin_uuid int) (EnkaUserInfo, error) {
	endpoint := "/api/uid/" + strconv.Itoa(genshin_uuid)
	resBody, _, err := sendRequest("GET", "https://enka.network", endpoint, "", nil)
	if err != nil {
		return EnkaUserInfo{}, echo.NewHTTPError(http.StatusInternalServerError, "failed to send request: "+err.Error())
	}

	var enkaUserInfo EnkaUserInfo
	if err := json.Unmarshal(resBody, &enkaUserInfo); err != nil {
		return EnkaUserInfo{}, echo.NewHTTPError(http.StatusInternalServerError, "failed to unmarshal response: "+err.Error())
	}

	return enkaUserInfo, nil
}
