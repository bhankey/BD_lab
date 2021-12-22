package reportservice

import (
	"context"
	"finance/internal/entities/accountentities"
	"finance/internal/entities/paymententities"
	"finance/internal/service"
)

type ReportService struct {
	*service.BaseService

	accountRepo         accountRepository
	paymentsRepo        paymentsRepository
	paymentsHistoryRepo paymentsHistoryRepository
}

type paymentsRepository interface {
	GetByAccountID(ctx context.Context, accountID int, year int) ([]paymententities.Payment, error)
	GetStatistic(ctx context.Context, accountID int) (paymententities.PaymentsStatistic, error)
}

type accountRepository interface {
	GetOne(ctx context.Context, accountID int) (accountentities.Account, error)
	GetAll(ctx context.Context) ([]accountentities.Account, error)
	GetDebtors(ctx context.Context, accountIDs []int) ([]accountentities.Account, error)
	GetAllDebtors(ctx context.Context) ([]accountentities.Account, error)
}

type paymentsHistoryRepository interface {
	GetPayment(ctx context.Context, accountID int, paymentID int) (paymententities.PaymentHistory, error)
}

func NewReportService(baseService *service.BaseService, paymentRepo paymentsRepository, paymentsHistoryRepo paymentsHistoryRepository, accountRepo accountRepository) *ReportService {
	return &ReportService{
		BaseService:         baseService,
		paymentsRepo:        paymentRepo,
		paymentsHistoryRepo: paymentsHistoryRepo,
		accountRepo:         accountRepo,
	}
}
