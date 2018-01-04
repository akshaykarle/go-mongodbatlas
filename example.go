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
	clusters, _, err := client.Cluster.List(gid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Clusters list: %v\n", clusters)

	// Cluster.Get example
	cluster, _, err := client.Cluster.Get(gid, clusters[0].Name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster get: %v\n", cluster)

	// Cluster.Create example
	providerSettings := mongodb.ProviderSettings{ProviderName: "AWS", RegionName: "US_EAST_1", InstanceSizeName: "M10"}
	params := &mongodb.Cluster{Name: "test", ReplicationFactor: 3, DiskSizeGB: 0.5, BackupEnabled: false, ProviderSettings: providerSettings}
	cluster, _, err = client.Cluster.Create(gid, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster created: %v\n", cluster)

	// Cluster.Update example
	providerSettings = mongodb.ProviderSettings{ProviderName: "AWS", RegionName: "US_EAST_1", InstanceSizeName: "M10"}
	params = &mongodb.Cluster{Name: "test", ReplicationFactor: 3, DiskSizeGB: 5, BackupEnabled: false, ProviderSettings: providerSettings}
	cluster, _, err = client.Cluster.Create(gid, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster updated: %v\n", cluster)

	// Cluster.Delete example
	_, err = client.Cluster.Delete(gid, "test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster deleted: %v\n", cluster)
}
