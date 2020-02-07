package main

import (
	"github.com/metal-stack/v"
)

func main() {
	const expected = "v0.0.1 (abcdef), devel, 1.1.1970"
	version := v.V.String()
	if version != expected {
		panic("expected:" + expected + " got:" + version)
	}
}
