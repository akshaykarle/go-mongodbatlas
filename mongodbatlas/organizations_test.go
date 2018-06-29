package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/orgs/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[],"results": [{"id":"5b20401dc0c6e334a832b26d","links":[],"name":"OrganizationFoo"}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	organizations, _, err := client.Organizations.List()
	fmt.Println(err)
	expected := []Organization{
		Organization{
			ID:   "5b20401dc0c6e334a832b26d",
			Name: "OrganizationFoo",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, organizations)
}

func TestOrganizationService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/orgs/5b20401dc0c6e334a832b26d", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"id":"5b20401dc0c6e334a832b26d", "links":[], "name":"OrganizationFoo"}`)
	})

	client := NewClient(httpClient)
	organization, _, err := client.Organizations.Get("5b20401dc0c6e334a832b26d")
	expected := &Organization{
		ID:   "5b20401dc0c6e334a832b26d",
		Name: "OrganizationFoo",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, organization)
}

func TestOrganizationService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/orgs/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"name": "OrganizationFoo",
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"id":"5b20401dc0c6e334a832b26d","name":"OrganizationFoo"}`)
	})

	client := NewClient(httpClient)
	params := &Organization{
		Name: "OrganizationFoo",
	}
	organization, _, err := client.Organizations.Create(params)
	expected := &Organization{
		ID:   "5b20401dc0c6e334a832b26d",
		Name: "OrganizationFoo",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, organization)
}

func TestOrganizationService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/orgs/5b20401dc0c6e334a832b26d", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"name": "OrganizationFoo",
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"id":"5b20401dc0c6e334a832b26d","links":[],"name":"OrganizationFoo"}`)
	})

	client := NewClient(httpClient)
	params := &Organization{
		Name: "OrganizationFoo",
	}
	organization, _, err := client.Organizations.Update("5b20401dc0c6e334a832b26d", params)
	expected := &Organization{
		ID:   "5b20401dc0c6e334a832b26d",
		Name: "OrganizationFoo",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, organization)
}

func TestOrganizationService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/orgs/5b20401dc0c6e334a832b26d", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})

	client := NewClient(httpClient)
	resp, err := client.Organizations.Delete("5b20401dc0c6e334a832b26d")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
