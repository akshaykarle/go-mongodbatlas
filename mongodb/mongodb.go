package mongodb

import (
	"net/http"

	"github.com/dghubble/sling"
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

const apiURL = "https://cloud.mongodb.com/api/atlas/v1.0/"

type Client struct {
	sling   *sling.Sling
	Cluster *ClusterService
}

func NewClient(username string, password string) *Client {
	t := dac.NewTransport(username, password)
	httpClient := &http.Client{Transport: &t}
	base := sling.New().Client(httpClient).Base(apiURL)
	return &Client{
		sling:   base,
		Cluster: newClusterService(base.New()),
	}
}
