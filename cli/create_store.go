package main

import (
	"github.com/blokhost/bnb_greenfield/services"
	"log"
)

func main() {

	svc := services.BnbGreenfieldService{}
	err := svc.Start()
	if err != nil {
		panic(err)
	}

	ifo, version, err := svc.NodeInfo()
	if err != nil {
		panic(err)
	}
	log.Printf("Info: %+v\n", ifo)
	log.Printf("Version: %+v\n", version)

	addr := "0x9148243a9f405a50703ac383E975999828f90D2d"

	cid, err := svc.CreateBucket(addr)
	if err != nil {
		panic(err)
	}

	log.Printf("Bucket: %v", cid)
}
