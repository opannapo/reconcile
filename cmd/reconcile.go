package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"reconcile/entity"
	"sync"
	"time"
)

func main() {
	var dirSystem, dirBank, startDate, endDate string
	flag.StringVar(&dirSystem, "dir-system", "", "Path Directory of CSV system file")
	flag.StringVar(&dirBank, "dir-bank", "", "Path Directory of CSV bank file")
	flag.StringVar(&startDate, "start-date", "", "start-date = 2025-03-09")
	flag.StringVar(&endDate, "end-date", "", "end-date = 2025-03-09")

	// Parse flag dari command line
	flag.Parse()
	if dirSystem == "" || dirBank == "" || startDate == "" || endDate == "" {
		log.Panic("Option dir-system , dir-bank , start-date , end-date required")
	}
	//validate date range

	//parse transaction system
	layoutDateOnly := "2006-01-02"
	tStart, err := time.Parse(layoutDateOnly, startDate)
	if err != nil {
		panic(err)
	}
	tEnd, err := time.Parse(layoutDateOnly, endDate)
	if err != nil {
		panic(err)
	}
	if tEnd.Before(tStart) {
		panic(fmt.Sprintf("invalid date range %s-%s", tStart, tEnd))
	}

	wg := sync.WaitGroup{}

	var tSysSlice map[string]entity.Transaction
	wg.Add(1) //system
	go func() {
		tSysSlice = parsingSystemData(dirSystem, tStart, tEnd, &wg)
	}()

	var bankTxMaps map[string]map[string]entity.Transaction
	wg.Add(1) //banmk
	go func() {
		bankTxMaps = parsingBankData(dirBank, tStart, tEnd, &wg)
	}()

	wg.Wait()

	for i, transaction := range tSysSlice {
		log.Println(i, ", system -> ", transaction)
	}

	for i, bm := range bankTxMaps {
		for ii, transaction := range bm {
			log.Println("bank file ", i, ", data bank ", ii, transaction, " -> ", transaction)
		}
	}

	mismatchesSysTxResult, mismatchesBankTxResult := findNotMatchingTx(tSysSlice, bankTxMaps)
	if len(mismatchesSysTxResult) > 0 {
		log.Println("mismatchesSysTxResult", mismatchesSysTxResult)
		generateReportMismatchesSys(mismatchesSysTxResult)
	}
	if len(mismatchesBankTxResult) > 0 {
		log.Println("mismatchesBankTxResult", mismatchesBankTxResult)
		generateReportMismatchesBank(mismatchesBankTxResult)
	}

	log.Println("completed")
}

func parsingSystemData(dir string, tStart, tEnd time.Time, wg *sync.WaitGroup) (tSysSlice map[string]entity.Transaction) {
	defer wg.Done()

	filesSystem, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error read direktori:", err)
		return
	}
	tSys := entity.Wrapper{TransactionType: "system", DateRange: []time.Time{tStart, tEnd}}
	tSysSlice, err = tSys.ParseToSlice(dir + "/" + filesSystem[0].Name())
	if err != nil {
		log.Panic(err)
	}

	return
}

func parsingBankData(dir string, tStart, tEnd time.Time, wg *sync.WaitGroup) (bankTxMaps map[string]map[string]entity.Transaction) {
	defer wg.Done()

	bankTxMaps = map[string]map[string]entity.Transaction{}

	filesBank, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error read direktori:", err)
		return
	}

	for _, file := range filesBank {
		bankFileName := file.Name()
		tBank := entity.Wrapper{TransactionType: "bank", DateRange: []time.Time{tStart, tEnd}}
		tBankSlice, err := tBank.ParseToSlice(dir + "/" + bankFileName)
		if err != nil {
			log.Panic(err)
		}

		bankTxMaps[bankFileName] = tBankSlice
	}

	return
}

func findNotMatchingTx(sysMap map[string]entity.Transaction, bankMapArrays map[string]map[string]entity.Transaction) (mismatchesSysTxResult map[string]entity.Transaction, mismatchesBankTxResult map[string]map[string]entity.Transaction) {
	scanCount := 0
	for s, _ := range sysMap {
		scanCount++
		for keyBankFile, arr := range bankMapArrays {
			for keyOnSelectedBank, _ := range arr {
				if keyOnSelectedBank == s {
					log.Println("find system key", s, "on bank file", keyBankFile, "keyOnSelectedBank", keyOnSelectedBank, "MATCH")
					delete(sysMap, s)
					delete(arr, keyOnSelectedBank)
				}
			}
		}
	}

	//set data mismatchesSysTxResult
	mismatchesSysTxResult = sysMap

	//set data mismatchesBankTxResult
	mismatchesBankTxResult = map[string]map[string]entity.Transaction{}
	for keyBankFile, arr := range bankMapArrays {
		if len(bankMapArrays[keyBankFile]) > 0 {
			mismatchesBankTxResult[keyBankFile] = arr
		}
	}

	return
}

func generateReportMismatchesSys(mismatchesSysTxResult map[string]entity.Transaction) {
	dirPathMismatches := "sample-data/mismatches"
	_ = os.MkdirAll(dirPathMismatches, os.ModePerm)

	filename := filepath.Join(dirPathMismatches, "MISMATCHES-SYSTEM.csv")
	sysFile, _ := os.Create(filename)
	defer sysFile.Close()
	writer := csv.NewWriter(sysFile)
	defer writer.Flush()

	systemHeaders := []string{"TrxID", "Amount", "Type", "TransactionTime"}
	writer.Write(systemHeaders)
	for _, m := range mismatchesSysTxResult {
		amount := m.Amount
		txType := "CREDIT"
		if amount < 0 {
			txType = "DEBIT"
			positiveNum := math.Abs(amount)
			amount = positiveNum
		}

		writer.Write([]string{
			m.ID,
			fmt.Sprintf("%.3f", amount),
			txType,
			m.Date.String(),
		})
	}
}

func generateReportMismatchesBank(mismatchesBankTxResult map[string]map[string]entity.Transaction) {
	keyGroup := []string{}
	for s, _ := range mismatchesBankTxResult {
		keyGroup = append(keyGroup, s)
	}

	for i, key := range keyGroup {
		log.Println(i, key)
		bankMap := []entity.Transaction{}
		for s, tx := range mismatchesBankTxResult {
			if s == key {
				for _, transaction := range tx {
					bankMap = append(bankMap, transaction)
				}
			}
		}

		dirPathMismatches := "sample-data/mismatches"
		_ = os.MkdirAll(dirPathMismatches, os.ModePerm)

		filename := filepath.Join(dirPathMismatches, fmt.Sprintf("MISMATCHES-%s.csv", key))
		sysFile, _ := os.Create(filename)
		defer sysFile.Close()

		writer := csv.NewWriter(sysFile)
		defer writer.Flush()

		bankHeaders := []string{"UniqueID", "Amount", "Date"}
		writer.Write(bankHeaders)
		for i, transaction := range bankMap {
			log.Println(i, transaction)

			amount := transaction.Amount
			writer.Write([]string{
				transaction.ID,
				fmt.Sprintf("%.3f", amount),
				transaction.Date.String(),
			})
		}
	}
}
