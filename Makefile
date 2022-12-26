build:
	go build -o bin/http ./cmd/claim-api/http/main.go
	go build -o bin/lambda ./cmd/claim-api/lambda/main.go

test-unit:
	go test -v `go list ./internal/...` -tags=unit

test-integration: 
	go test -v `go list ./internal/...` -tags=integration

run-local:
	go run ./cmd/claim-api/http/main.go

run-lambda:
	go run ./cmd/claim-api/lambda/main.go
