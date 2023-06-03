package goriot

import "fmt"

const (
	matchByMatchIDEndpoint = "/lol/match/v5/matches/%s"
	matchListEndpoint      = "/lol/match/v5/matches/by-puuid/%s/ids"
)

type MatchListArgs struct {
	StartTime *int64
	EndTime   *int64
	Queue     *int
	Type      *string
	Start     *int
	Count     *int
}

func (args MatchListArgs) Build() string {
	output := "?"
	if args.StartTime != nil {
		output += fmt.Sprintf("startTime=%d&", *args.StartTime)
	}
	if args.EndTime != nil {
		output += fmt.Sprintf("endTime=%d&", *args.EndTime)
	}
	if args.Queue != nil {
		output += fmt.Sprintf("queue=%d&", *args.Queue)
	}
	if args.Type != nil {
		output += fmt.Sprintf("type=%s&", *args.Type)
	}
	if args.Start != nil {
		output += fmt.Sprintf("start=%d&", *args.Start)
	}
	if args.Count != nil {
		output += fmt.Sprintf("count=%d&", *args.Count)
	}

	if len(output) == 1 {
		return ""
	}

	return output[:len(output)-1]
}

func (api *RiotAPI) GetMatchByMatchID(matchID, region string) (*MatchDTO, error) {
	if matchID == "" {
		return nil, fmt.Errorf("matchID cannot be empty")
	}

	reg := Region(region)
	if !reg.Valid() {
		return nil, fmt.Errorf("invalid region: %s", region)
	}

	var match MatchDTO
	endpoint := fmt.Sprintf(reg.Regional()+matchByMatchIDEndpoint, matchID)
	err := api.fetch(endpoint, &match)
	return &match, err
}

func (api *RiotAPI) GetMatchListByPUUID(puuid, region string, args MatchListArgs) ([]string, error) {
	if puuid == "" {
		return nil, fmt.Errorf("puuid cannot be empty")
	}

	reg := Region(region)
	if !reg.Valid() {
		return nil, fmt.Errorf("invalid region: %s", region)
	}

	var matchList []string
	endpoint := fmt.Sprintf(reg.Regional()+matchListEndpoint+args.Build(), puuid)
	err := api.fetch(endpoint, &matchList)
	return matchList, err
}
