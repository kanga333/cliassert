build:
	go build .

clean:
	rm -f cli-assert


test:
	go test .

.PHONY: build test
