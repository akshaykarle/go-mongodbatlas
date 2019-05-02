package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivateIPModeService_Enable(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/privateIpMode", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{"enabled": true}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"enabled":true}`)
	})

	client := NewClient(httpClient)
	resp, err := client.PrivateIPMode.Enable("123")

	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

// While this test passes I think it is incorrect.
func TestPrivateIPModeService_Disable(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/privateIpMode", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{}`)
	})

	client := NewClient(httpClient)
	resp, err := client.PrivateIPMode.Disable("123")

	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
