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

	// Cluster.List example
	clusters, resp, err := client.Cluster.List(gid)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp)
	}
	fmt.Println(clusters)

	// Cluster.Get example
	cluster, resp, err := client.Cluster.Get(gid, clusters[0].Name)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp)
	}
	fmt.Println(cluster)

	// Cluster.Create example
	providerSettings := mongodb.ProviderSettings{ProviderName: "AWS", RegionName: "US_EAST_1", InstanceSizeName: "M0"}
	params := &mongodb.Cluster{Name: "test", MongoDBMajorVersion: "3.4", ReplicationFactor: 3, DiskSizeGB: 0.5, ProviderSettings: providerSettings}
	cluster, resp, err = client.Cluster.Create(gid, params)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp)
	}
	fmt.Println(cluster)

	// Cluster.Update example
	providerSettings = mongodb.ProviderSettings{ProviderName: "AWS", RegionName: "US_EAST_1", InstanceSizeName: "M0"}
	params = &mongodb.Cluster{Name: "test", MongoDBMajorVersion: "3.4", ReplicationFactor: 3, DiskSizeGB: 5, ProviderSettings: providerSettings}
	cluster, resp, err = client.Cluster.Create(gid, params)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp)
	}
	fmt.Println(cluster)

	// Cluster.Delete example
	resp, err = client.Cluster.Delete(gid, "test")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp)
	}
	fmt.Println(cluster)
}
