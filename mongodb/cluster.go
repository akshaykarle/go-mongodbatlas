package mongodb

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type ClusterService struct {
	sling *sling.Sling
}

func newClusterService(sling *sling.Sling) *ClusterService {
	return &ClusterService{
		sling: sling.Path("groups/"),
	}
}

type ProviderSettings struct {
	providerName        string `json:"providerName"`
	backingProviderName string `json:"backingProviderName"`
	regionName          string `json:"regionName"`
	instanceSizeName    string `json:"instanceSizeName"`
}

type Cluster struct {
	Name                string           `json:"name"`
	MongoDBMajorVersion string           `json:"mongoDBMajorVersion"`
	BackupEnabled       bool             `json:"backupEnabled"`
	ProviderSettings    ProviderSettings `json:"providerSettings"`
}

func (c *ClusterService) List(gid string) ([]Cluster, *http.Response, error) {
	clusters := new([]Cluster)
	apiError := new(APIError)
	path := fmt.Sprint("%s/clusters", gid)
	resp, err := c.sling.New().Get(path).Receive(clusters, apiError)
	return *clusters, resp, relevantError(err, *apiError)
}
