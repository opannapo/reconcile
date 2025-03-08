package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dirPathSystem := "sample-data/system"
	dirPathBank := "sample-data/bank"
	_ = os.MkdirAll(dirPathSystem, os.ModePerm)
	_ = os.MkdirAll(dirPathBank, os.ModePerm)

	systemHeaders := []string{"TrxID", "Amount", "TransactionTime"}
	bankHeaders := []string{"UniqueID", "Amount", "Date"}

	//generate master system transaction
	generateCSV(Option{Headers: systemHeaders, Filename: filepath.Join(dirPathSystem, "system.csv"), DateFormat: "2006-01-02T15:04:05", NumRows: 100})

	//generate bank transaction
	generateCSV(Option{Headers: bankHeaders, Filename: filepath.Join(dirPathBank, "bank_1.csv"), DateFormat: "2006-01-02", NumRows: 20})
	generateCSV(Option{Headers: bankHeaders, Filename: filepath.Join(dirPathBank, "bank_2.csv"), DateFormat: "2006-01-02", NumRows: 30})
	generateCSV(Option{Headers: bankHeaders, Filename: filepath.Join(dirPathBank, "bank_3.csv"), DateFormat: "2006-01-02", NumRows: 70})
}

type Option struct {
	Headers    []string
	Filename   string
	DateFormat string
	NumRows    int
}

func generateCSV(opt Option) {
	file, err := os.Create(opt.Filename)
	if err != nil {
		fmt.Println("Error creating system transaction file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(opt.Headers)
	for i := 0; i < opt.NumRows; i++ {
		trxID := fmt.Sprintf("SYS-%06d", i+1)
		amount := rand.Float64() * 100
		date := time.Now().AddDate(0, 0, -rand.Intn(10))
		writer.Write([]string{trxID, fmt.Sprintf("%.1f", amount), date.Format(opt.DateFormat)})
	}

	return
}
