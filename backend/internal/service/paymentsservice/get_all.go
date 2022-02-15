package paymentsservice

import (
	"context"

	"github.com/bhankey/BD_lab/backend/internal/entities/paymententities"
)

func (s *PaymentsService) GetAll(ctx context.Context) ([]paymententities.Payment, error) {
	accounts, err := s.accountRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	accountIDs := make([]int, 0, len(accounts))
	for _, account := range accounts {
		accountIDs = append(accountIDs, account.ID)
	}

	return s.paymentsRepo.GetClientsPayments(ctx, accountIDs)
}
