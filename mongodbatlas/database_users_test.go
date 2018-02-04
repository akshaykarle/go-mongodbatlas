package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseUserService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/databaseUsers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[],"results":[{"databaseName":"test","username":"test","roles":[{"databaseName":"test","roleName":"read"}]}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	databaseUsers, _, err := client.DatabaseUsers.List("123")
	expectedRoles := []Role{Role{DatabaseName: "test", RoleName: "read"}}
	expected := []DatabaseUser{DatabaseUser{DatabaseName: "test", Username: "test", Roles: expectedRoles}}
	assert.Nil(t, err)
	assert.Equal(t, expected, databaseUsers)
}

func TestDatabaseUserService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/databaseUsers/admin/test", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"databaseName":"test","username":"test","roles":[{"databaseName":"test","roleName":"read"}]}`)
	})

	client := NewClient(httpClient)
	databaseUser, _, err := client.DatabaseUsers.Get("123", "test")
	expectedRoles := []Role{Role{DatabaseName: "test", RoleName: "read"}}
	expected := &DatabaseUser{DatabaseName: "test", Username: "test", Roles: expectedRoles}
	assert.Nil(t, err)
	assert.Equal(t, expected, databaseUser)
}

func TestDatabaseUserService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/databaseUsers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"username":     "test",
			"password":     "test",
			"databaseName": "test",
			"roles": []interface{}{map[string]interface{}{
				"databaseName": "test",
				"roleName":     "read",
			}},
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"databaseName":"test","username":"test","roles":[{"databaseName":"test","roleName":"read"}]}`)
	})

	client := NewClient(httpClient)
	roles := []Role{Role{DatabaseName: "test", RoleName: "read"}}
	params := &DatabaseUser{DatabaseName: "test", Username: "test", Password: "test", Roles: roles}
	databaseUser, _, err := client.DatabaseUsers.Create("123", params)
	expected := &DatabaseUser{DatabaseName: "test", Username: "test", Roles: roles}
	assert.Nil(t, err)
	assert.Equal(t, expected, databaseUser)
}

func TestDatabaseUserService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/databaseUsers/admin/test", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"password": "secure",
			"roles": []interface{}{map[string]interface{}{
				"databaseName": "test",
				"roleName":     "readWrite",
			}},
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"databaseName":"test","username":"test","roles":[{"databaseName":"test","roleName":"readWrite"}]}`)
	})

	client := NewClient(httpClient)
	roles := []Role{Role{DatabaseName: "test", RoleName: "readWrite"}}
	params := &DatabaseUser{Password: "secure", Roles: roles}
	databaseUser, _, err := client.DatabaseUsers.Update("123", "test", params)
	expected := &DatabaseUser{DatabaseName: "test", Username: "test", Roles: roles}
	assert.Nil(t, err)
	assert.Equal(t, expected, databaseUser)
}

func TestDatabaseUserService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/databaseUsers/admin/test", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})

	client := NewClient(httpClient)
	resp, err := client.DatabaseUsers.Delete("123", "test")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
