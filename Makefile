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

# Deployments

deploy-staged:
	docker build . \
		--build-arg VERSION=${VERSION} \
		--build-arg VERSION_SUFFIX=${VERSION_SUFFIX} \
		-t docker.pkg.github.com/anime-skip/backend/api:staged \
		-t registry.heroku.com/staged-api-service/web
	docker push registry.heroku.com/staged-api-service/web
	heroku container:release -a staged-api-service web
deploy-prod-only:
	docker build . \
		--build-arg VERSION=${VERSION} \
		--build-arg VERSION_SUFFIX=${VERSION_SUFFIX} \
		-t docker.pkg.github.com/anime-skip/backend/api:prod \
		-t registry.heroku.com/prod-api-service/web
	docker push registry.heroku.com/prod-api-service/web
	heroku container:release -a prod-api-service web
deploy-prod-test-only:
	docker build . \
		--build-arg VERSION=${VERSION} \
		--build-arg VERSION_SUFFIX=${VERSION_SUFFIX} \
		-t docker.pkg.github.com/anime-skip/backend/api:prod \
		-t registry.heroku.com/prod-api-test-service/web
	docker push registry.heroku.com/prod-api-test-service/web
	heroku container:release -a prod-api-test-service web
deploy-prod:
	docker build . \
		--build-arg VERSION=${VERSION} \
		--build-arg VERSION_SUFFIX=${VERSION_SUFFIX} \
		-t docker.pkg.github.com/anime-skip/backend/api:prod \
		-t registry.heroku.com/prod-api-service/web \
		-t registry.heroku.com/prod-api-test-service/web
	docker push registry.heroku.com/prod-api-service/web
	docker push registry.heroku.com/prod-api-test-service/web
	heroku container:release -a prod-api-service web
	heroku container:release -a prod-api-test-service web
