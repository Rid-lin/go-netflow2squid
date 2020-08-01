PWD := $(shell pwd)
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename $(PWD))
GOOS := windows
GOARCH := amd64
TAG := $(VERSION)__$(GOOS)_$(GOARCH)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-w -s -X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

.PHONY: build

build: buildwothoutdebug_linux pack

buildfordebug:
	go build -o build/$(PROJECTNAME) -v ./

buildwothoutdebug:
	go build $(LDFLAGS) -o build/$(PROJECTNAME) -v ./

buildwothoutdebug_linux:
	set GOOS=linux&& go build $(LDFLAGS) -o build/$(PROJECTNAME)_$(TAG) -v ./

buildwodebug_linux:
	set GOOS=$(GOOS)&&set GOARCH=$(GOARCH)&& go build $(LDFLAGS) -o build/$(PROJECTNAME)_$(TAG) -v ./


run: build
	build/$(PROJECTNAME)
	
.DUFAULT_GOAL := build

pack:
	upx --ultra-brute build/$(PROJECTNAME)*

name:
	@echo "SHELL = $(SHELL)"
	@echo "CURDIR = $(CURDIR)"
	@echo "PWD = $(PWD)"
	@echo "VERSION = $(VERSION)"
	@echo "BUILD = $(BUILD)"
	@echo "PROJECTNAME = $(PROJECTNAME)"
	@echo "TAG = $(TAG)"