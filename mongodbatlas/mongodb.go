package mongodbatlas

import (
	"net/http"

	"github.com/dghubble/sling"
)

const atlasURL = "https://cloud.mongodb.com/api/atlas/v1.0/"

// Client is a MongoDB Atlas client for making MongoDB API requests.
type Client struct {
	sling               *sling.Sling
	Root                *RootService
	Whitelist           *WhitelistService
	Projects            *ProjectService
	Clusters            *ClusterService
	Containers          *ContainerService
	Peers               *PeerService
	DatabaseUsers       *DatabaseUserService
	Organizations       *OrganizationService
	AlertConfigurations *AlertConfigurationService
	SnapshotSchedule    *SnapshotScheduleService
	AtlasUsers          *AtlasUserService
	PrivateIPMode       *PrivateIPModeService
}

// NewClient returns a new Client using MongoDB Atlas API Base URL
func NewClient(httpClient *http.Client) *Client {
	return NewCustomURLClient(httpClient, atlasURL)
}

// NewCustomURLClient returns a new Client using provided API Base URL
func NewCustomURLClient(httpClient *http.Client, apiURL string) *Client {
	base := sling.New().Client(httpClient).Base(apiURL)

	return &Client{
		sling:               base,
		Root:                newRootService(base.New()),
		Whitelist:           newWhitelistService(base.New()),
		Projects:            newProjectService(base.New()),
		Clusters:            newClusterService(base.New()),
		Containers:          newContainerService(base.New()),
		Peers:               newPeerService(base.New()),
		DatabaseUsers:       newDatabaseUserService(base.New()),
		Organizations:       newOrganizationService(base.New()),
		AlertConfigurations: newAlertConfigurationService(base.New()),
		SnapshotSchedule:    newSnapshotScheduleService(base.New()),
		AtlasUsers:          newAtlasUserService(base.New()),
		PrivateIPMode:       newPrivateIPModeService(base.New()),
	}
}
