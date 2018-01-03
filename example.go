package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akshaykarle/mongodb-atlas-go/mongodb"
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

func main() {
	username := os.Args[1]
	password := os.Args[2]
	gid := os.Args[3]
	t := dac.NewTransport(username, password)
	httpClient := &http.Client{Transport: &t}
	client := mongodb.NewClient(httpClient)
	clusters, resp, err := client.Cluster.List(gid)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp)
	}
	fmt.Println(clusters)
}
