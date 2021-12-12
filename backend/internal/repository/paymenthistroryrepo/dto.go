package paymenthistroryrepo

import "time"

type paymentHistory struct {
	ID        int       `db:"id"`
	AccountID int       `db:"account_id"`
	PaymentID int       `db:"payment_id"`
	SumBefore float64   `db:"sum_before"`
	Date      time.Time `db:"date"`
}
