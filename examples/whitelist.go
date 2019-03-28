package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	dac "github.com/akshaykarle/go-http-digest-auth-client"
	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
)

func main() {
	username := os.Args[1]
	password := os.Args[2]
	gid := os.Args[3]
	t := dac.NewTransport(username, password)
	httpClient := &http.Client{Transport: &t}
	client := ma.NewClient(httpClient)

	// Whitelist.List example
	whitelists, _, err := client.Whitelist.List(gid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Whitelist list: %v\n", whitelists)

	// Whitelist.Get example
	whitelist, _, err := client.Whitelist.Get(gid, whitelists[0].CidrBlock)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Whitelist get: %v\n", whitelist)

	// Whitelist.Create example
	params := []ma.Whitelist{
		ma.Whitelist{
			Comment:   "test",
			GroupID:   gid,
			IPAddress: "179.154.224.127",
		},
	}
	whitelists, _, err = client.Whitelist.Create(gid, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Whitelist created: %v\n", whitelists)

	// Whitelist.Delete example
	_, err = client.Whitelist.Delete(gid, whitelists[0].IPAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Whitelist deleted: %v\n", whitelist)
}
