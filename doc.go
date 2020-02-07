/*
Package v simplifies printing version information of an application.

Add the following to your main.go

	package main

	import (
		"log"
		"github.com/metal-stack/v"
	)

	func main() {
		log.Info("application", "version", v.V)
	}

And in your Makefile:


	.ONESHELL:
	SHA := $(shell git rev-parse --short=8 HEAD)
	GITVERSION := $(shell git describe --long --all)
	BUILDDATE := $(shell date -Iseconds)
	VERSION := $(or ${VERSION},devel)

	BINARY := application

	${BINARY}: clean test
		GGO_ENABLED=0 \
		GO111MODULE=on \
		go build \
		-tags netgo \
		-ldflags "-X 'github.com/metal-stack/v.Version=$(VERSION)' \
				-X 'github.com/metal-stack/v.Revision=$(GITVERSION)' \
				-X 'github.com/metal-stack/v.GitSHA1=$(SHA)' \
				-X 'github.com/metal-stack/v.BuildDate=$(BUILDDATE)'" \
		-o application

*/
package v
