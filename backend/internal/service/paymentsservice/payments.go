package paymentsservice

import (
	"context"
	"finance/internal/entities/accountentities"
	"finance/internal/entities/paymententities"
	"finance/internal/service"
)

type PaymentsService struct {
	*service.BaseService

	accountRepo         accountRepository
	paymentsRepo        paymentsRepository
	paymentsHistoryRepo paymentsHistoryRepository
}

type paymentsRepository interface {
	Create(ctx context.Context, accountID int, sum float64, reason string) (int, error)
	GetAll(ctx context.Context) ([]paymententities.Payment, error)
	GetClientPayments(ctx context.Context, accountID int) ([]paymententities.Payment, error)
	GetClientsPayments(ctx context.Context, accountID []int) ([]paymententities.Payment, error)
}

type accountRepository interface {
	GetAll(ctx context.Context) ([]accountentities.Account, error)
	GetOne(ctx context.Context, accountID int) (accountentities.Account, error)
	ChangeSum(ctx context.Context, accountID int, sum float64) error
}

type paymentsHistoryRepository interface {
	Create(ctx context.Context, accountID int, paymentID int, sumBefore float64) error
}

func NewPaymentsService(baseService *service.BaseService, paymentRepo paymentsRepository, paymentsHistoryRepo paymentsHistoryRepository, accountRepo accountRepository) *PaymentsService {
	return &PaymentsService{
		BaseService:         baseService,
		paymentsRepo:        paymentRepo,
		paymentsHistoryRepo: paymentsHistoryRepo,
		accountRepo:         accountRepo,
	}
}
