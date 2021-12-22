package accountservise

import (
	"context"
	"finance/internal/entities/accountentities"
	"finance/internal/service"
)

type AccountService struct {
	*service.BaseService

	accountRepo AccountRepository
}

type AccountRepository interface {
	Create(ctx context.Context, name string, userID int) error
	GetAll(ctx context.Context) ([]accountentities.Account, error)
	Update(ctx context.Context, accountID int, name string) error
	Delete(ctx context.Context, accountID int) error
}

func NewAccountService(baseService *service.BaseService, accountRepo AccountRepository) *AccountService {
	return &AccountService{
		BaseService: baseService,
		accountRepo: accountRepo,
	}
}
