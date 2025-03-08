package main

import (
	"flag"
	"log"
)

func main() {
	var actionType, bankName string
	flag.StringVar(&actionType, "action-type", "", "Action Type: `system` or `bank`")
	flag.StringVar(&bankName, "bank-name", "", "Bank Name: Name Of Bank")
	flag.Parse()

	if actionType == "" {
		log.Panic("action-type required")
	}

	if actionType == "bank" && bankName == "" {
		log.Panic("bank-name required for action-type bank")
	}

}
