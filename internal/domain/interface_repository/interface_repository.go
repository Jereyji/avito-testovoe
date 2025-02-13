package repos

import (
	"context"

	"github.com/Jereyji/avito-testovoe/internal/domain/entity"
	"github.com/google/uuid"
)

type IRepository interface {
	FetchUser(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	FetchUserByUsername(ctx context.Context, username string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error

	FetchMerchByID(ctx context.Context, merchID uuid.UUID) (*entity.Merch, error)
	CreateMerch(ctx context.Context, merch *entity.Merch) error
	UpdateMerch(ctx context.Context, merch *entity.Merch) error
	DeleteMerch(ctx context.Context, merchID uuid.UUID) error

	FetchPurchasesByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Purchase, error)
	CreatePurchase(ctx context.Context, purchase *entity.Purchase) error
	UpdatePurchase(ctx context.Context, purchase *entity.Purchase) error
	DeletePurchase(ctx context.Context, purchaseID uuid.UUID) error

	FetchTransactionsByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Transaction, error)
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) error
	UpdateTransaction(ctx context.Context, transaction *entity.Transaction) error
	DeleteTransaction(ctx context.Context, transactionID uuid.UUID) error
}
