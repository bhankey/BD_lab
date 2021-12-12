package accountservise

import "context"

func (s *AccountService) Create(ctx context.Context, name string, userID int) error {
	return s.accountRepo.Create(ctx, name, userID)
}
