.PHONY: build
build:
	@./scripts/build.sh

.PHONY: run
run:
	@./scripts/run.sh

.PHONY: watch
watch:
	@modd

.PHONY: gqlgen
gqlgen:
	@./scripts/gqlgen.sh

.PHONY: clean
clean:
	@go clean --modcache
