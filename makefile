test-integration:
	TEST_INTEGRATION=true go test ./... -v
start:
	go run main.go serve