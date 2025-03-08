package entity

import "time"

type BankTransaction struct {
	UniqueID string
	Amount   float64
	Date     time.Time
}

func (s BankTransaction) Parse(ID string, amount float64, date time.Time) {
	s.UniqueID = ID
	s.Amount = amount
	s.Date = date
}
