package accountservise

import "context"

func (s *AccountService) Update(ctx context.Context, accountID int, newName string) error {
	return s.accountRepo.Update(ctx, accountID, newName)
}

func (s *AccountService) Delete(ctx context.Context, accountID int) error {
	return s.accountRepo.Delete(ctx, accountID)
}
