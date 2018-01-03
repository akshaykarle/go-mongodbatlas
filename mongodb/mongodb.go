package mongodb

import (
	"net/http"

	"github.com/dghubble/sling"
)

const apiURL = "https://cloud.mongodb.com/api/atlas/v1.0/"

type Client struct {
	sling   *sling.Sling
	Cluster *ClusterService
}

func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(apiURL)
	return &Client{
		sling:   base,
		Cluster: newClusterService(base.New()),
	}
}
