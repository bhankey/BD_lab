package accountservise

import (
	"context"

	"github.com/bhankey/BD_lab/backend/internal/entities/accountentities"
)

func (s *AccountService) GetAll(ctx context.Context) ([]accountentities.Account, error) {
	return s.accountRepo.GetAll(ctx)
}
