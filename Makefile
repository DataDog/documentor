.POSIX:
.SUFFIXES:

PREFIX=/usr/local
BINDIR=bin
MANDIR=share/man
PKGDIR=./

GO=go
GIT=git
RM=rm
INSTALL=install
SCDOC=scdoc

GOBUILD_OPTS=-trimpath -buildmode=pie -ldflags '-s -w' -mod 'readonly' -modcacherw

all: build doc

pre-commit: tidy fmt lint vulnerabilities test build clean # Runs all pre-commit checks.

commit: pre-commit # Commits the changes to the repository.
	$(GIT) commit -s

push: # Pushes the changes to the repository.
	$(GIT) push origin trunk

build: # Builds an application binary.
	$(GO) build $(GOBUILD_OPTS) $(PKGDIR)

build/linux: # Builds an application binary for Linux.
	GOOS=linux GOARCH=amd64 $(GO) build $(GOBUILD_OPTS) $(PKGDIR)

build/darwin: # Builds an application binary for macOS.
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOBUILD_OPTS) $(PKGDIR)

build/windows: # Builds an application binary for Windows.
	GOOS=windows GOARCH=amd64 $(GO) build $(GOBUILD_OPTS) $(PKGDIR)

doc: # Builds the manpage.
	$(SCDOC) <doc/documentor.1.scd >documentor.1

install: # Installs the release binary.
	$(INSTALL) -d \
		$(DESTDIR)$(PREFIX)/$(BINDIR)/ \
		$(DESTDIR)$(PREFIX)/$(MANDIR)/man1/
	$(INSTALL) -pm 0755 documentor $(DESTDIR)$(PREFIX)/$(BINDIR)/
	$(INSTALL) -pm 0644 documentor.1 $(DESTDIR)$(PREFIX)/$(MANDIR)/man1/

tidy: # Updates the go.mod file to use the latest versions of all direct and indirect dependencies.
	$(GO) mod tidy

fmt: # Formats Go source files in this repository.
	$(GO) run mvdan.cc/gofumpt@latest -e -extra -w .

lint: # Runs golangci-lint using the config at the root of the repository.
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./...

vulnerabilities: # Analyzes the codebase and looks for vulnerabilities affecting it.
	$(GO) run golang.org/x/vuln/cmd/govulncheck@latest ./...

test: # Runs unit tests.
	$(GO) test -cover -race -vet all -mod readonly ./...

test/coverage: # Generates a coverage profile and open it in a browser.
	$(GO) test -coverprofile cover.out ./...
	$(GO) tool cover -html=cover.out

clean: # Cleans cache files from tests and deletes any build output.
	$(RM) -f cover.out documentor documentor.1

.PHONY: all pre-commit commit push build doc install tidy fmt lint vulnerabilities test test/coverage clean
