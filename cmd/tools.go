package main

import (
	"encoding/csv"
	"fmt"
	"log"
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

	//generate master system transaction
	generateCSV(Option{
		SystemFilename: filepath.Join(dirPathSystem, "SYSTEM.csv"),
		NumRows:        1000,
		BankFilenames: [][]string{
			{filepath.Join(dirPathBank, "BCA.csv"), "BCA"},
			{filepath.Join(dirPathBank, "MANDIRI.csv"), "MANDIRI"},
			{filepath.Join(dirPathBank, "BRI.csv"), "BRI"},
		},
	})
}

type Master struct {
	ID     string
	Amount float64
	Type   string
	Date   time.Time
}
type Option struct {
	SystemFilename string
	BankFilenames  [][]string
	NumRows        int
}

func generateCSV(opt Option) {
	systemHeaders := []string{"TrxID", "Amount", "Type", "TransactionTime"}
	bankHeaders := []string{"UniqueID", "Amount", "Date"}

	//writer file system
	sysFile, err := os.Create(opt.SystemFilename)
	if err != nil {
		fmt.Println("Error creating system transaction file:", err)
		return
	}
	defer sysFile.Close()
	writer := csv.NewWriter(sysFile)
	defer writer.Flush()

	//writer bank
	writersBank := []*csv.Writer{}
	for _, filename := range opt.BankFilenames {
		f, err := os.Create(filename[0])
		if err != nil {
			fmt.Println("Error creating system transaction file:", err)
			return
		}
		defer f.Close()

		w := csv.NewWriter(f)
		writersBank = append(writersBank, w)
	}
	defer func() {
		for _, wb := range writersBank {
			wb.Flush()
		}
	}()

	//generate file system
	writer.Write(systemHeaders)
	sysTxTypes := []string{"DEBIT", "CREDIT"}
	master := []Master{}
	for i := 0; i < opt.NumRows; i++ {
		trxID := fmt.Sprintf("SYS%06d", i+1)
		amount := rand.Float64() * 100
		date := time.Now().AddDate(0, -rand.Intn(12), -rand.Intn(10))
		txType := sysTxTypes[rand.Intn(len(sysTxTypes))]
		master = append(master, Master{
			ID:     trxID,
			Amount: amount,
			Type:   txType,
			Date:   date,
		})

		writer.Write([]string{
			trxID,
			fmt.Sprintf("%.3f", amount),
			txType,
			date.Format("2006-01-02T15:04:05"),
		})
	}

	//generate file bank
	bankSizes := generateRandomSizes(opt.NumRows, len(opt.BankFilenames))
	log.Println("bankSizes", bankSizes)
	start := 0
	for i, size := range bankSizes {
		end := start + size
		fmt.Printf("Bank %s len=%d. Start %d - end %d\n", opt.BankFilenames[i][1], size, start, end)

		partMaster := master[start:end]
		writersBank[i].Write(bankHeaders)
		for ii, m := range partMaster {
			bankName := fmt.Sprintf("%s%d", opt.BankFilenames[i][1], ii)
			date := m.Date.Format("2006-01-02")
			amount := m.Amount
			if m.Type == "DEBIT" {
				amount = -amount
			}
			writersBank[i].Write([]string{
				bankName,
				fmt.Sprintf("%.3f", amount),
				date,
			})
		}
		start = end
	}

	return
}

func generateRandomSizes(totalLen, childCount int) []int {
	rand.Seed(time.Now().UnixNano())
	sizes := make([]int, childCount)
	remaining := totalLen
	for i := 0; i < childCount-1; i++ {
		sizes[i] = rand.Intn(remaining/(childCount-i)) + 1
		remaining -= sizes[i]
	}

	sizes[childCount-1] = remaining

	return sizes
}
