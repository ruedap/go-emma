NAME := emma
VERSION := v0.10.0.0
REVISION := $(shell git rev-parse --short HEAD)
SRCS := $(shell find . -type f -name '*.go')

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): $(SRCS)
	go build -o bin/$(NAME)

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
