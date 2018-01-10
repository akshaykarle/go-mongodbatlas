package mongodb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeerService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/peers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[],"results":[{"awsAccountId" : "999900000000","connectionId" : null,"errorStateName" : null,"id" : "1112222b3bf99403840e8934","routeTableCidrBlock" : "10.15.0.0/16","statusName" : "INITIATING","vpcId" : "vpc-abc123abc123","containerId" : "1112222b3bf99403840e8934"}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	peers, _, err := client.Peers.List("123")
	expected := []Peer{
		Peer{
			AwsAccountID:        "999900000000",
			ID:                  "1112222b3bf99403840e8934",
			RouteTableCidrBlock: "10.15.0.0/16",
			StatusName:          "INITIATING",
			VpcID:               "vpc-abc123abc123",
			ContainerID:         "1112222b3bf99403840e8934",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, peers)
}

func TestPeerService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/peers/1112222b3bf99403840e8934", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"awsAccountId" : "999900000000","connectionId" : null,"errorStateName" : null,"id" : "1112222b3bf99403840e8934","routeTableCidrBlock" : "10.15.0.0/16","statusName" : "INITIATING","vpcId" : "vpc-abc123abc123","containerId" : "1112222b3bf99403840e8934"}`)
	})

	client := NewClient(httpClient)
	peer, _, err := client.Peers.Get("123", "1112222b3bf99403840e8934")
	expected := &Peer{
		AwsAccountID:        "999900000000",
		ID:                  "1112222b3bf99403840e8934",
		RouteTableCidrBlock: "10.15.0.0/16",
		StatusName:          "INITIATING",
		VpcID:               "vpc-abc123abc123",
		ContainerID:         "1112222b3bf99403840e8934",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, peer)
}

func TestPeerService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/peers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"vpcId":               "vpc-123456",
			"awsAccountId":        "abc123abc123",
			"routeTableCidrBlock": "192.168.0.0/24",
			"containerId":         "1112222b3bf99403840e8934",
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"awsAccountId" : "abc123abc123","connectionId" : null,"errorStateName" : null,"id" : "1112222b3bf99403840e8934","routeTableCidrBlock" : "192.168.0.0/24","statusName" : "INITIATING","vpcId" : "vpc-123456","containerId" : "1112222b3bf99403840e8934"}`)
	})

	client := NewClient(httpClient)
	params := &Peer{
		VpcID:               "vpc-123456",
		AwsAccountID:        "abc123abc123",
		RouteTableCidrBlock: "192.168.0.0/24",
		ContainerID:         "1112222b3bf99403840e8934",
	}
	peer, _, err := client.Peers.Create("123", params)
	expected := &Peer{
		ID:                  "1112222b3bf99403840e8934",
		VpcID:               "vpc-123456",
		AwsAccountID:        "abc123abc123",
		RouteTableCidrBlock: "192.168.0.0/24",
		ContainerID:         "1112222b3bf99403840e8934",
		StatusName:          "INITIATING",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, peer)
}

func TestPeerService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/peers/1112222b3bf99403840e8934", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{"routeTableCidrBlock": "10.15.0.0/16"}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"awsAccountId" : "abc123abc123","connectionId" : null,"errorStateName" : null,"id" : "1112222b3bf99403840e8934","routeTableCidrBlock" : "10.15.0.0/16","statusName" : "INITIATING","vpcId" : "vpc-123456","containerId" : "1112222b3bf99403840e8934"}`)
	})

	client := NewClient(httpClient)
	params := &Peer{RouteTableCidrBlock: "10.15.0.0/16"}
	peer, _, err := client.Peers.Update("123", "1112222b3bf99403840e8934", params)
	expected := &Peer{
		ID:                  "1112222b3bf99403840e8934",
		VpcID:               "vpc-123456",
		AwsAccountID:        "abc123abc123",
		RouteTableCidrBlock: "10.15.0.0/16",
		ContainerID:         "1112222b3bf99403840e8934",
		StatusName:          "INITIATING",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, peer)
}

func TestPeerService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/peers/1112222b3bf99403840e8934", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})

	client := NewClient(httpClient)
	resp, err := client.Peers.Delete("123", "1112222b3bf99403840e8934")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
