.SHELLFLAGS: -ec

# Run Petcare REST API
.PHONY: run
run:
	@echo "Starting Petcare.."
	@go run cmd/petcare/main.go 