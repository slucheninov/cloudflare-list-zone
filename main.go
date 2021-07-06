package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func main() {
	// Init vars
	//var idzone = make(map[string]string)
	// date UTC now

	// Construct a new API object and get env api key
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CF_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	// Most API calls require a Context
	ctx := context.Background()

	// GET all id domain
	z, err := api.ListZones(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", z)
	for _, zi := range z {
		if zi.Status == "active" && zi.Plan.IsSubscribed {
			//idzone[zi.ID] = zi.Name
			//fmt.Printf("%v %v\n", zi.ID, zi.Name)
			fmt.Printf("%v %v\n", zi.Name, zi.ID)

		}
	}

}
