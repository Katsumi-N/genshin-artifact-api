package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/Katsumi-N/genshin-artifact-api/domain"
	"github.com/labstack/echo/v4"
)

// CharacterService は、アプリケーション層のサービスインターフェースです。
// これは、ハンドラーがアプリケーション層の機能にアクセスするために使用します。
type CharacterService interface {
	GetCharacter(ctx context.Context, enkaId string) (*domain.Character, error)
	GetCharacters(ctx context.Context) ([]domain.Character, error)
}

type CharacterHandler struct {
	service CharacterService
}

// NewCharacterHandler は新しいCharacterHandlerを作成し、返します。
func NewCharacterHandler(service CharacterService) *CharacterHandler {
	return &CharacterHandler{service: service}
}

func (h *CharacterHandler) FetchCharacter(c echo.Context) error {
	enkaId := c.Param("id")
	ctx := c.Request().Context()

	character, err := h.service.GetCharacter(ctx, enkaId) // usecase層に定義された関数を呼び出す
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "character not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch character: "+err.Error())
	}
	return c.JSON(http.StatusOK, character)
}

func (h *CharacterHandler) FetchCharacters(c echo.Context) error {
	ctx := c.Request().Context()
	characters, err := h.service.GetCharacters(ctx) // usecase層に定義された関数を呼び出す
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch characters: "+err.Error())
	}
	return c.JSON(http.StatusOK, characters)
}
