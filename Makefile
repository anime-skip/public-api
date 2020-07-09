.PHONY: build
build:
	@./scripts/build.sh

.PHONY: run
run:
	@./scripts/run.sh

.PHONY: watch
watch:
	@modd

.PHONY: gen
gen:
	@./scripts/gqlgen.sh

.PHONY: clean
clean:
	@go clean --modcache

.PHONY: init
init:
	@./scripts/init.sh

.PHONY: help
help:
	@./scripts/help.sh
