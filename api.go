package goriot

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type RiotAPI struct {
	key    string
	json   jsoniter.API
	client *http.Client
}

func NewRiotAPI(key string) *RiotAPI {
	return &RiotAPI{key, jsoniter.ConfigFastest, &http.Client{}}
}

func (api *RiotAPI) fetch(endpoint string, result interface{}) error {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Riot-Token", api.key)

	res, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("api returned status code %d: %s", res.StatusCode, res.Status)
	}

	return api.json.NewDecoder(res.Body).Decode(result)
}

func (api *RiotAPI) SupportedRegions() []string {
	var regions []string
	for region := range routingValues {
		regions = append(regions, region.String())
	}
	return regions
}
