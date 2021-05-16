VERSION = $(shell jq -r .version meta.json)
VERSION_SUFFIX = $(shell TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')

build:
	@docker build --build-arg VERSION=$(VERSION) --build-arg VERSION_SUFFIX=$(VERSION_SUFFIX) . -t anime-skip/backend/api:dev
run: build
	@./scripts/run.sh
watch:
	@modd

test:
	@./scripts/test.sh

services:
	@docker-compose -f docker-compose.dev.yml up --remove-orphans
reset-services:
	@docker-compose -f docker-compose.dev.yml up --remove-orphans -V

gen:
	@./scripts/gqlgen.sh
clean:
	@go clean --modcache
	@go mod download
init:
	@./scripts/init.sh
help:
	@./scripts/help.sh
