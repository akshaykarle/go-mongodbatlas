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
	ProviderName        string `json:"providerName"`
	BackingProviderName string `json:"backingProviderName"`
	RegionName          string `json:"regionName"`
	InstanceSizeName    string `json:"instanceSizeName"`
}

type Cluster struct {
	Id                  string           `json:"id"`
	GroupId             string           `json:"groupId"`
	Name                string           `json:"name"`
	MongoDBVersion      string           `json:"mongoDBVersion"`
	DiskSizeGB          float64          `json:"diskSizeGB"`
	MongoDBMajorVersion string           `json:"mongoDBMajorVersion"`
	BackupEnabled       bool             `json:"backupEnabled"`
	StateName           string           `json:"stateName"`
	ReplicationFactor   int              `json:"replicationFactor"`
	ProviderSettings    ProviderSettings `json:"providerSettings"`
}

type listResponse struct {
	Results    []Cluster `json:"results"`
	TotalCount int       `json:"totalCount"`
}

func (c *ClusterService) List(gid string) ([]Cluster, *http.Response, error) {
	response := new(listResponse)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/clusters", gid)
	resp, err := c.sling.New().Get(path).Receive(response, apiError)
	return response.Results, resp, relevantError(err, *apiError)
}

func (c *ClusterService) Get(gid string, name string) (*Cluster, *http.Response, error) {
	cluster := new(Cluster)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/clusters/%s", gid, name)
	resp, err := c.sling.New().Get(path).Receive(cluster, apiError)
	return cluster, resp, relevantError(err, *apiError)
}
