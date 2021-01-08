GOCMD=go
GOLINT=golint
GOFMT=gofmt
MAKE=make

.DEFAULT_GOAL := all

all :
.PHONY : all

.PHONY: install
install:
	$(shell cd /; $(GOCMD) get -u golang.org/x/lint/golint)
	$(GOCMD) mod vendor

.PHONY: test
test:
	$(GOCMD) test -mod=vendor ./...

.PHONY: lint
lint:
	./scripts/gofmt_test.sh
	$(GOLINT) ./... | grep -v vendor/ && exit 1 || exit 0
	$(GOCMD) vet -mod=vendor ./... | grep -v vendor/ && exit 1 || exit 0

.PHONY: fix
fix:
	$(GOFMT) -w .

### Utilities
.PHONY: version
version:
	git tag $(V)
	./scripts/changelog.sh
	git add ./docs/changelogs/CHANGELOG_$(V).md
	git commit --allow-empty -m "Build $(V)"
	git tag --delete $(V)
	git tag $(V)
