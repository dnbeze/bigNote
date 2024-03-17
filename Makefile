# Run your Go application without building a binary
run:
	@templ generate
	@go run cmd/*.go

# Build your Go application binary
build: clean
	@templ generate
	@go build -o $(BINARY_NAME) ./cmd/...

# Clean up the previously built binary
clean:
	@if [ -f $(BINARY_NAME) ] ; then echo "Removing $(BINARY_NAME)"; rm $(BINARY_NAME) ; fi
