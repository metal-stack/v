all: build test

build:
	GGO_ENABLED=0 \
	GO111MODULE=on \
	go build \
		-tags netgo \
		-ldflags "-X 'github.com/metal-pod/v.Version=v0.0.1' \
        	      -X 'github.com/metal-pod/v.Revision=devel' \
            	  -X 'github.com/metal-pod/v.GitSHA1=abcdef' \
              	  -X 'github.com/metal-pod/v.BuildDate=1.1.1970'" \
	-o v

test:
	./v
