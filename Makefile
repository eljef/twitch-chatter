.PHONY: help all build deps_get deps_update gofmt lint_clean lint_install lint_run release

VERSION := "0.0.2"
NULL :=

# all runs help
all : help

# help lists out targets
help :
	$(info $(NULL))
	$(info ** Available Targets **)
	$(info $(NULL))
	$(info $(NULL)	build		- builds a development executable into the build direcotry)
	$(info $(NULL)	deps_get	- download the dependencies for this project to the vendor folder)
	$(info $(NULL)	deps_update	- update the dependencies for this project)
	$(info $(NULL)	gofmt		- runs gofmt, formatting all project source files)
	$(info $(NULL)	lint_clean	- cleans the lint tools cache)
	$(info $(NULL)	lint_run	- runs linting tools for this project)
	$(info $(NULL)	release		- builds the releases for the set version)
	$(info $(NULL))
	@:

# build builds an executable into the build directory for testing.
build :
	$(info $(NULL))
	mkdir -p build
	go build -o build/twitch-chatter ./cmd/twitch-chatter

# deps_get downloads dependencies for the project
deps_get :
	$(info $(NULL))
	go mod download
	go mod vendor
	@echo

# deps_update updates dependencies for the project
deps_update :
	$(info $(NULL))
	go get -t -u ./...
	@echo

# gofmt runs gofmt on directories
gofmt :
	$(info $(NULL))
	gofmt -w ./cmd/twitch-chatter/ ./internal/
	@echo

# lint_clean cleans the linting tools cache
lint_clean :
	$(info $(NULL))
	golangci-lint cache clean
	@echo

# lint_run runs linting tools for this project
lint_run :
	$(info $(NULL))
	golangci-lint run ./cmd/... ./internal/...
	@echo

# release builds a release build
release :
	$(info $(NULL))
	@if [ -d "release-$(VERSION)" ]; then echo Release Folder Exists. Please Remove.; echo ; exit 1; fi

	@scripts/build-releases "$(VERSION)"

	@echo
	@echo All Releases Built In release-$(VERSION) Directory
	@echo

