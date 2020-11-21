.SHELLFLAGS: -ec

# Run Petcare REST API
.PHONY: run
run:
	@echo "Starting Petcare.."
	@go run cmd/petcare/main.go --port 80

# Run unit tests of petcare project
.PHONY: test
test:
	@echo "Testing api.."
	@go test ./...
