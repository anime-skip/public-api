build:
	@docker build . --build-arg DEV=true -t anime-skip/backend/api:dev
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
