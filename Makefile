VERSION = 0.0.1

.PHONY: devel-deps
devel-deps:
	go get -u \
		golang.org/x/lint/golint \
		github.com/mattn/goveralls \
		github.com/Songmu/goxz \
		github.com/tcnksm/ghr

.PHONY: lint
lint: devel-deps
	go vet
	golint -set_exit_status

.PHONY: build
build:
	go build ./cmd/cliassert

.PHONY: clean
clean:
	rm -f cli-assert

.PHONY: test
test:
	go test

.PHONY: update-test
update-test:
	go test -update

.PHONY: crossbuild
crossbuild:
	goxz -pv=v$(VERSION) -os=linux,darwin -d=./dist/v$(VERSION) ./cmd/cliassert

.PHONY: release
release:
	ghr v$(VERSION) dist/v$(VERSION)

.PHONY: cover
cover: devel-deps
	goveralls
