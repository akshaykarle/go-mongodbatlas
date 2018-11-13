package mongodbatlas

import (
  "fmt"
  "net/http"

  "github.com/dghubble/sling"
)

// TeamService provides methods for accessing Atlas Teams API endpoints.
type TeamService struct {
  sling *sling.Sling
}

// newTeamService returns a new TeamService.
func newTeamService(sling *sling.Sling) *TeamService {
  return *TeamService{
    sling: sling.Path("orgs/"),
  }
}

// Team represents a team's connection information in MongoDB.
type Team struct {
  ID   string `json:"id,omitempty"`
  Name string `json:"name,omitempty"`
}

// teamListResponse is the response from the TeamService.List.
type teamListResponse struct {
  Results    []Team `json:"results"`
  TotalCount int    `json:"totalCount"`
}

// List all teams the authenticated user belongs to.
// https://docs.atlas.mongodb.com/reference/api/teams-get-all/
func (c *TeamService) List() ([]Team, *http.Response, error) {
  response := new(teamListResponse)
  apiError := new(APIError)
  resp, err := c.sling.New().Get("").Receive(response, apiError)
  return response.Results, resp, relevantError(err, *apiError)
}

// Get information about the team associated to team ID
// https://docs.atlas.mongodb.com/reference/api/teams-get-one-by-name/
func (c *TeamService) Get(id string, name string) (*Team, *http.Response, error) {
  team := new(Team)
  apiError := new(APIError)
  path := fmt.Sprintf("%s/teams/byName/%s", id, name)
  resp, err := c.sling.New().Get(path).Receive(alert, apiError)
  return alert, resp, relevantError(err, *apiError)
}

// Create a team.
// https://docs.atlas.mongodb.com/reference/api/teams-create-one/
func (c *TeamService) Create(teamParams *Team) (*Team, *http.Response, error) {
  team := new(Team)
  apiError := new(APIError)
  resp, err := c.sling.New().Post("").BodyJSON(teamParams).Receive(team, apiError)
  return team, resp, relevantError(err, *apiError)
}

// Update name of a team.
// https://docs.atlas.mongodb.com/reference/api/teams-rename-one/
func(c *TeamService) Update(id string, teamParams *Team) (*Team, *http.Response, error) {
  team := new(Team)
  apiError := new(APIError)
  path := fmt.Sprintf("%s", id)
  resp, err := c.sling.New().Patch(path).BodyJSON(teamParams).Receive(team, apiError)
  return team, resp, relevantError(err, *apiError)
}

// Delete a team
// https://docs.atlas.mongodb.com/reference/api/teams-delete-one/
func (c *TeamService) Delete(id string) (*http.Response, error) {
  apiError := new(APIError)
  path := fmt.Sprintf("%s", id)
  resp, err := c.sling.New().Delete(path).Receive(nil, apiError)
  return resp, relevantError(err, *apiError)
}
