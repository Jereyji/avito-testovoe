package entity

import "github.com/google/uuid"

type Merch struct {
	ID    uuid.UUID
	Name  string
	Price int32
}

func NewMerch(name string, price int32) *Merch {
	return &Merch{
		ID:    uuid.New(),
		Name:  name,
		Price: price,
	}
}
