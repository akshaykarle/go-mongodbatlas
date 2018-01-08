package mongodb

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
