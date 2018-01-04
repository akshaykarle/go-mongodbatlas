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
	ProviderName        string `json:"providerName,omitempty"`
	BackingProviderName string `json:"backingProviderName,omitempty"`
	RegionName          string `json:"regionName,omitempty"`
	InstanceSizeName    string `json:"instanceSizeName,omitempty"`
}

type Cluster struct {
	Id                  string           `json:"id,omitempty"`
	GroupId             string           `json:"groupId,omitempty"`
	Name                string           `json:"name,omitempty"`
	MongoDBVersion      string           `json:"mongoDBVersion,omitempty"`
	DiskSizeGB          float64          `json:"diskSizeGB,omitempty"`
	MongoDBMajorVersion string           `json:"mongoDBMajorVersion,omitempty"`
	BackupEnabled       bool             `json:"backupEnabled,omitempty"`
	StateName           string           `json:"stateName,omitempty"`
	ReplicationFactor   int              `json:"replicationFactor,omitempty"`
	ProviderSettings    ProviderSettings `json:"providerSettings,omitempty"`
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

func (c *ClusterService) Create(gid string, cluster *Cluster) (*Cluster, *http.Response, error) {
	apiError := new(APIError)
	path := fmt.Sprintf("%s/clusters", gid)
	resp, err := c.sling.New().Post(path).BodyJSON(cluster).Receive(cluster, apiError)
	return cluster, resp, relevantError(err, *apiError)
}

func (c *ClusterService) Update(gid string, name string, cluster *Cluster) (*Cluster, *http.Response, error) {
	apiError := new(APIError)
	path := fmt.Sprintf("%s/clusters/%s", gid, name)
	resp, err := c.sling.New().Patch(path).BodyJSON(cluster).Receive(cluster, apiError)
	return cluster, resp, relevantError(err, *apiError)
}

func (c *ClusterService) Delete(gid string, name string) (*http.Response, error) {
	cluster := new(Cluster)
	apiError := new(APIError)
	path := fmt.Sprintf("%s/clusters/%s", gid, name)
	resp, err := c.sling.New().Delete(path).Receive(cluster, apiError)
	return resp, relevantError(err, *apiError)
}
