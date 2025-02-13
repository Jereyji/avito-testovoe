package entity

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uuid.UUID
	Username       string
	HashedPassword string
	Balance        int32
}

func NewUser(username, password string, balance int32) (*User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:             uuid.New(),
		Username:       username,
		HashedPassword: hashedPassword,
		Balance:        balance,
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func (u User) VerifyPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrInvalidUsernameOrPassword
		}

		return err
	}

	return nil
}
