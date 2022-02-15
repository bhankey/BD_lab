package reportservice

import (
	"context"
	"time"

	"github.com/bhankey/BD_lab/backend/internal/entities/paymententities"
	"github.com/bhankey/BD_lab/backend/internal/entities/reportentities"
)

func (s *ReportService) GetTurnOverSheets(
	ctx context.Context,
	accountIDs []int,
	year int,
) ([]reportentities.TurnOverReport, error) {
	if len(accountIDs) == 0 {
		accounts, err := s.accountRepo.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		accountIDs = make([]int, 0, len(accounts))
		for _, account := range accounts {
			accountIDs = append(accountIDs, account.ID)
		}
	}

	report, err := s.getTurnOverReportByAccountID(ctx, accountIDs, year)
	if err != nil {
		return nil, err
	}

	return report, nil
}

func (s *ReportService) getTurnOverReportByAccountID(
	ctx context.Context,
	accountIDs []int,
	year int,
) ([]reportentities.TurnOverReport, error) {
	report := make([]reportentities.TurnOverReport, 0)

	for _, accountID := range accountIDs {
		payments, err := s.paymentsRepo.GetByAccountID(ctx, accountID, year)
		if err != nil {
			return nil, err
		}

		if len(payments) == 0 {
			s.Logger.WithField("account_id", accountID).
				Warn("reportservice.GetTurnOverSheets.accountWithoutPaymentsForYear", accountID)

			continue
		}

		firstYearPayment := payments[0]

		paymentsByMonth := map[time.Month][]paymententities.Payment{}
		for _, payment := range payments {
			payments := paymentsByMonth[payment.Date.Month()]

			payments = append(payments, payment)
			paymentsByMonth[payment.Date.Month()] = payments

			if payment.Date.Before(firstYearPayment.Date) {
				firstYearPayment = payment
			}
		}

		detailsByMonth := map[time.Month]reportentities.MonthDetails{}

		// Лучше сразу до цикла запросить все данные у БД и распарсить их в мапу, но лень
		firstPaymentHistory, err := s.paymentsHistoryRepo.GetPayment(ctx, firstYearPayment.AccountID, firstYearPayment.ID)
		if err != nil {
			return nil, err
		}

		s.getTurnOverMonthDetails(paymentsByMonth, detailsByMonth, firstPaymentHistory.SumBefore)

		report = append(report, reportentities.TurnOverReport{
			AccountID:   accountID,
			StartingSum: firstPaymentHistory.SumBefore,
			EndSum:      detailsByMonth[time.December].Sum,
			MothDetails: detailsByMonth,
		})
	}

	return report, nil
}

func (s *ReportService) getTurnOverMonthDetails(
	paymentsByMonth map[time.Month][]paymententities.Payment,
	detailsByMonth map[time.Month]reportentities.MonthDetails,
	sum float64) {
	for month := time.January; month <= time.December; month++ {
		payments, ok := paymentsByMonth[month]
		if !ok {
			detailsByMonth[month] = reportentities.MonthDetails{
				Sum: sum,
			}
		}

		income := 0.0
		outgo := 0.0
		for _, payment := range payments {
			sum += payment.Sum

			if payment.Sum >= 0 {
				income += payment.Sum
			} else {
				outgo += payment.Sum
			}
		}

		detailsByMonth[month] = reportentities.MonthDetails{
			Income: income,
			Outgo:  outgo,
			Sum:    sum,
		}
	}
}
