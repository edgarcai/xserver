package main

import (
	_ "net/http/pprof"

	"rtmfpserver/core"
)

func main() {
	xserver.Start()
}
