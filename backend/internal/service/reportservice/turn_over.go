package reportservice

import (
	"context"
	"finance/internal/entities/paymententities"
	"finance/internal/entities/reportentities"
	"time"
)

func (s *ReportService) GetTurnOverSheets(ctx context.Context, accountIDs []int, year int) ([]reportentities.TurnOverReport, error) {
	report := make([]reportentities.TurnOverReport, 0)

	for _, accountID := range accountIDs {
		payments, err := s.paymentsRepo.GetByAccountID(ctx, accountID, year)
		if err != nil {
			return nil, err
		}

		if len(payments) == 0 {
			s.Logger.Warn("reportservice.GetTurnOverSheets.accountWithoutPaymentsForYear", accountID)

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

		sum := firstPaymentHistory.SumBefore
		for month := time.January; month <= time.December; month++ {
			payments, ok := paymentsByMonth[month]
			if !ok {
				detailsByMonth[month] = reportentities.MonthDetails{
					Income: 0,
					Outgo:  0,
					Sum:    sum,
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

		report = append(report, reportentities.TurnOverReport{
			AccountID:   accountID,
			StartingSum: firstPaymentHistory.SumBefore,
			EndSum:      detailsByMonth[time.December].Sum,
			MothDetails: detailsByMonth,
		})
	}

	return report, nil
}
