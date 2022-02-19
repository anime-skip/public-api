VERSION = $(shell jq -r .version meta.json)-$(shell TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')

build:
	@docker build --build-arg VERSION=$(VERSION) . -t anime-skip/timestamp-service/server:dev
	@echo
	@docker image ls | grep "anime-skip/timestamp-service/server"
	@echo
run: 
	docker-compose up --build --abort-on-container-exit --exit-code-from timestamps_service
docker-clean:
	docker-compose up -V 
watch:
	modd
gen:
	go generate ./...
