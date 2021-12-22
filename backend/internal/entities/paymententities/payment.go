package paymententities

import "time"

type Payment struct {
	ID        int
	AccountID int
	Reason    string
	Sum       float64
	Date      time.Time
}
