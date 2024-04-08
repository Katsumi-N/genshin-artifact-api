package repository

import (
	"context"

	"github.com/Katsumi-N/genshin-artifact-api/domain"
	"github.com/Katsumi-N/genshin-artifact-api/util"
	"github.com/jmoiron/sqlx"
)

type characterRepository struct {
	db *sqlx.DB
}

func NewCharacterRepository(db *sqlx.DB) CharacterRepository {
	return &characterRepository{db: db}
}

func (r *characterRepository) GetCharacter(ctx context.Context, enkaId string) (*domain.Character, error) {
	var character domain.Character
	err := util.Db.GetContext(ctx, &character, "SELECT * FROM characters WHERE enka_id = ?", enkaId)
	if err != nil {
		return nil, err
	}
	return &character, nil
}

func (r *characterRepository) GetCharacters(ctx context.Context) ([]domain.Character, error) {
	var characters []domain.Character
	err := util.Db.SelectContext(ctx, &characters, "SELECT * FROM characters")
	if err != nil {
		return nil, err
	}
	return characters, nil
}
