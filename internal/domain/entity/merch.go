package entity

import "github.com/google/uuid"

type Merch struct {
	ID    uuid.UUID
	Name  string
	Price int32
	Size  string
}
