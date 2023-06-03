package goriot

import "fmt"

const (
	summonerByNameEndpoint      = "/lol/summoner/v4/summoners/by-name/%s"
	summonerByAccountIDEndpoint = "/lol/summoner/v4/summoners/by-account/%s"
	summonerByPUUIDEndpoint     = "/lol/summoner/v4/summoners/by-puuid/%s"
)

func (api *RiotAPI) GetSummonerByName(name, region string) (*SummonerDTO, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	reg := Region(region)
	if !reg.Valid() {
		return nil, fmt.Errorf("invalid region: %s", region)
	}

	var result SummonerDTO
	endpoint := fmt.Sprintf(reg.Platform()+summonerByNameEndpoint, name)
	err := api.fetch(endpoint, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *RiotAPI) GetSummonerByAccountID(accountID, region string) (*SummonerDTO, error) {
	if accountID == "" {
		return nil, fmt.Errorf("accountID cannot be empty")
	}

	reg := Region(region)
	if !reg.Valid() {
		return nil, fmt.Errorf("invalid region: %s", region)
	}

	var result SummonerDTO
	endpoint := fmt.Sprintf(reg.Platform()+summonerByAccountIDEndpoint, accountID)
	err := api.fetch(endpoint, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *RiotAPI) GetSummonerByPUUID(summonerID, region string) (*SummonerDTO, error) {
	if summonerID == "" {
		return nil, fmt.Errorf("summonerID cannot be empty")
	}

	reg := Region(region)
	if !reg.Valid() {
		return nil, fmt.Errorf("invalid region: %s", region)
	}

	var result SummonerDTO
	endpoint := fmt.Sprintf(reg.Platform()+summonerByPUUIDEndpoint, summonerID)
	err := api.fetch(endpoint, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
