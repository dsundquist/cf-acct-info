package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/cloudflare/cloudflare-go"
)

func main() {
	// Construct a new API object using a global API key
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	// alternatively, you can use a scoped API token
	// api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	ctx := context.Background()

	// Print API details
	var zones []cloudflare.Zone
	zones, err = api.ListZones(ctx)

	if len(zones) == 0 || err != nil {
		log.Fatal(err)
	}

	// Print Account Information
	fmt.Printf("\nUsername   : %v \n", api.APIEmail)
	fmt.Printf("Account ID : %v \n\n", zones[0].Account.ID)

	// Using text/tabwriter for formating
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	// Print Zone Headers
	fmt.Fprintf(w, "| Sites: \t | Zone ID: \t | Plan: \t | Apex record: \t | Proxied \n")
	fmt.Fprintf(w, "| ------------------ \t | -------------------------------- \t | ------------------ \t | ---------------- \t | ---- \n")

	// Print Each Zone
	for _, zone := range zones {

		// Setup for APEX Records
		var rr cloudflare.DNSRecord
		rr.Name = zone.Name
		myRecords, _ := api.DNSRecords(ctx, zone.ID, rr)

		// Print Each
		if len(myRecords) > 0 {
			fmt.Fprintf(w, "| %v \t | %v \t | %v \t | %v \t | %t \n", zone.Name, zone.ID, zone.Plan.Name, myRecords[0].Content, *myRecords[0].Proxied)
		} else {
			fmt.Fprintf(w, "| %v \t | %v \t | %v \t | n/a \t | n/a \n", zone.Name, zone.ID, zone.Plan.Name)
		}
	}
	w.Flush()

}
