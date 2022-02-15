package paymentsservice

import (
	"context"

	"github.com/bhankey/BD_lab/backend/internal/entities/paymententities"
)

func (s *PaymentsService) GetClientPayments(ctx context.Context, accountID int) ([]paymententities.Payment, error) {
	return s.paymentsRepo.GetClientPayments(ctx, accountID)
}
