test:
	@rm -f coverage.out coverage.html
	@echo "Run project tests"
	@echo "------------------------------------------------------------"
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo ""
	@echo "Coverage output files Save!"
