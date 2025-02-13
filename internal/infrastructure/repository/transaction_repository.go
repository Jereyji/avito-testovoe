package repository

import (
	"context"

	"github.com/Jereyji/avito-testovoe/internal/domain/entity"
	"github.com/Jereyji/avito-testovoe/internal/infrastructure/repository/queries"
	"github.com/google/uuid"
)

func (r *Repository) FetchTransactionsByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Transaction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rows, err := r.db.Query(ctx, queries.QueryFetchTransactionsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entity.Transaction
	for rows.Next() {
		var transaction entity.Transaction
		if err := rows.Scan(
			&transaction.ID,
			&transaction.SenderID,
			&transaction.RecepientID,
			&transaction.Amount,
			&transaction.CreatedAt,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *Repository) CreateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryCreateTransaction,
		transaction.ID,
		transaction.SenderID,
		transaction.RecepientID,
		transaction.Amount,
		transaction.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryUpdateTransaction,
		transaction.ID,
		transaction.Amount,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteTransaction(ctx context.Context, transactionID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.QueryDeleteTransaction, transactionID)
	if err != nil {
		return err
	}

	return nil
}
