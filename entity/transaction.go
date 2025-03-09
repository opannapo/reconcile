package entity

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Transaction struct {
	Key    string
	ID     string
	Amount float64
	Date   time.Time
}

type Wrapper struct {
	TransactionType string
	DateRange       []time.Time
}

func (w Wrapper) ParseToSlice(filePath string) (result []Transaction, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	layoutDateOnly := "2006-01-02"
	timeNormalized := time.Time{}
	for _, row := range rows[1:] {
		ID := row[0]
		amount, _ := strconv.ParseFloat(row[1], 64)
		date := time.Time{}

		if w.TransactionType == "system" {
			isDebit := row[2] == "DEBIT"
			if isDebit {
				amount = -amount
			}
			date, _ = time.Parse("2006-01-02T15:04:05", row[3])
			timeNormalized, _ = time.Parse("2006-01-02T15:04:05", row[3])
		} else {
			date, _ = time.Parse("2006-01-02", row[2])
		}

		if (date.After(w.DateRange[0]) || date.Equal(w.DateRange[0])) && (date.Before(w.DateRange[1]) || date.Equal(w.DateRange[1])) {
			//normalize date format to general format 2006-01-02
			result = append(result, Transaction{
				Key:    fmt.Sprintf("%s@%f", timeNormalized.Format(layoutDateOnly), amount),
				ID:     ID,
				Amount: amount,
				Date:   date,
			})
		}
	}

	return result, nil
}
