package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akshaykarle/mongodb-atlas-go/mongodb"
)

func main() {
	username := os.Args[1]
	password := os.Args[2]
	gid := os.Args[3]
	client := mongodb.New(username, password)
	clusters, resp, err := client.Cluster.List(gid)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp)
	}
	fmt.Println(clusters)
}
