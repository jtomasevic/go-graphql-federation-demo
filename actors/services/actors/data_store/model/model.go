package model

import "github.com/google/uuid"

type Actor struct {
	ID         uuid.UUID
	Name       string
	MovieImdbs string
	ImdbID     string
}
