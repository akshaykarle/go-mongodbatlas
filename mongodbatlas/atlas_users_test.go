package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtlasUserService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/users/789", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{
			"emailAddress": "john.doe@example.com",
			"firstName": "John",
			"id": "789",
			"lastName": "Doe",
			"links": [],
			"mobileNumber" : "2125550198",
			"roles": [
				{
					"orgId": "111",
					"roleName": "ORG_OWNER"
				},
				{
					"groupId": "222",
					"roleName": "GROUP_OWNER"
				}
			],
			"teamIds": [
				"333"
			],
			"username": "john.doe@example.com"
		}`)
	})

	client := NewClient(httpClient)
	atlasUser, _, err := client.AtlasUsers.Get("789")
	expectedRoles := []AtlasRole{{OrgID: "111", RoleName: "ORG_OWNER"}, {GroupID: "222", RoleName: "GROUP_OWNER"}}
	expected := &AtlasUser{
		EmailAddress: "john.doe@example.com",
		FirstName:    "John",
		LastName:     "Doe",
		ID:           "789",
		MobileNumber: "2125550198",
		Roles:        expectedRoles,
		TeamIDs:      []string{"333"},
		Username:     "john.doe@example.com",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, atlasUser)
}

func TestAtlasUserService_GetByName(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/users/byName/john.doe@example.com", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{
			"emailAddress": "john.doe@example.com",
			"firstName": "John",
			"id": "789",
			"lastName": "Doe",
			"links": [],
			"mobileNumber": "2125550198",
			"roles": [
				{
					"orgId": "111",
					"roleName": "ORG_OWNER"
				},
				{
					"groupId": "222",
					"roleName": "GROUP_OWNER"
				}
			],
			"teamIds": [
				"333"
			],
			"username": "john.doe@example.com"
		}`)
	})

	client := NewClient(httpClient)
	atlasUser, _, err := client.AtlasUsers.GetByName("john.doe@example.com")
	expectedRoles := []AtlasRole{{OrgID: "111", RoleName: "ORG_OWNER"}, {GroupID: "222", RoleName: "GROUP_OWNER"}}
	expected := &AtlasUser{
		EmailAddress: "john.doe@example.com",
		FirstName:    "John",
		LastName:     "Doe",
		ID:           "789",
		MobileNumber: "2125550198",
		Roles:        expectedRoles,
		TeamIDs:      []string{"333"},
		Username:     "john.doe@example.com",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, atlasUser)
}

func TestAtlasUserService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/users/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"username":     "john.doe@example.com",
			"emailAddress": "john.doe@example.com",
			"firstName":    "John",
			"lastName":     "Doe",
			"password":     "myPassword1@",
			"mobileNumber": "2125550198",
			"roles": []interface{}{map[string]interface{}{
				"orgId":    "111",
				"roleName": "ORG_OWNER",
			}, map[string]interface{}{
				"groupId":  "222",
				"roleName": "GROUP_OWNER",
			}},
			"country": "US",
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{
			"username": "john.doe@example.com",
			"emailAddress": "john.doe@example.com",
			"firstName": "John",
			"lastName": "Doe",
			"id": "789",
			"mobileNumber" : "2125550198",
			"links": [],
			"roles": [
				{
					"orgId": "111",
					"roleName": "ORG_OWNER"
				},
				{
					"groupId": "222",
					"roleName": "GROUP_OWNER"
				}
			],
			"teamIds": []
		}`)
	})

	client := NewClient(httpClient)
	roles := []AtlasRole{{OrgID: "111", RoleName: "ORG_OWNER"}, {GroupID: "222", RoleName: "GROUP_OWNER"}}
	params := &AtlasUser{
		EmailAddress: "john.doe@example.com",
		Username:     "john.doe@example.com",
		FirstName:    "John",
		LastName:     "Doe",
		Password:     "myPassword1@",
		MobileNumber: "2125550198",
		Country:      "US",
		Roles:        roles,
	}
	atlasUser, _, err := client.AtlasUsers.Create(params)
	expected := &AtlasUser{
		EmailAddress: "john.doe@example.com",
		ID:           "789",
		Username:     "john.doe@example.com",
		FirstName:    "John",
		LastName:     "Doe",
		MobileNumber: "2125550198",
		Roles:        roles,
		TeamIDs:      []string{},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, atlasUser)
}

func TestAtlasUserService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/users/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"country":      "US",
			"mobileNumber": "2125550198",
			"roles": []interface{}{map[string]interface{}{
				"orgId":    "111",
				"roleName": "ORG_OWNER",
			}, map[string]interface{}{
				"groupId":  "222",
				"roleName": "GROUP_OWNER",
			}},
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{
			"username": "john.doe@example.com",
			"emailAddress": "john.doe@example.com",
			"firstName": "John",
			"lastName": "Doe",
			"id": "123",
			"mobileNumber" : "2125550198",
			"links": [],
			"roles": [
				{
					"orgId": "111",
					"roleName": "ORG_OWNER"
				},
				{
					"groupId": "222",
					"roleName": "GROUP_OWNER"
				}
			],
			"teamIds": []
		}`)
	})

	client := NewClient(httpClient)
	roles := []AtlasRole{{OrgID: "111", RoleName: "ORG_OWNER"}, {GroupID: "222", RoleName: "GROUP_OWNER"}}
	params := &AtlasUser{
		Country:      "US",
		MobileNumber: "2125550198",
		Roles:        roles,
	}
	atlasUser, _, err := client.AtlasUsers.Update("123", params)
	expected := &AtlasUser{
		EmailAddress: "john.doe@example.com",
		ID:           "123",
		Username:     "john.doe@example.com",
		FirstName:    "John",
		LastName:     "Doe",
		MobileNumber: "2125550198",
		Roles:        roles,
		TeamIDs:      []string{},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, atlasUser)
}
