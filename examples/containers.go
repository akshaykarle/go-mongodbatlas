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

	// Container.List example
	containers, _, err := client.Containers.List(gid, "AWS")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Container containers list: %v\n", containers)

	// Container.Get example
	container, _, err := client.Containers.Get(gid, containers[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Container container get: %v\n", container)

	// Container.Create example
	params := &ma.Container{
		AtlasCidrBlock: "10.1.0.0/21",
		ProviderName:   "AWS",
		RegionName:     "US_EAST_1",
	}
	container, _, err = client.Containers.Create(gid, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Container container created: %v\n", container)

	// Container.Update example
	params = &ma.Container{
		AtlasCidrBlock: "192.168.248.0/21",
		ProviderName:   "AWS",
	}
	container, _, err = client.Containers.Update(gid, container.ID, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Container container updated: %v\n", container)

	// Container.Delete example
	_, err = client.Containers.Delete(gid, container.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Container container deleted: %v\n", container)
}
