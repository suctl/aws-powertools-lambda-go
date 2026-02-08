test:
	rm -f coverage.out
	go test -coverprofile=coverage.out --count=1 -cover -v ./... 

view-report: test
	go tool cover -html=coverage.out
