package paymentsservice

import (
	"context"
)

func (s *PaymentsService) Create(ctx context.Context, accountID int, sum float64, reason string) error {
	account, err := s.accountRepo.GetOne(ctx, accountID)
	if err != nil {
		return err
	}

	paymentID, err := s.paymentsRepo.Create(ctx, accountID, sum, reason)
	if err != nil {
		return err
	}

	if err := s.paymentsHistoryRepo.Create(ctx, accountID, paymentID, account.Sum); err != nil {
		return err
	}

	newSum := account.Sum + sum
	if err := s.accountRepo.ChangeSum(ctx, accountID, newSum); err != nil {
		return err
	}

	return nil
}
