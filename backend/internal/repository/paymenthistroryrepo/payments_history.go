package paymenthistroryrepo

import (
	"context"
	"database/sql"
	"errors"
	"finance/internal/entities/paymententities"
	"finance/internal/repository"
	"github.com/jmoiron/sqlx"
)

type PaymentsHistoryRepo struct {
	*repository.BaseRepository

	db *sqlx.DB
}

func NewPaymentsHistoryRepo(baseRepo *repository.BaseRepository, db *sqlx.DB) *PaymentsHistoryRepo {
	return &PaymentsHistoryRepo{
		BaseRepository: baseRepo,
		db:             db,
	}
}

func (r *PaymentsHistoryRepo) Create(ctx context.Context, accountID int, paymentID int, sumBefore float64) error {
	const query = `
	INSERT INTO payments_history (
		account_id,
		payment_id,
		sum_before
	)
	VALUES (
		$1,
		$2,
		$3
	)
	`

	if _, err := r.db.ExecContext(ctx, query, accountID, paymentID, sumBefore); err != nil {
		r.Logger.Error("paymenthistrotyrepo.Create.QueryError", err)

		return err
	}

	return nil
}

func (r *PaymentsHistoryRepo) GetPayment(ctx context.Context, accountID int, paymentID int) (paymententities.PaymentHistory, error) {
	const query = `
		SELECT 
			id,
		    account_id,
		    payment_id,
		    date,
		    sum_before
		FROM payments_history
		WHERE account_id = $1 AND payment_id = $2
`

	row := paymentHistory{}
	if err := r.db.GetContext(ctx, &row, query, accountID, paymentID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return paymententities.PaymentHistory{}, repository.NoEntity
		}
		r.Logger.Error("paymenthistoryrepo.GetPayment.QueryError")

		return paymententities.PaymentHistory{}, err
	}

	return paymententities.PaymentHistory{
		ID:        row.ID,
		AccountID: row.AccountID,
		PaymentID: row.PaymentID,
		SumBefore: row.SumBefore,
		Date:      row.Date,
	}, nil
}
