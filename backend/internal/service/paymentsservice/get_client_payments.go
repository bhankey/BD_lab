package paymentsservice

import (
	"context"
	"finance/internal/entities/paymententities"
)

func (s *PaymentsService) GetClientPayments(ctx context.Context, accountID int) ([]paymententities.Payment, error) {
	return s.paymentsRepo.GetClientPayments(ctx, accountID)
}
