package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reconcile/entity"
	"time"
)

func main() {
	dirPath := "sample-data"
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Gagal membuat direktori:", err)
		return
	}

	_ = os.MkdirAll(dirPath, os.ModePerm)

	//generate master system transaction
	generateCSV[entity.SystemTransaction](filepath.Join(dirPath, "system.csv"), "2006-01-02T15:04:05", 100)

	//generate bank transaction
	generateCSV[entity.SystemTransaction](filepath.Join(dirPath, "bank_1.csv"), "2006-01-02", 20)
	generateCSV[entity.SystemTransaction](filepath.Join(dirPath, "bank_2.csv"), "2006-01-02", 30)
	generateCSV[entity.SystemTransaction](filepath.Join(dirPath, "bank_3.csv"), "2006-01-02", 70)
}

func generateCSV[T entity.Transaction](filename string, dateFormat string, numRows int) (sysDataResult []T) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating system transaction file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"TrxID", "Amount", "TransactionTime"})
	for i := 0; i < numRows; i++ {
		var t T
		trxID := fmt.Sprintf("SYS-%06d", i+1)
		amount := rand.Float64() * 100
		date := time.Now().AddDate(0, 0, -rand.Intn(10))
		t.Parse(trxID, amount, date)
		sysDataResult = append(sysDataResult, t)

		writer.Write([]string{trxID, fmt.Sprintf("%.1f", amount), date.Format(dateFormat)})
	}

	return
}
