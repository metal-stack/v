all: build test

build:
	GGO_ENABLED=0 \
	GO111MODULE=on \
	go build \
		-tags netgo \
		-ldflags "-X 'github.com/metal-stack/v.Version=v1.0.1' \
                  -X 'github.com/metal-stack/v.Revision=devel' \
                  -X 'github.com/metal-stack/v.GitSHA1=abcdef' \
                  -X 'github.com/metal-stack/v.BuildDate=1.1.1970'" \
	-o v

test:
	./v
