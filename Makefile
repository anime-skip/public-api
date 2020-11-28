build:
	@docker build . -t anime-skip/backend/api:dev
run:
	@docker build -q . -t anime-skip/backend/api:dev
	@docker run --network=host --env-file .env -p 8081:8081 anime-skip/backend/api:dev
watch:
	@modd

test:
	@./scripts/test.sh

gen:
	@./scripts/gqlgen.sh
clean:
	@go clean --modcache
	@go mod download
init:
	@./scripts/init.sh
help:
	@./scripts/help.sh
