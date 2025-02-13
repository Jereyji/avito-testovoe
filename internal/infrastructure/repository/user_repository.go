package repository

import (
	"context"
	"errors"

	"github.com/Jereyji/avito-testovoe/internal/domain/entity"
	repos "github.com/Jereyji/avito-testovoe/internal/domain/interface_repository"
	"github.com/Jereyji/avito-testovoe/internal/infrastructure/repository/queries"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) FetchUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var user entity.User
	err := r.db.QueryRow(ctx, queries.QueryFetchUserByUsername, username).Scan(
		&user.ID,
		&user.Username,
		&user.HashedPassword,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repos.ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *Repository) FetchUser(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var user entity.User
	err := r.db.QueryRow(ctx, queries.QueryFetchUserByID, userID).Scan(
		&user.ID,
		&user.Username,
		&user.HashedPassword,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repos.ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryCreateUser,
		user.ID,
		user.Username,
		user.HashedPassword,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == repos.CodeDuplicateRow {
			return repos.ErrRowExist
		}

		return err
	}

	return nil
}

func (r *Repository) UpdateUser(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryUpdateUser,
		user.ID,
		user.Username,
		user.HashedPassword,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryDeleteUser, userID)
	if err != nil {
		return err
	}

	return nil
}
