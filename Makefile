benchmark:
	go test -bench=. -benchmem

coverage:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

pre-commit:
	go fmt ./...
	golangci-lint run
	make test

test:
	go test ./...