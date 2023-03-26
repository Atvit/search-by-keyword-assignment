.PHONY: run
run:
	go run main.go
.PHONY: test
test:
	go test ./search
.PHONY: coverage
coverage:
	go test ./search -coverprofile=./search/coverage.out
	go tool cover -func ./search/coverage.out