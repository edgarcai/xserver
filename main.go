package main

import "rtmfpserver/core"

func main() {
	/**
	-rtmfp=1945,1946 -ncpu=1 -parallel=32 -apps=introduction,askFor
	-manage=300 -retrans=300,500,1000,1500,1500,2500,3000,4000,5000,7500,10000,15000
	-http=6000 -debug -heartbeat=5
	**/
	xserver.Start()
}
