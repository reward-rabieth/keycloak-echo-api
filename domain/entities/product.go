package entities

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	Price     float32
}
