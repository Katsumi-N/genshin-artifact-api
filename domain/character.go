package domain

import (
	"time"
)

type Character struct {
	ID        int64     `json:"id" db:"id"`
	EnkaId    string    `json:"enka_id" db:"enka_id"`
	Name      string    `json:"name" db:"name"`
	ImageUrl  string    `json:"image_url" db:"image_url"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}
