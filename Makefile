build:
	@docker build . -t backend:dev
run: build
	@docker run --rm=false --network=host --env-file .env -p 8081:8081 backend:dev
watch:
	@modd

test:
	@./scripts/test.sh

gen:
	@./scripts/gqlgen.sh
clean:
	@go clean --modcache
init:
	@./scripts/init.sh
help:
	@./scripts/help.sh
