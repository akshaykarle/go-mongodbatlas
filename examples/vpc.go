package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	dac "github.com/akshaykarle/go-http-digest-auth-client"
	"github.com/akshaykarle/mongodb-atlas-go/mongodb"
)

func main() {
	username := os.Args[1]
	password := os.Args[2]
	gid := os.Args[3]
	t := dac.NewTransport(username, password)
	httpClient := &http.Client{Transport: &t}
	client := mongodb.NewClient(httpClient)

	// VPC.List example
	peers, _, err := client.VPC.List(gid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("VPC peers list: %v\n", peers)

	// VPC.Get example
	peer, _, err := client.VPC.Get(gid, peers[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("VPC peer get: %v\n", peer)

	// VPC.Create example
	params := &Peer{
		VpcID:               "vpc-123456",
		AwsAccountID:        "abc123abc123",
		RouteTableCidrBlock: "192.168.0.0/24",
		ContainerID:         "1112222b3bf99403840e8934",
	}
	peer, _, err = client.VPC.Create(gid, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("VPC peer created: %v\n", peer)

	// VPC.Update example
	params = &Peer{RouteTableCidrBlock: "192.168.0.0/24"}
	peer, _, err = client.VPC.Update(gid, peer.ID, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("VPC peer updated: %v\n", peer)

	// VPC.Delete example
	_, err = client.VPC.Delete(gid, peer.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("VPC peer deleted: %v\n", peer)
}
