package mongodb

import (
	"github.com/dghubble/sling"
)

const apiURL = "https://cloud.mongodb.com/api/atlas/v1.0/"

type Client struct {
	sling   *sling.Sling
	Cluster *ClusterService
}

func New(username string, token string) *Client {
	base := sling.New().Base(apiURL).SetBasicAuth(username, token)
	return &Client{
		sling:   base,
		Cluster: newClusterService(base.New()),
	}
}
