package entity

import "time"

type Transaction interface {
	SystemTransaction | BankTransaction
	Parse(ID string, amount float64, date time.Time)
	Assign(ID string, amount float64, date time.Time)
}
