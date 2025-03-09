package entity

import (
	"time"
)

type SystemTransaction struct {
	TrxID           string
	Amount          float64
	TransactionTime time.Time
}
