VERSION = $(shell jq -r .version meta.json)-$(shell TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')

compile:
	@go build -o bin/server cmd/server/main.go
build:
	@docker build --build-arg VERSION=$(VERSION) . -t anime-skip/public-api/server:dev
	@echo
	@docker image ls | grep "anime-skip/public-api/server"
	@echo
run: pre-run
	docker-compose up --build --abort-on-container-exit --exit-code-from timestamps_service
run-clean: pre-run
	docker-compose up --build --abort-on-container-exit --exit-code-from timestamps_service -V
pre-run:
	@touch .env
watch:
	modd
gen:
	go generate ./...
test: compile
	LOG_LEVEL=3 go test ./...
