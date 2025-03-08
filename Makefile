.PHONY: tools-generate-csv reconcile


tools-generate-csv:
	echo "generate-csv"
	go run ./cmd/tools.go

reconcile:
	echo "reconcile system path ${system}, bank path ${bank}"
	go run ./cmd/reconcile.go --dir-system=${system} --dir-bank=${bank}
