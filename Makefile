test:
	@echo "--> Running tests"
	@go test -v -cover ./...

install:
	@echo "--> Building"
	@go build