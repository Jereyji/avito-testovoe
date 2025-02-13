package repository

import (
	"context"

	"github.com/Jereyji/avito-testovoe/internal/domain/entity"
	"github.com/Jereyji/avito-testovoe/internal/infrastructure/repository/queries"
	"github.com/google/uuid"
)

func (r *Repository) FetchPurchasesByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Purchase, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rows, err := r.db.Query(ctx, queries.QueryFetchPurchasesByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []entity.Purchase
	for rows.Next() {
		var purchase entity.Purchase
		if err := rows.Scan(
			&purchase.ID,
			&purchase.UserID,
			&purchase.MerchID,
			&purchase.CreatedAt,
		); err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return purchases, nil
}

func (r *Repository) CreatePurchase(ctx context.Context, purchase *entity.Purchase) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryCreatePurchase,
		purchase.ID,
		purchase.UserID,
		purchase.MerchID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePurchase(ctx context.Context, purchase *entity.Purchase) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryUpdatePurchase,
		purchase.ID,
		purchase.UserID,
		purchase.MerchID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeletePurchase(ctx context.Context, purchaseID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryDeletePurchase, purchaseID)
	if err != nil {
		return err
	}

	return nil
}
