.DEFAULT_GOAL := help
SHELL := /bin/bash

#help: @ list available tasks on this project
help:
	@grep -E '[a-zA-Z\.\-]+:.*?@ .*$$' $(MAKEFILE_LIST)| tr -d '#'  | awk 'BEGIN {FS = ":.*?@ "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#build: @ build artifact
build:
	@yarn build && go generate pkg/app/server.go && go build cmd/monitoring/main.go

#clean: @ clean builds and statics
clean:
	@rm -rf main dist pkg/app/dist .cache .parcel-cache

#init: @ install project and init dependencies
init:
	@echo "[INIT] Install project and init dependencies"
	@echo "[INIT][1/3] install and setup pre-commit"
	pip install pre-commit
	pre-commit --version
	pre-commit install
	@echo "[INIT][2/3] commitlint, conventional commit, husky and newman installation"
	npm install --save-dev @commitlint/{config-conventional,cli} husky
	npx husky install
	npx husky add .husky/commit-msg "npx --no -- commitlint --edit \"$1\""
	@echo "[INIT][3/3] download development scripts"
	git clone https://github.com/b3lb/b3lb-scripts scripts

#cluster.init: @ initialize development cluster (initialize influxdb and telegraf)
cluster.init: cluster.influxdb cluster.telegraf

#cluster.start: @ start development cluster
cluster.start:
	@make -f ./scripts/Makefile cluster.start

#cluster.stop: @ stop development cluster
cluster.stop:
	@make -f ./scripts/Makefile cluster.stop

#cluster.influxdb: @ initialize influxdb database
cluster.influxdb:
	@make -f ./scripts/Makefile cluster.influxdb

#cluster.telegraf: @ initialize bigbluebutton telegraf configuration
cluster.telegraf:
	@make -f ./scripts/Makefile cluster.telegraf

#cluster.consul: @ start development cluster using consul coniguration provider
cluster.consul:
	@make -f ./scripts/Makefile cluster.consul