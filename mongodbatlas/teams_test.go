package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/orgs/123/teams", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[],"results": [{"id":"i123","name":"test"}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	teams, _, err := client.Team.List("123")
	expected := []Team{Team{Name: "test"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, teams)
}

func TestTeamService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/orgs/123/teams/test", func(w http.ResponseWriter, r *http.Request) {
		asserMethod(t, "GET", r)
		fmt.Fprintf(w, `{"name":"test"}`)
	})

	client := NewClient(httpClient)
	team, _, err := team.Teams.Get("123", "test")
	expected := &Team{Name: "test"}
	assert.Nil(t, err)
	assert.Equal(t, expected, team)
}

func TestTeamServer_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/orgs/123/teams", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{"name": "test"}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"name":"test"}`)
	})

	client := NewClient(httpClient)
	params := &Team{Name: "test"}
	team, _, err := team.Teams.Create("123", params)
	expected := &Team{Name: "test"}
	assert.Nil(t, err)
	assert.Equal(t, expected, cluster)
}

func TestTeamService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/orgs/123/teams/test", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{"name": "test"}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"name":"test"}`)
	})

	client := NewClient(httpClient)
	params := &Team{Name: "test2"}
	team, _, err := client.Teams.Update("123", "test", params)
	expected := &Team{Name: "test2"}
	assert.Nil(t, err)
	assert.Equal(t, expected, cluster)
}

func TestTeamService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/orgs/123/teams/test", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})

	client := NewClient(httpClient)
	resp, err := client.Teams.Delete("123", "test")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
