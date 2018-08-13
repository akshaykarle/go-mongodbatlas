package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlertConfigurationService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/alertConfigs", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"links":[],"results":[{"id":"533dc40ae4b00835ff81eaee","groupId":"535683b3794d371327b"}],"totalCount":1}`)
	})

	client := NewClient(httpClient)
	alertConfigurations, _, err := client.AlertConfigurations.List("123")
	expected := []AlertConfiguration{AlertConfiguration{ID: "533dc40ae4b00835ff81eaee", GroupID: "535683b3794d371327b"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, alertConfigurations)
}

func TestAlertConfigurationService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/alertConfigs/533dc40ae4b00835ff81eaee", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{"id":"533dc40ae4b00835ff81eaee","groupId":"535683b3794d371327b"}`)
	})

	client := NewClient(httpClient)
	alert, _, err := client.AlertConfigurations.Get("123", "533dc40ae4b00835ff81eaee")
	expected := &AlertConfiguration{ID: "533dc40ae4b00835ff81eaee", GroupID: "535683b3794d371327b"}
	assert.Nil(t, err)
	assert.Equal(t, expected, alert)
}

func TestAlertConfigurationService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/alertConfigs", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"metricThreshold": map[string]interface{}{
				"threshold": 12.0,
			},
			"eventTypeName": "HOST_RESTARTED",
			"enabled":       true,
		}

		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"id":"533dc40ae4b00835ff81eaee","groupId":"535683b3794d371327b"}`)
	})

	client := NewClient(httpClient)
	params := &AlertConfiguration{EventTypeName: "HOST_RESTARTED", Enabled: true}
	params.MetricThreshold.Threshold = 12
	alert, _, err := client.AlertConfigurations.Create("123", params)
	expected := &AlertConfiguration{ID: "533dc40ae4b00835ff81eaee", GroupID: "535683b3794d371327b"}
	assert.Nil(t, err)
	assert.Equal(t, expected, alert)
}

func TestAlertConfiguration_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/alertConfigs/533dc40ae4b00835ff81eaee", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"metricThreshold": map[string]interface{}{
				"threshold": 74.0,
			},
			"eventTypeName": "HOST_RESTARTED",
			"enabled":       true,
		}

		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{"id":"533dc40ae4b00835ff81eaee","groupId":"535683b3794d371327b"}`)
	})

	client := NewClient(httpClient)
	params := &AlertConfiguration{EventTypeName: "HOST_RESTARTED", Enabled: true}
	params.MetricThreshold.Threshold = 74
	alert, _, err := client.AlertConfigurations.Update("123", "533dc40ae4b00835ff81eaee", params)
	expected := &AlertConfiguration{ID: "533dc40ae4b00835ff81eaee", GroupID: "535683b3794d371327b"}
	assert.Nil(t, err)
	assert.Equal(t, expected, alert)
}

func TestAlertConfiguration_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/alertConfigs/alertConfigs", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		fmt.Fprintf(w, `{}`)
	})

	client := NewClient(httpClient)
	resp, err := client.AlertConfigurations.Delete("123", "alertConfigs")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
