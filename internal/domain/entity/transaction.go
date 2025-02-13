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

func NewTransaction(senderID, recepientID uuid.UUID, amount int32) *Transaction {
	return &Transaction{
		ID:          uuid.New(),
		SenderID:    senderID,
		RecepientID: recepientID,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
}
