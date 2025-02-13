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

func (r *Repository) FetchMerchByID(ctx context.Context, merchID uuid.UUID) (*entity.Merch, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var merch entity.Merch
	err := r.db.QueryRow(ctx, queries.QueryFetchMerchByID, merchID).Scan(
		&merch.ID,
		&merch.Name,
		&merch.Price,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repos.ErrNotFound
		}

		return nil, err
	}

	return &merch, nil
}

func (r *Repository) CreateMerch(ctx context.Context, merch *entity.Merch) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryCreateMerch,
		merch.ID,
		merch.Name,
		merch.Price,
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

func (r *Repository) UpdateMerch(ctx context.Context, merch *entity.Merch) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryUpdateMerch,
		merch.ID,
		merch.Name,
		merch.Price,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteMerch(ctx context.Context, merchID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryDeleteMerch, merchID)
	if err != nil {
		return err
	}

	return nil
}
