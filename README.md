# v

[![GoDoc](https://godoc.org/github.com/metal-pod/v?status.svg)](https://godoc.org/github.com/metal-pod/v)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/metal-pod/v/blob/master/LICENSE)


v simplifies printing version information of an application.

## Quickstart

Add the following to your main.go


```go
package main

import (
    "log"
    "github.com/metal-pod/v"
)

func main() {
    log.Info("application", "version", v.V)
}

```

Modify you build target in the Makefile:

```Makefile
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
    -ldflags "-X 'github.com/metal-pod/v.Version=$(VERSION)' \
              -X 'github.com/metal-pod/v.Revision=$(GITVERSION)' \
              -X 'github.com/metal-pod/v.GitSHA1=$(SHA)' \
              -X 'github.com/metal-pod/v.BuildDate=$(BUILDDATE)'" \
    -o application

```

Compile and run:

```bash
make application
```

Expected output:

```bash
[...]
INFO[10-29|14:22:34] application   version="devel (0b016992), heads/master-0-g0b01699, 2019-05-29T14:22:26+01:00"
```
