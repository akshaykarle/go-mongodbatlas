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

	// Clusters.List example
	clusters, _, err := client.Clusters.List(gid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Clusters list: \n%v\n\n", clusters)

	// Clusters.Get example
	cluster, _, err := client.Clusters.Get(gid, clusters[0].Name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Clusters get: \n%v\n\n", cluster)

	// Clusters.Create example
	providerSettings := ma.ProviderSettings{ProviderName: "AWS", RegionName: "US_EAST_1", InstanceSizeName: "M10"}
	params := &ma.Cluster{Name: "test", ReplicationFactor: 3, DiskSizeGB: 1, BackupEnabled: false, ProviderSettings: providerSettings}
	cluster, _, err = client.Clusters.Create(gid, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster created: \n%v\n\n", cluster)

	// Clusters.Update example
	providerSettings = ma.ProviderSettings{ProviderName: "AWS", RegionName: "US_EAST_1", InstanceSizeName: "M10"}
	params = &ma.Cluster{Name: "test", ReplicationFactor: 3, DiskSizeGB: 5, BackupEnabled: true, ProviderSettings: providerSettings}
	cluster, _, err = client.Clusters.Update(gid, "test", params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster updated: \n%v\n\n", cluster)

	// SnapshotSchedule.Get example
	snapshotSchedule, _, err := client.SnapshotSchedule.Get(gid, "test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Snapshot Schedule Get: \n%v\n\n", snapshotSchedule)

	// SnapshotSchedule.Update example
	snapshotScheduleParams := &ma.SnapshotSchedule{
		SnapshotIntervalHours:          10,
		SnapshotRetentionDays:          4,
		DailySnapshotRetentionDays:     7,
		PointInTimeWindowHours:         24,
		WeeklySnapshotRetentionWeeks:   4,
		MonthlySnapshotRetentionMonths: 13,
	}
	snapshotSchedule, _, err = client.SnapshotSchedule.Update(gid, "test", snapshotScheduleParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Snapshot Schedule updated: \n%v\n\n", snapshotSchedule)

	// Clusters.Delete example
	_, err = client.Clusters.Delete(gid, "test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster deleted: \n%v\n\n", cluster)
}
