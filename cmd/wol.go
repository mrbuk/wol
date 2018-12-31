package main

import (
	"flag"
	"log"
	"os"

	wol "github.com/ghthor/gowol"
)

func main() {

	macPtr := flag.String("mac", "", "mac address in format aa:bb:cc to send WOL packet to")
	broadcastPtr := flag.String("bcast", "", "broadcast address to send WOL packet to")

	flag.Parse()

	if *macPtr == "" || *broadcastPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("Sending WOL packet to '%s'\n", *macPtr)
	err := wol.MagicWake(*macPtr, *broadcastPtr)
	if err != nil {
		log.Fatalln(err)
	}
}
