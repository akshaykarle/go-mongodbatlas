package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerService_List_Aws(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[],"results":[{"atlasCidrBlock":"10.2.0.0/24","id":"1112222b3bf99403840e8934","providerName":"AWS","provisioned":false,"regionName":"US_EAST_1","vpcId":null}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	containers, _, err := client.Containers.List("123", "AWS")
	expected := []Container{
		Container{
			ID:             "1112222b3bf99403840e8934",
			AtlasCidrBlock: "10.2.0.0/24",
			ProviderName:   "AWS",
			RegionName:     "US_EAST_1",
			Provisioned:    false,
			VpcID:          "",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, containers)
}

func TestContainerService_List_Gcp(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[],"results":[{"atlasCidrBlock":"10.2.0.0/24","id":"1112222b3bf99403840e8934","providerName":"GCP","provisioned":false,"gcpProjectId":null,"networkName":null}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	containers, _, err := client.Containers.List("123", "GCP")
	expected := []Container{
		Container{
			ID:             "1112222b3bf99403840e8934",
			AtlasCidrBlock: "10.2.0.0/24",
			ProviderName:   "GCP",
			GcpProjectID:   "",
			Provisioned:    false,
			NetworkName:    "",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, containers)
}

func TestContainerService_Get_Aws(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers/1112222b3bf99403840e8934", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"atlasCidrBlock":"10.2.0.0/24","id":"1112222b3bf99403840e8934","providerName":"AWS","provisioned":false,"regionName":"US_EAST_1","vpcId":null}`)
	})

	client := NewClient(httpClient)
	container, _, err := client.Containers.Get("123", "1112222b3bf99403840e8934")
	expected := &Container{
		ID:             "1112222b3bf99403840e8934",
		AtlasCidrBlock: "10.2.0.0/24",
		ProviderName:   "AWS",
		RegionName:     "US_EAST_1",
		Provisioned:    false,
		VpcID:          "",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, container)
}

func TestContainerService_Get_Gcp(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers/1112222b3bf99403840e8934", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"atlasCidrBlock":"10.2.0.0/24","id":"1112222b3bf99403840e8934","providerName":"GCP","provisioned":false,"gcpProjectId":null,"networkName":null}`)
	})

	client := NewClient(httpClient)
	container, _, err := client.Containers.Get("123", "1112222b3bf99403840e8934")
	expected := &Container{
		ID:             "1112222b3bf99403840e8934",
		AtlasCidrBlock: "10.2.0.0/24",
		ProviderName:   "GCP",
		GcpProjectID:   "",
		Provisioned:    false,
		NetworkName:    "",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, container)
}

func TestContainerService_Create_Aws(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"atlasCidrBlock": "192.168.248.0/21",
			"providerName":   "AWS",
			"regionName":     "US_EAST_1",
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"atlasCidrBlock":"192.168.248.0/21","id":"1112222b3bf99403840e8934","providerName":"AWS","provisioned":false,"regionName":"US_EAST_1","vpcId":null}`)
	})

	client := NewClient(httpClient)
	params := &Container{
		AtlasCidrBlock: "192.168.248.0/21",
		ProviderName:   "AWS",
		RegionName:     "US_EAST_1",
	}
	container, _, err := client.Containers.Create("123", params)
	expected := &Container{
		ID:             "1112222b3bf99403840e8934",
		AtlasCidrBlock: "192.168.248.0/21",
		ProviderName:   "AWS",
		RegionName:     "US_EAST_1",
		Provisioned:    false,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, container)
}

func TestContainerService_Create_Gcp(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"atlasCidrBlock": "192.168.248.0/21",
			"providerName":   "GCP",
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"atlasCidrBlock":"192.168.248.0/21","id":"1112222b3bf99403840e8934","providerName":"GCP","provisioned":false,"networkName":null,"gcpProjectId":null}`)
	})

	client := NewClient(httpClient)
	params := &Container{
		AtlasCidrBlock: "192.168.248.0/21",
		ProviderName:   "GCP",
	}
	container, _, err := client.Containers.Create("123", params)
	expected := &Container{
		ID:             "1112222b3bf99403840e8934",
		AtlasCidrBlock: "192.168.248.0/21",
		ProviderName:   "GCP",
		NetworkName:    "",
		Provisioned:    false,
		GcpProjectID:   "",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, container)
}

func TestContainerService_Update_Aws(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers/1112222b3bf99403840e8934", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{"atlasCidrBlock": "192.168.268.0/21"}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"atlasCidrBlock":"192.168.268.0/21","id":"1112222b3bf99403840e8934","providerName":"AWS","provisioned":false,"regionName":"US_EAST_1","vpcId":null}`)
	})

	client := NewClient(httpClient)
	params := &Container{AtlasCidrBlock: "192.168.268.0/21"}
	container, _, err := client.Containers.Update("123", "1112222b3bf99403840e8934", params)
	expected := &Container{
		ID:             "1112222b3bf99403840e8934",
		AtlasCidrBlock: "192.168.268.0/21",
		ProviderName:   "AWS",
		RegionName:     "US_EAST_1",
		Provisioned:    false,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, container)
}

func TestContainerService_Update_Gcp(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers/1112222b3bf99403840e8934", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{"atlasCidrBlock": "192.168.268.0/21"}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"atlasCidrBlock":"192.168.268.0/21","id":"1112222b3bf99403840e8934","providerName":"GCP","provisioned":false,"gcpProjectId":null,"networkName":null}`)
	})

	client := NewClient(httpClient)
	params := &Container{AtlasCidrBlock: "192.168.268.0/21"}
	container, _, err := client.Containers.Update("123", "1112222b3bf99403840e8934", params)
	expected := &Container{
		ID:             "1112222b3bf99403840e8934",
		AtlasCidrBlock: "192.168.268.0/21",
		ProviderName:   "GCP",
		GcpProjectID:   "",
		Provisioned:    false,
		NetworkName:    "",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, container)
}

// Only a single Test is required here as, the call does not change dependant on provider
// neither does the return. However, it is worth noting there does not seem to be official
// documentation on the deletion of containers.
func TestContainerService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers/1112222b3bf99403840e8934", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})

	client := NewClient(httpClient)
	resp, err := client.Containers.Delete("123", "1112222b3bf99403840e8934")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
