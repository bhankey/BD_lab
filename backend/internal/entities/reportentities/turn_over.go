package reportentities

import "time"

type TurnOverReport struct {
	AccountID   int
	StartingSum float64
	EndSum      float64
	MothDetails map[time.Month]MonthDetails
}
