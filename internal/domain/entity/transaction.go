package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID
	SenderID    uuid.UUID
	RecepientID uuid.UUID
	Amount      int32
	CreatedAt   time.Time
}
