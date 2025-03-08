package entity

import "time"

type SystemTransaction struct {
	TrxID           string
	Amount          float64
	TransactionTime time.Time
}

func (s *SystemTransaction) Assign(ID string, amount float64, date time.Time) {
	s.TrxID = ID
	s.Amount = amount
	s.TransactionTime = date
}

func (s *SystemTransaction) Parse(ID string, amount float64, date time.Time) {
	s.TrxID = ID
	s.Amount = amount
	s.TransactionTime = date
}
