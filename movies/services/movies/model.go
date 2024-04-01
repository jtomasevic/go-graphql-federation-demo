package movies

import (
	"github.com/google/uuid"
)



type Movie struct {
	ID    uuid.UUID 
	Title string    
	ImdbID string
}