package repository

import (
	"context"

	"github.com/Katsumi-N/genshin-artifact-api/domain"
)

type CharacterRepository interface {
	GetCharacter(ctx context.Context, enkaId string) (*domain.Character, error)
	GetCharacters(ctx context.Context) ([]domain.Character, error)
}
