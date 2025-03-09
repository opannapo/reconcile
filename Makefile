.PHONY: tools-generate-csv reconcile


tools-generate-csv:
	echo "generate-csv"
	go run ./cmd/tools.go

# make reconcile system=/home/legion/PROJECT-CODE/PRIB/GIT/reconcile/sample-data/system bank=/home/legion/PROJECT-CODE/PRIB/GIT/reconcile/sample-data/bank start=2025-01-02 end=2025-03-01
reconcile:
	echo "reconcile system path ${system}, bank path ${bank} , start date ${start} , end date ${end}"
	go run ./cmd/reconcile.go --dir-system=${system} --dir-bank=${bank} --start-date=${start} --end-date=${end}
