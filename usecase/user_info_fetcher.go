package usecase

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Katsumi-N/genshin-artifact-api/domain"
	"github.com/Katsumi-N/genshin-artifact-api/infrastructure/repository"
)

type UserInfoFetcher struct {
	httpClient repository.HttpClientRepository // domain層のインターフェース
}

func NewUserInfoFetcher(httpClient repository.HttpClientRepository) *UserInfoFetcher {
	return &UserInfoFetcher{
		httpClient: httpClient,
	}
}

func (u *UserInfoFetcher) FetchUserInfo(genshin_uuid int) (domain.EnkaUserInfo, error) {
	responseBody, statusCode, err := u.httpClient.SendRequest("GET", "https://enka.network", "/api/uid/"+strconv.Itoa(genshin_uuid), "", nil)
	if err != nil || statusCode != http.StatusOK {
		return domain.EnkaUserInfo{}, err // Error handling may need to be more sophisticated
	}

	var userInfo domain.EnkaUserInfo
	if err := json.Unmarshal(responseBody, &userInfo); err != nil {
		return domain.EnkaUserInfo{}, err
	}

	return userInfo, nil
}
