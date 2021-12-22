package paymentrepo

import (
	"context"
	"database/sql"
	"errors"
	"finance/internal/entities/paymententities"
	"finance/internal/repository"
)

func (r *PaymentsRepo) GetStatistic(ctx context.Context, accountID int) (paymententities.PaymentsStatistic, error) {
	const query = `
		SELECT
                account_id,
                SUM(CASE WHEN sum > 0 THEN sum ELSE 0 END) as income,
                SUM(CASE WHEN sum < 0 THEN sum ELSE 0 END) as outgo
        FROM
                payments
        WHERE account_id = $1
        GROUP BY account_id
`

	row := struct {
		AccountID int     `db:"account_id"`
		Income    float64 `db:"income"`
		Outgo     float64 `db:"outgo"`
	}{}

	if err := r.db.GetContext(ctx, &row, query, accountID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return paymententities.PaymentsStatistic{}, repository.NoEntity
		}

		r.Logger.Error("paymentrepo.GetPaymentsStatistic.QueryError")

		return paymententities.PaymentsStatistic{}, err
	}

	result := paymententities.PaymentsStatistic{
		AccountID:     row.AccountID,
		SummaryIncome: row.Income,
		SummaryOutgo:  row.Outgo,
	}

	return result, nil

}
