build:
	go build ./cmd/cliassert

clean:
	rm -f cli-assert

test:
	go test

.PHONY: build test
