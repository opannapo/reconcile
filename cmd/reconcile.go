package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	findNotMatchingTx(tSysSlice, bankTxMaps)
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

func findNotMatchingTx(sysMap map[string]entity.Transaction, bankMapArrays map[string]map[string]entity.Transaction) (sysTxResult map[string]entity.Transaction, bankTxResult map[string]map[string]entity.Transaction) {
	scanCount := 0
	for s, _ := range sysMap {
		scanCount++
		for keyBankFile, arr := range bankMapArrays {
			for keyOnSelectedBank, _ := range arr {
				log.Println("find system key", s, "on bank file", keyBankFile, "keyOnSelectedBank", keyOnSelectedBank, "scanCount", scanCount, "len(sysMap)", len(sysMap))
				if keyOnSelectedBank == s {
					log.Println("find system key", s, "on bank file", keyBankFile, "keyOnSelectedBank", keyOnSelectedBank, "MATCH")
					delete(sysMap, s)
					delete(arr, keyOnSelectedBank)
				}
			}
			log.Println("OnProgress current data sysMap count", len(sysMap))
		}
	}

	log.Println("Final current data sysMap", sysMap)
	log.Println("Final current data bankMapArrays", bankMapArrays)

	return
}

func findNotMatching(mainArray []int, arrays ...[]int) map[string][]int {
	allOtherValues := make(map[int]bool)
	for _, arr := range arrays {
		for _, v := range arr {
			allOtherValues[v] = true
		}
	}

	mainNotMatch := []int{}
	for _, v := range mainArray {
		if !allOtherValues[v] {
			mainNotMatch = append(mainNotMatch, v)
		}
	}

	results := map[string][]int{"ArrayMain_not_match": mainNotMatch}

	for i, arr := range arrays {
		notMatch := []int{}
		for _, v := range arr {
			if !contains(mainArray, v) {
				notMatch = append(notMatch, v)
			}
		}
		results[fmt.Sprintf("Array%d_not_match", i+2)] = notMatch
	}

	return results
}

func contains(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
