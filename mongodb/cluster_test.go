package mongodb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClusterService_List(t *testing.T) {
	_, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/groups/123/clusters", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `[{"name":"test","mongoDBMajorVersion":"3.4.9"}]`)
	})

	client := New("test", "password")
	clusters, _, err := client.Cluster.List("123")
	expected := []Cluster{Cluster{Name: "test", MongoDBMajorVersion: "3.4.9"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, clusters)
}
