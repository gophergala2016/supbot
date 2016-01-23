.PHONY: help run build build_pkgs install clean

help:
	@echo "run:     Run code in dev mode."
	@echo "build:   Build code."
	#@echo "test:    Run tests."
	@echo "install: Install binary."
	@echo "clean:   Clean up."

run:
	@(cd ./cmd/supbot && \
	fresh -c ../../etc/fresh-runner.conf -w=../..)

build: build_pkgs
	@mkdir -p ./bin
	@rm -f ./bin/*
	go build -o ./bin/supbot ./cmd/supbot

build_pkgs:
	go build ./...

#test:
	#go test

install: build
	go install ./...

clean:
	@rm -rf ./bin

deps:
	@glock sync -n github.com/pxue/supbot < Glockfile

update_deps:
	@glock save -n github.com/pxue/supbot > Glockfile

