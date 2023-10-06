OUT=build
PRJ=shorty

clean:
	@rm -rf $(OUT)

format:
	@gofmt -s -w .

test:
	@go test -v ./...

build: clean format
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o $(OUT)/$(PRJ)

run: build
	@./$(OUT)/$(PRJ)
