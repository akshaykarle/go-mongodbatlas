package mongodb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClusterService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/clusters", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[],"results":[{"name":"test","mongoDBMajorVersion":"3.4.9"}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	clusters, _, err := client.Cluster.List("123")
	expected := []Cluster{Cluster{Name: "test", MongoDBMajorVersion: "3.4.9"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, clusters)
}

func TestClusterService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/clusters/test", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"name":"test","mongoDBMajorVersion":"3.4.9"}`)
	})

	client := NewClient(httpClient)
	cluster, _, err := client.Cluster.Get("123", "test")
	expected := &Cluster{Name: "test", MongoDBMajorVersion: "3.4.9"}
	assert.Nil(t, err)
	assert.Equal(t, expected, cluster)
}

func TestClusterService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/clusters", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"name":                "test",
			"mongoDBMajorVersion": "3.4.9",
			"replicationFactor":   float64(3),
			"diskSizeGB":          0.5,
			"backupEnabled":       false,
			"providerSettings": map[string]interface{}{
				"providerName":     "AWS",
				"regionName":       "US_EAST_1",
				"instanceSizeName": "M0",
			},
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"name":"test","mongoDBMajorVersion":"3.4.9"}`)
	})

	client := NewClient(httpClient)
	providerSettings := ProviderSettings{ProviderName: "AWS", RegionName: "US_EAST_1", InstanceSizeName: "M0"}
	params := &Cluster{Name: "test", MongoDBMajorVersion: "3.4.9", ReplicationFactor: 3, DiskSizeGB: 0.5, BackupEnabled: false, ProviderSettings: providerSettings}
	cluster, _, err := client.Cluster.Create("123", params)
	expected := params
	assert.Nil(t, err)
	assert.Equal(t, expected, cluster)
}

func TestClusterService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/clusters/test", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"name":                "test",
			"mongoDBMajorVersion": "3.4.9",
			"replicationFactor":   float64(3),
			"diskSizeGB":          float64(5),
			"backupEnabled":       false,
			"providerSettings": map[string]interface{}{
				"providerName":     "AWS",
				"regionName":       "US_EAST_1",
				"instanceSizeName": "M0",
			},
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"name":"test","mongoDBMajorVersion":"3.4.9"}`)
	})

	client := NewClient(httpClient)
	providerSettings := ProviderSettings{ProviderName: "AWS", RegionName: "US_EAST_1", InstanceSizeName: "M0"}
	params := &Cluster{Name: "test", MongoDBMajorVersion: "3.4.9", ReplicationFactor: 3, DiskSizeGB: 5, BackupEnabled: false, ProviderSettings: providerSettings}
	cluster, _, err := client.Cluster.Update("123", "test", params)
	expected := params
	assert.Nil(t, err)
	assert.Equal(t, expected, cluster)
}

func TestClusterService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/clusters/test", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})

	client := NewClient(httpClient)
	resp, err := client.Cluster.Delete("123", "test")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
