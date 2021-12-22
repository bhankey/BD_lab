package paymentrepo

import "time"

type payment struct {
	ID        int       `db:"id"`
	AccountID int       `db:"account_id"`
	Reason    string    `db:"reason"`
	Sum       float64   `db:"sum"`
	Date      time.Time `db:"date"`
}
