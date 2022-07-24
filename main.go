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
	fmt.Fprintf(w, "| ------------------ \t | -------------------------------- \t | ------------------ \t | ---------------- \t | ---- \n")
	w.Flush()

	// Cloudflared Tunnels Basic Usage
	var tunnels []cloudflare.Tunnel
	tunnelParam := &cloudflare.TunnelListParams{
		AccountID: zones[0].Account.ID,
	}
	tunnels, err = api.Tunnels(ctx, *tunnelParam)

	if err != nil {
		log.Fatal("Error looking up tunnels")
	}

	fmt.Println("\tActive Tunnels: ")
	for i, tunnel := range tunnels {
		if tunnel.DeletedAt == nil && len(tunnel.Connections) != 0 {
			fmt.Printf("%v: %v - %v\n", i, tunnel.Name, tunnel.ID)

			for _, connection := range tunnel.Connections {
				fmt.Printf("\t%v - %v - %v\n", connection.ColoName, connection.OpenedAt, connection.ClientID)
			}
		}

	}

	// Zero Trust Rules Basic Usage
	var activeOnly bool = true
	var teamsRules []cloudflare.TeamsRule
	teamsRules, err = api.TeamsRules(ctx, zones[0].Account.ID)
	var httpRules []cloudflare.TeamsRule
	var l4Rules []cloudflare.TeamsRule
	var dnsRules []cloudflare.TeamsRule

	if err != nil {
		log.Fatal("Error looking up Zero Trust Account")
	}

	fmt.Printf("\nTeams Rules Info: (Printing only Active: %v)\n", activeOnly)

	for _, teamRule := range teamsRules {
		if teamRule.Filters[0] == cloudflare.DnsFilter {
			if !(!teamRule.Enabled && activeOnly) { // Because truth tables
				dnsRules = append(dnsRules, teamRule)
			}
		}
		if teamRule.Filters[0] == cloudflare.L4Filter {
			if !(!teamRule.Enabled && activeOnly) {
				l4Rules = append(l4Rules, teamRule)
			}
		}
		if teamRule.Filters[0] == cloudflare.HttpFilter {
			if !(!teamRule.Enabled && activeOnly) {
				httpRules = append(httpRules, teamRule)
			}
		}

	}

	fmt.Println("\tDNS Rules:")
	for i, dnsRule := range dnsRules {
		fmt.Printf("\t\t%v: %v - %v - %v \n", i+1, dnsRule.Name, dnsRule.Filters, dnsRule.Action)
		if dnsRule.Traffic != "" {
			fmt.Printf("\t\t\t%v\n", dnsRule.Traffic)
		}
		if dnsRule.Identity != "" {
			fmt.Printf("\t\t\t%v\n", dnsRule.Identity)
		}
	}

	fmt.Println("\tL4 Rules:")
	for i, l4Rule := range l4Rules {
		fmt.Printf("\t\t%v: %v - %v - %v \n", i+1, l4Rule.Name, l4Rule.Filters, l4Rule.Action)
		if l4Rule.Traffic != "" {
			fmt.Printf("\t\t\t%v\n", l4Rule.Traffic)
		}
		if l4Rule.Identity != "" {
			fmt.Printf("\t\t\t%v\n", l4Rule.Identity)
		}
	}

	fmt.Println("\tHTTP Rules:")
	for i, httpRule := range httpRules {
		fmt.Printf("\t\t%v: %v - %v - %v \n", i+1, httpRule.Name, httpRule.Filters, httpRule.Action)
		if httpRule.Traffic != "" {
			fmt.Printf("\t\t\t%v\n", httpRule.Traffic)
		}
		if httpRule.Identity != "" {
			fmt.Printf("\t\t\t%v\n", httpRule.Identity)
		}
	}

}
