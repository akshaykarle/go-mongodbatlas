package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"appName":"MongoDB Atlas","build":"cb21a95a96766eb4a962a37dfe2c57be69ccc1e6","links":[{"href":"https://cloud.mongodb.com/api/atlas/v1.0","rel":"self"},{"href":"https://cloud.mongodb.com/api/atlas/v1.0/users/123","rel":"http://mms.mongodb.com/user"}],"throttling":false}`)
	})

	client := NewClient(httpClient)
	root, _, err := client.Root.Get()
	expected := &Root{
		AppName: "MongoDB Atlas",
		Build:   "cb21a95a96766eb4a962a37dfe2c57be69ccc1e6",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, root)
}

func TestCloudManagerRootService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/public/v1.0/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"appName":"MongoDB Cloud Manager","build":"2cf9112fd6b8f59c03a9d1d4cbf32004424e728e","links":[{"href":"https://cloud.mongodb.com/api/public/v1.0","rel":"self"},{"href":"https://cloud.mongodb.com/api/public/v1.0/orgs/000000000000000000000000/apiKeys/000000000000000000000000","rel":"http://mms.mongodb.com/apiKeys"},{"href":"https://cloud.mongodb.com/api/public/v1.0/groups","rel":"http://mms.mongodb.com/groups"},{"href":"https://cloud.mongodb.com/api/public/v1.0/admin/backup","rel":"http://mms.mongodb.com/backupAdmin"}],"throttling":false}`)
	})

	client := NewCustomURLClient(httpClient, "https://cloud.mongodb.com/api/public/v1.0/")
	root, _, err := client.Root.Get()
	expected := &Root{
		AppName: "MongoDB Cloud Manager",
		Build:   "2cf9112fd6b8f59c03a9d1d4cbf32004424e728e",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, root)
}
