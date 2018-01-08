package mongodb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[],"results":[{"atlasCidrBlock":"10.2.0.0/24","id":"1112222b3bf99403840e8934","providerName":"AWS","provisioned":false,"regionName":"US_EAST_1","vpcId":null}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	containers, _, err := client.Containers.List("123")
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

func TestContainerService_Get(t *testing.T) {
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

func TestContainerService_Create(t *testing.T) {
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
	expected := params
	assert.Nil(t, err)
	assert.Equal(t, expected, container)
}

func TestContainerService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/containers/1112222b3bf99403840e8934", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{"atlasCidrBlock": "192.168.248.0/21"}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"atlasCidrBlock":"192.168.248.0/21","id":"1112222b3bf99403840e8934","providerName":"AWS","provisioned":false,"regionName":"US_EAST_1","vpcId":null}`)
	})

	client := NewClient(httpClient)
	params := &Container{AtlasCidrBlock: "192.168.248.0/21"}
	container, _, err := client.Containers.Update("123", "1112222b3bf99403840e8934", params)
	expected := params
	assert.Nil(t, err)
	assert.Equal(t, expected, container)
}

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
