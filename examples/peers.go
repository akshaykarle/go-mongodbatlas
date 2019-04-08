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

	// Peers.List example
	peers, _, err := client.Peers.List(gid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peers peers list: %v\n", peers)

	// Peers.Get example
	peer, _, err := client.Peers.Get(gid, peers[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peers peer get: %v\n", peer)

	// Peers.Create example
	params := &Peer{
		VpcID:               "vpc-123456",
		AwsAccountID:        "abc123abc123",
		RouteTableCidrBlock: "192.168.0.0/24",
		ContainerID:         "1112222b3bf99403840e8934",
	}
	peer, _, err := client.Peers.Create(gid, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peers peer created: %v\n", peer)

	// Peers.Update example
	params := &Peer{RouteTableCidrBlock: "192.168.0.0/24"}
	peer, _, err = client.Peers.Update(gid, peer.ID, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peers peer updated: %v\n", peer)

	// Peers.Delete example
	_, err := client.Peers.Delete(gid, peer.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peers peer deleted: %v\n", peer)
}
