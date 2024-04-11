## This is a self-documented Makefile. For usage information, run `make help`:
##
## For more information, refer to https://suva.sh/posts/well-documented-makefiles/

VERSION     = 1.0.3
BUILD_DATE  = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

.PHONY: build
build: ## Build docker image
	docker build \
		--rm --compress \
		--build-arg VERSION="$(VERSION)" \
		--build-arg BUILD_DATE="$(BUILD_DATE)" \
		--tag imagelist/smartctl_exporter:latest \
		--tag imagelist/smartctl_exporter:$(VERSION) \
		.

.PHONY: build_bin
build_bin: ## Build bin
	go build -o smartctl_exporter -ldflags "-w -s -X main.Version=${VERSION} -X main.BuildDate=${BUILD_DATE}"

.PHONY: push
push: ## Push docker image to docker hub
	docker push imagelist/smartctl_exporter:latest
	docker push imagelist/smartctl_exporter:$(VERSION)

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
