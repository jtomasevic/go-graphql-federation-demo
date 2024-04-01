package actors

import (
	"github.com/google/uuid"
)

type Actor struct {
	ID   uuid.UUID 
	Name string
	MovieImdbIds []string
	ImdbID string
}

func (Actor) IsEntity() {}