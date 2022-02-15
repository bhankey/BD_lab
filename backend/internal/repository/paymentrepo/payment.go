package paymentrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/bhankey/BD_lab/backend/internal/entities/paymententities"
	"github.com/bhankey/BD_lab/backend/internal/repository"
	"github.com/jmoiron/sqlx"
)

type PaymentsRepo struct {
	*repository.BaseRepository

	db *sqlx.DB
}

func NewPaymentsRepo(baseRepo *repository.BaseRepository, db *sqlx.DB) *PaymentsRepo {
	return &PaymentsRepo{
		BaseRepository: baseRepo,
		db:             db,
	}
}

func (r *PaymentsRepo) Create(ctx context.Context, accountID int, sum float64, reason string) (int, error) {
	const query = `
		INSERT INTO payments (
			account_id,
			sum,
			reason
		)
		VALUES (
				   $1,
				   $2,
				   $3
			   )
		RETURNING id
`

	id := 0
	if err := r.db.QueryRowContext(ctx, query, accountID, sum, reason).Scan(&id); err != nil {
		r.Logger.Error("paymentrepo.Create.QueryError", err)

		return 0, err
	}

	return id, nil
}

func (r *PaymentsRepo) GetByAccountID(ctx context.Context, accountID int, year int) ([]paymententities.Payment, error) {
	const query = `
		SELECT 
			id,
		    account_id,
		    reason,
		    sum,
		    date
		FROM payments
		WHERE account_id = $1 %s
		ORDER BY date 
`
	where := ""
	params := []interface{}{accountID}
	if year != 0 {
		where += " AND date <= $2 AND date > $3"
		loc := time.UTC

		params = append(params, time.Date(year, 12, 31, 23, 59, 59, 59, loc))
		params = append(params, time.Date(year-1, 12, 31, 23, 59, 59, 59, loc))
	}

	resultingQuery := fmt.Sprintf(query, where)

	rows := make([]payment, 0)
	if err := r.db.SelectContext(ctx, &rows, resultingQuery, params...); err != nil {
		r.Logger.Error("paymentrepo.GetByAccountID.QueryError")

		return nil, err
	}

	result := make([]paymententities.Payment, 0, len(rows))
	for _, row := range rows {
		result = append(result, paymententities.Payment{
			ID:        row.ID,
			AccountID: row.AccountID,
			Reason:    row.Reason,
			Sum:       row.Sum,
			Date:      row.Date,
		})
	}

	return result, nil
}

func (r *PaymentsRepo) GetAll(ctx context.Context) ([]paymententities.Payment, error) {
	const query = `
		SELECT 
			id,
		    account_id,
		    reason,
		    sum,
		    date
		FROM payments
		ORDER BY date 
`

	rows := make([]payment, 0)
	if err := r.db.SelectContext(ctx, &rows, query); err != nil {
		r.Logger.Error("paymentrepo.GetByAccountID.QueryError")

		return nil, err
	}

	result := make([]paymententities.Payment, 0, len(rows))
	for _, row := range rows {
		result = append(result, paymententities.Payment{
			ID:        row.ID,
			AccountID: row.AccountID,
			Reason:    row.Reason,
			Sum:       row.Sum,
			Date:      row.Date,
		})
	}

	return result, nil
}

func (r *PaymentsRepo) GetClientPayments(ctx context.Context, accountID int) ([]paymententities.Payment, error) {
	const query = `
		SELECT 
			id,
		    account_id,
		    reason,
		    sum,
		    date
		FROM payments
		WHERE account_id = $1
		ORDER BY date 
`

	rows := make([]payment, 0)
	if err := r.db.SelectContext(ctx, &rows, query, accountID); err != nil {
		r.Logger.Error("paymentrepo.GetByAccountID.QueryError")

		return nil, err
	}

	result := make([]paymententities.Payment, 0, len(rows))
	for _, row := range rows {
		result = append(result, paymententities.Payment{
			ID:        row.ID,
			AccountID: row.AccountID,
			Reason:    row.Reason,
			Sum:       row.Sum,
			Date:      row.Date,
		})
	}

	return result, nil
}

func (r *PaymentsRepo) GetClientsPayments(ctx context.Context, accountIDs []int) ([]paymententities.Payment, error) {
	const query = `
		SELECT 
			id,
		    account_id,
		    reason,
		    sum,
		    date
		FROM payments
		WHERE account_id IN (?)
		ORDER BY date 
`

	resultingQuery, params, err := sqlx.In(query, accountIDs)
	if err != nil {
		return nil, err
	}

	resultingQuery = r.db.Rebind(resultingQuery)
	rows := make([]payment, 0)
	if err := r.db.SelectContext(ctx, &rows, resultingQuery, params...); err != nil {
		r.Logger.Error("accountrepo.GetDebtors.QueryError")

		return nil, err
	}

	result := make([]paymententities.Payment, 0, len(rows))
	for _, row := range rows {
		result = append(result, paymententities.Payment{
			ID:        row.ID,
			AccountID: row.AccountID,
			Reason:    row.Reason,
			Sum:       row.Sum,
			Date:      row.Date,
		})
	}

	return result, nil
}
