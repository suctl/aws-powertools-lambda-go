test:
	rm -f coverage.out
	go test -coverprofile=coverage.out -coverpkg=./... ./...

view-report: coverage-report
	go tool cover -html=coverage.out
