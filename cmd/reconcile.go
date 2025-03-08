package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"reconcile/entity"
	"strconv"
	"time"
)

func main() {
	var dirSystem string
	var dirBank string
	flag.StringVar(&dirSystem, "dir-system", "", "Path Directory of CSV system file")
	flag.StringVar(&dirBank, "dir-bank", "", "Path Directory of CSV bank file")

	// Parse flag dari command line
	flag.Parse()
	if dirSystem == "" || dirBank == "" {
		log.Panic("dir-system & dir-bank required")
	}

	log.Println("dirSystem", dirSystem)
	log.Println("dirBank", dirBank)

	_, err := parseCSVSystem(dirSystem+"/system.csv", "2006-01-02T15:04:05")
	if err != nil {
		log.Panic(err)
	}

}

func parseCSVSystem(filePath string, dateFormat string) (result []entity.SystemTransaction, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range rows[1:] {
		ID := row[0]
		amount, _ := strconv.ParseFloat(row[1], 64)
		date, _ := time.Parse(dateFormat, row[2])
		t := entity.SystemTransaction{
			TrxID:           ID,
			Amount:          amount,
			TransactionTime: date,
		}
		result = append(result, t)
	}

	return result, nil
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
