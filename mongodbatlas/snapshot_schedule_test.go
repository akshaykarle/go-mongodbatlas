package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshotScheduleService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/clusters/Cluster0/snapshotSchedule", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		fmt.Fprintf(w, `{
			"groupId":"123",
			"clusterId":"456",
			"snapshotIntervalHours":6,
			"snapshotRetentionDays":2,
			"dailySnapshotRetentionDays":7,
			"pointInTimeWindowHours":24,
			"weeklySnapshotRetentionWeeks": 4,
			"monthlySnapshotRetentionMonths": 13,
			"links": []
		}`)
	})

	client := NewClient(httpClient)
	snapshotSchedule, _, err := client.SnapshotSchedule.Get("123", "Cluster0")
	expected := &SnapshotSchedule{
		GroupID:                        "123",
		ClusterID:                      "456",
		SnapshotIntervalHours:          float64(6),
		SnapshotRetentionDays:          float64(2),
		DailySnapshotRetentionDays:     float64(7),
		PointInTimeWindowHours:         float64(24),
		WeeklySnapshotRetentionWeeks:   float64(4),
		MonthlySnapshotRetentionMonths: float64(13),
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, snapshotSchedule)
}

func TestSnapshotScheduleService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/atlas/v1.0/groups/123/clusters/Cluster0/snapshotSchedule", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		w.Header().Set("Content-Type", "application/json")
		expectedBody := map[string]interface{}{
			"snapshotIntervalHours":          float64(6),
			"snapshotRetentionDays":          float64(2),
			"dailySnapshotRetentionDays":     float64(7),
			"pointInTimeWindowHours":         float64(24),
			"weeklySnapshotRetentionWeeks":   float64(4),
			"monthlySnapshotRetentionMonths": float64(13),
		}
		assertReqJSON(t, expectedBody, r)
		fmt.Fprintf(w, `{
			"groupId":"123",
			"clusterId":"456",
			"snapshotIntervalHours":6,
			"snapshotRetentionDays":2,
			"dailySnapshotRetentionDays":7,
			"pointInTimeWindowHours":24,
			"weeklySnapshotRetentionWeeks": 4,
			"monthlySnapshotRetentionMonths": 13,
			"links": []
		}`)
	})

	client := NewClient(httpClient)
	params := &SnapshotSchedule{
		SnapshotIntervalHours:          float64(6),
		SnapshotRetentionDays:          float64(2),
		DailySnapshotRetentionDays:     float64(7),
		PointInTimeWindowHours:         float64(24),
		WeeklySnapshotRetentionWeeks:   float64(4),
		MonthlySnapshotRetentionMonths: float64(13),
	}
	snapshotSchedule, _, err := client.SnapshotSchedule.Update("123", "Cluster0", params)
	expected := &SnapshotSchedule{
		GroupID:                        "123",
		ClusterID:                      "456",
		SnapshotIntervalHours:          float64(6),
		SnapshotRetentionDays:          float64(2),
		DailySnapshotRetentionDays:     float64(7),
		PointInTimeWindowHours:         float64(24),
		WeeklySnapshotRetentionWeeks:   float64(4),
		MonthlySnapshotRetentionMonths: float64(13),
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, snapshotSchedule)
}
