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

func NewPurchase(userID, merchID uuid.UUID) *Purchase {
	return &Purchase{
		ID:        uuid.New(),
		UserID:    userID,
		MerchID:   merchID,
		CreatedAt: time.Now(),
	}
}
