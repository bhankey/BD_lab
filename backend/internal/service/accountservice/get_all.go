package accountservise

import (
	"context"
	"finance/internal/entities/accountentities"
)

func (s *AccountService) GetAll(ctx context.Context) ([]accountentities.Account, error) {
	return s.accountRepo.GetAll(ctx)
}
