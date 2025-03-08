.PHONY: tools-generate-csv-system tools-generate-csv-bank tools-generate-csv

tools-generate-csv-system:
	echo "generate-csv-system"
	go run ./cmd/tools.go --action-type=system

tools-generate-csv-bank:
	echo "generate-csv-bank ${bank-name}"
	go run ./cmd/tools.go --action-type=bank --bank-name=${bank-name}


tools-generate-csv:
	echo "generate-csv"
	go run ./cmd/tools.go

