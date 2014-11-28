test:
	@echo "--> Running tests"
	@go test -v -cover ./...

install:
	@echo "--> Installing"
	@go install