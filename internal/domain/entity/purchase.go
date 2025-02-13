package entity

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	MerchID   uuid.UUID
	CreatedAt time.Time
}
