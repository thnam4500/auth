build:
	@go build -o bin/go-auth
run: build
	@./bin/go-auth	
test:
	@go test -v ./..
