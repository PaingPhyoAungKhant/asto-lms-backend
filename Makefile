# Makefile for Asto LMS Microservice

# Default Target 
help:
	@echo "Available Commends:"
	@echo " run     - Run the API Gateway"
	@echo " test    - Run tests"
	@echo " clean   - Clean build artifacts"

# Run the API Gateway 
run:
	@echo "Starting API Gateway"
	go run cmd/gateway/main.go 
# Run Tests 
test: 
	@echo "Running Tests"
	go test -v ./...


# Clean build artifacts 
clean:
	@echo "Cleaning build artifacts"
	go clean


