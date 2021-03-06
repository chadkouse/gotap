package main

import (
	"log"
	"flag"
	"./tap"
)

var prot = flag.String("prot", "tcp", "Layer 3 protocol (tcp, tcp4, tcp6)")
var dest = flag.String("dest", "localhost:11211", "Host:port to connect to")

func main() {
	flag.Parse()
	log.Printf("Connecting to %s/%s", *prot, *dest)

	var args tap.TapArguments
	args.Backfill = 131313
	args.VBuckets = []uint16{0, 2, 4}
	args.ClientName = "go_go_gadget_tap"
	client := tap.Connect(*prot, *dest, args)

	for op := range client.Feed() {
		log.Printf("Tap OP:  %s\n", op.ToString())
	}
}
