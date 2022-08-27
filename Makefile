VERSION = $(shell jq -r .version meta.json)-$(shell TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')

compile:
	@go build -o bin/server cmd/server/main.go
build:
	@docker build --build-arg GO_OPTIONS=-trimpath --build-arg VERSION=$(VERSION) --build-arg STAGE=production . -t anime-skip/public-api/server:dev
	@echo
	@docker image ls | grep "anime-skip/public-api/server"
	@echo
run: pre-run
	VERSION=$(VERSION) docker-compose up --build --abort-on-container-exit --exit-code-from public_api
run-clean: pre-run
	docker-compose up --build --abort-on-container-exit --exit-code-from public_api -V
pre-run:
	@touch .env
watch:
	modd
gen:
	go generate ./...
test: compile
	LOG_LEVEL=3 ginkgo ./...
backfill-anilist-shows:
	go run ./cmd/backfill-anilist-shows
validate-timestamps:
	go run ./cmd/validate-timestamps

get-prod-env:
	heroku config -a prod-public-api --shell | cat > .env.prod
run-prod:
	VERSION=$(VERSION) docker-compose -f docker-compose.prod.yml up --build --abort-on-container-exit --exit-code-from public_api
