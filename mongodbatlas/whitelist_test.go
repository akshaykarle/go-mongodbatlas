package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhitelistService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/whitelist", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[],"results":[{"cidrBlock":"10.15.0.0/16","comment":"test","groupId":"123","ipAddress":"10.15.0.1"}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	whitelist, _, err := client.Whitelist.List("123")
	expected := []Whitelist{
		Whitelist{
			CidrBlock: "10.15.0.0/16",
			Comment:   "test",
			GroupID:   "123",
			IPAddress: "10.15.0.1",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, whitelist)
}

func TestWhitelistService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/whitelist/10.15.0.0/16", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"cidrBlock":"10.15.0.0/16","comment":"test","groupId":"123","ipAddress":"10.15.0.1"}`)
	})

	client := NewClient(httpClient)
	whitelist, _, err := client.Whitelist.Get("123", "10.15.0.0/16")
	expected := &Whitelist{
		CidrBlock: "10.15.0.0/16",
		Comment:   "test",
		GroupID:   "123",
		IPAddress: "10.15.0.1",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, whitelist)
}

func TestWhitelistService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/whitelist", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := []interface{}{
			map[string]interface{}{
				"cidrBlock": "10.15.0.0/16",
				"comment":   "test",
				"groupId":   "123",
			},
		}
		assertReqJSONList(t, expectedBody, r)
		fmt.Fprintf(w, `{"links":[],"results":[{"cidrBlock":"10.15.0.0/16","comment":"test","groupId":"123","ipAddress":"10.15.0.1"}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	params := []Whitelist{
		Whitelist{
			CidrBlock: "10.15.0.0/16",
			Comment:   "test",
			GroupID:   "123",
		},
	}
	whitelist, _, err := client.Whitelist.Create("123", params)
	expected := []Whitelist{
		Whitelist{
			CidrBlock: "10.15.0.0/16",
			Comment:   "test",
			GroupID:   "123",
			IPAddress: "10.15.0.1",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, whitelist)
}

func TestWhitelistService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/whitelist/10.15.0.0/16", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})

	client := NewClient(httpClient)
	resp, err := client.Whitelist.Delete("123", "10.15.0.0/16")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
