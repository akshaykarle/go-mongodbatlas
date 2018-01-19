package mongodb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[], "results":[{"clusterCount": 2, "created":"2016-07-14T14:19:33Z", "id":"5a0a1e7e0f2912c554080ae6", "links":[], "name":"ProjectBar", "orgId":"5a0a1e7e0f2912c554080adc"}], "totalCount": 1}`)
	})

	client := NewClient(httpClient)
	projects, _, err := client.Projects.List()
	fmt.Println(err)
	expected := []Project{
		Project{
			ID:           "5a0a1e7e0f2912c554080ae6",
			Name:         "ProjectBar",
			OrgID:        "5a0a1e7e0f2912c554080adc",
			ClusterCount: 2,
			Created:      "2016-07-14T14:19:33Z",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, projects)
}

func TestProjectService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"clusterCount": 2, "created":"2016-07-14T14:19:33Z", "id":"5a0a1e7e0f2912c554080ae6", "links":[], "name":"ProjectBar", "orgId":"5a0a1e7e0f2912c554080adc"}`)
	})

	client := NewClient(httpClient)
	project, _, err := client.Projects.Get("5a0a1e7e0f2912c554080ae6")
	expected := &Project{
		ID:           "5a0a1e7e0f2912c554080ae6",
		Name:         "ProjectBar",
		OrgID:        "5a0a1e7e0f2912c554080adc",
		ClusterCount: 2,
		Created:      "2016-07-14T14:19:33Z",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, project)
}

func TestProjectService_GetByName(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/byName/ProjectBar", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"clusterCount": 2, "created":"2016-07-14T14:19:33Z", "id":"5a0a1e7e0f2912c554080ae6", "links":[], "name":"ProjectBar", "orgId":"5a0a1e7e0f2912c554080adc"}`)
	})

	client := NewClient(httpClient)
	project, _, err := client.Projects.GetByName("ProjectBar")
	expected := &Project{
		ID:           "5a0a1e7e0f2912c554080ae6",
		Name:         "ProjectBar",
		OrgID:        "5a0a1e7e0f2912c554080adc",
		ClusterCount: 2,
		Created:      "2016-07-14T14:19:33Z",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, project)
}

func TestProjectService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"orgId": "5a0a1e7e0f2912c554080adc",
			"name":  "ProjectBar",
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"clusterCount": 2, "created":"2016-07-14T14:19:33Z", "id":"5a0a1e7e0f2912c554080ae6", "links":[], "name":"ProjectBar", "orgId":"5a0a1e7e0f2912c554080adc"}`)
	})

	client := NewClient(httpClient)
	params := &Project{
		Name:  "ProjectBar",
		OrgID: "5a0a1e7e0f2912c554080adc",
	}
	project, _, err := client.Projects.Create(params)
	expected := &Project{
		ID:           "5a0a1e7e0f2912c554080ae6",
		Name:         "ProjectBar",
		OrgID:        "5a0a1e7e0f2912c554080adc",
		ClusterCount: 2,
		Created:      "2016-07-14T14:19:33Z",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, project)
}
