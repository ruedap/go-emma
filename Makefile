NAME := emma

.DEFAULT_GOAL := build

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif
	glide install

.PHONY: esc
esc:
	esc -o data.go -pkg emma data/

.PHONY: deps
deps: glide esc

.PHONY: build
build:
	go build -o bin/$(NAME) cmd/$(NAME)/main.go

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: test
test:
	go test -cover -v `glide novendor`
