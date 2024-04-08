package usecase

import (
	"context"

	"github.com/Katsumi-N/genshin-artifact-api/domain"
	"github.com/Katsumi-N/genshin-artifact-api/infrastructure/repository"
)

type CharacterService struct {
	repo repository.CharacterRepository // リポジトリのインターフェース
}

func NewCharacterService(repo repository.CharacterRepository) *CharacterService {
	return &CharacterService{
		repo: repo,
	}
}

func (s *CharacterService) GetCharacter(ctx context.Context, enkaId string) (*domain.Character, error) {
	return s.repo.GetCharacter(ctx, enkaId)
}

func (s *CharacterService) GetCharacters(ctx context.Context) ([]domain.Character, error) {
	return s.repo.GetCharacters(ctx)
}
