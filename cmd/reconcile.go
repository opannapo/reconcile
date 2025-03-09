package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reconcile/entity"
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

	//system
	filesSystem, err := os.ReadDir(dirSystem)
	if err != nil {
		fmt.Println("Error read direktori:", err)
		return
	}
	tSys := entity.Wrapper{TransactionType: "system", DateRange: []time.Time{tStart, tEnd}}
	tSysSlice, err := tSys.ParseToSlice(dirSystem + "/" + filesSystem[0].Name())
	if err != nil {
		log.Panic(err)
	}
	for i, transaction := range tSysSlice {
		log.Println(i, ", data -> ", transaction)
	}

	//bank
	bankTxMaps := map[string][]entity.Transaction{}
	filesBank, err := os.ReadDir(dirBank)
	if err != nil {
		fmt.Println("Error read direktori:", err)
		return
	}

	for _, file := range filesBank {
		bankFileName := file.Name()
		tBank := entity.Wrapper{TransactionType: "bank", DateRange: []time.Time{tStart, tEnd}}
		tBankSlice, err := tBank.ParseToSlice(dirBank + "/" + bankFileName)
		if err != nil {
			log.Panic(err)
		}
		bankTxMaps[bankFileName] = tBankSlice
		for i, transaction := range bankTxMaps[bankFileName] {
			log.Println(i, ", data bank ", bankFileName, " -> ", transaction)
		}
	}

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
