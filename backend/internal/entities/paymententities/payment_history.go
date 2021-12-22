package paymententities

import "time"

type PaymentHistory struct {
	ID        int
	AccountID int
	PaymentID int
	SumBefore float64
	SumAfter  float64
	Date      time.Time
}
