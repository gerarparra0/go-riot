package goriot

import (
	"reflect"
	"testing"
)

const (
	testMatchID = "NA1_4673336201"
	validPUUID  = "xVojEdk4AjaHfyHn5NLULbjinNRd9sLff3Jzd2is1VRF5AE6_z9RjirxUsUhvoQDhAnVgNvkSX6BMA"
)

func TestRiotAPI_MatchByMatchID(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		matchID string
		region  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *MatchDTO
		wantErr bool
	}{
		{
			name: "MatchByMatchID - valid matchID",
			fields: fields{
				key: devKey,
			},
			args: args{
				matchID: testMatchID,
				region:  "NA",
			},
			want:    &MatchDTO{},
			wantErr: false,
		},
		{
			name: "MatchByMatchID - empty matchID",
			fields: fields{
				key: devKey,
			},
			args: args{
				matchID: "",
				region:  "NA",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "MatchByMatchID - invalid region",
			fields: fields{
				key: devKey,
			},
			args: args{
				matchID: testMatchID,
				region:  "invalid",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "MatchByMatchID - match not found",
			fields: fields{
				key: devKey,
			},
			args: args{
				matchID: "invalid",
				region:  "NA",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewRiotAPI(tt.fields.key)
			got, err := api.GetMatchByMatchID(tt.args.matchID, tt.args.region)
			tt.want = got
			if (err != nil) != tt.wantErr {
				t.Errorf("RiotAPI.GetMatchByMatchID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RiotAPI.GetMatchByMatchID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRiotAPI_MatchListByPUUID(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		puuid  string
		region string
		args   MatchListArgs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "MatchListByPUUID - valid input",
			fields: fields{
				key: devKey,
			},
			args: args{
				puuid:  validPUUID,
				region: "NA",
				args: MatchListArgs{
					Start: intRef(0),
					Count: intRef(20),
				},
			},
			want:    make([]string, 20),
			wantErr: false,
		},
		{
			name: "MatchListByPUUID - invalid puuid",
			fields: fields{
				key: devKey,
			},
			args: args{
				puuid:  "invalid",
				region: "NA",
				args:   MatchListArgs{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "MatchListByPUUID - invalid region",
			fields: fields{
				key: devKey,
			},
			args: args{
				puuid:  validPUUID,
				region: "invalid",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "MatchListByPUUID - no args",
			fields: fields{
				key: devKey,
			},
			args: args{
				puuid:  validPUUID,
				region: "NA",
				args:   MatchListArgs{},
			},
			want:    make([]string, 20),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewRiotAPI(tt.fields.key)
			got, err := api.GetMatchListByPUUID(tt.args.puuid, tt.args.region, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("RiotAPI.GetMatchListByPUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("RiotAPI.GetMatchListByPUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
