package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	wol "github.com/ghthor/gowol"
)

func main() {

	macPtr := flag.String("mac", "", "mac address in format aa:bb:cc to send WOL packet to")
	broadcastPtr := flag.String("bcast", "", "broadcast address to send WOL packet to")

	flag.Parse()

	// check for required arguments
	if *macPtr == "" || *broadcastPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// check if provided mac is valid
	if _, err := net.ParseMAC(*macPtr); err != nil {
		log.Fatalf("Error: %s", err)
	}

	// check if provided broadcast is reachable
	err := checkBroadcast(*broadcastPtr)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	log.Printf("Sending WOL packet to '%s'\n", *macPtr)
	err = wol.MagicWake(*macPtr, *broadcastPtr)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func checkBroadcast(broadcastString string) error {
	broadcast := net.ParseIP(broadcastString)
	if broadcast == nil {
		return errors.New(fmt.Sprintf("%s not a valid IP address", broadcastString))
	}
	log.Printf("Using '%s' as broadcast", broadcast)

	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Printf("Error occured checking interface '%v'. Error: %s\n", i, err.Error())
		}
		for _, addr := range addrs {
			if ipaddr, ok := addr.(*net.IPNet); ok {
				if ipaddr.Contains(broadcast) {
					return nil
				}
			}
		}
	}

	return errors.New(fmt.Sprintf("Could not find interface containing broadcast %s", broadcast))
}
