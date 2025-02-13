package repos

import "errors"

var (
	CodeDuplicateRow = "23505"

	ErrRowExist = errors.New("object is exist")
	ErrNotFound = errors.New("object not found")
)