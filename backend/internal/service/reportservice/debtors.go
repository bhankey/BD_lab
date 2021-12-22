package reportservice

import (
	"context"
	"finance/internal/entities/accountentities"
	"finance/internal/entities/reportentities"
)

func (s *ReportService) GetDebtors(ctx context.Context, accountIDs []int) ([]reportentities.DebtorsDetails, error) {
	var err error
	debtors := make([]accountentities.Account, 0, len(accountIDs))
	if len(accountIDs) == 0 {
		debtors, err = s.accountRepo.GetAllDebtors(ctx)
	} else {
		debtors, err = s.accountRepo.GetDebtors(ctx, accountIDs)
	}

	if err != nil {
		return nil, err
	}

	result := make([]reportentities.DebtorsDetails, 0, len(debtors))
	for _, debtor := range debtors {
		// Лучше сразу до цикла запросить все данные у БД и распарсить их в мапу, но лень
		statistic, err := s.paymentsRepo.GetStatistic(ctx, debtor.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, reportentities.DebtorsDetails{
			AccountID: debtor.ID,
			Income:    statistic.SummaryIncome,
			Outgo:     statistic.SummaryOutgo,
			Own:       debtor.Sum,
		})
	}

	return result, nil
}
