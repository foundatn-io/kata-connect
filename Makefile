.PHONY: lint
lint:
	cd proto; buf lint

.PHONY: breaking
breaking:
	cd proto; buf breaking --against ../.git#branch=master,ref=HEAD~1,subdir=proto/

.PHONY: build
build:
	cd proto; buf build

.PHONY: gen
gen:
	cd proto; buf generate