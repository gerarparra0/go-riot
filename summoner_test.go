package goriot

import (
	"reflect"
	"testing"
)

const (
	devKey         = "RGAPI-5ec9e2dc-47ff-4baf-a0a3-bbb201e49e00"
	testSummonerID = "7WOTQ0plvmbJ2Kx9TwZdLKUmTJgnpm-FTNDXUsxtFCEtpwo"
	testAccountID  = "O_C6_yfCaAgj9GgV9biST8LGEk9O0QlPdxCQlsg6oVhlUp8"
	testPUUID      = "xVojEdk4AjaHfyHn5NLULbjinNRd9sLff3Jzd2is1VRF5AE6_z9RjirxUsUhvoQDhAnVgNvkSX6BMA"
	testName       = "Insulin Abuser"
	testRegion     = "NA"
)

func TestRiotAPI_GetSummonerByName(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		name   string
		region string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SummonerDTO
		wantErr bool
	}{
		{
			name: "GetSummonerByName - Insulin Abuser",
			fields: fields{
				key: devKey,
			},
			args: args{
				name:   testName,
				region: testRegion,
			},
			want:    &SummonerDTO{},
			wantErr: false,
		},
		{
			name: "GetSummonerByName - empty name",
			fields: fields{
				key: devKey,
			},
			args: args{
				name:   "",
				region: testRegion,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetSummonerByName - invalid region",
			fields: fields{
				key: devKey,
			},
			args: args{
				name:   testName,
				region: "invalid",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetSummonerByName - summoner not found",
			fields: fields{
				key: devKey,
			},
			args: args{
				name:   "asdfasdfasd kekekrrrrr",
				region: testRegion,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewRiotAPI(tt.fields.key)
			got, err := api.GetSummonerByName(tt.args.name, tt.args.region)
			tt.want = got
			if (err != nil) != tt.wantErr {
				t.Errorf("RiotAPI.GetSummonerByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RiotAPI.GetSummonerByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRiotAPI_GetSummonerByAccountID(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		accountID string
		region    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SummonerDTO
		wantErr bool
	}{
		{
			name: "GetSummonerByAccountID - Insulin Abuser",
			fields: fields{
				key: devKey,
			},
			args: args{
				accountID: testAccountID,
				region:    testRegion,
			},
			want:    &SummonerDTO{},
			wantErr: false,
		},
		{
			name: "GetSummonerByAccountID - empty accountID",
			fields: fields{
				key: devKey,
			},
			args: args{
				accountID: "",
				region:    testRegion,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetSummonerByAccountID - invalid region",
			fields: fields{
				key: devKey,
			},
			args: args{
				accountID: testAccountID,
				region:    "invalid",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetSummonerByAccountID - summoner not found",
			fields: fields{
				key: devKey,
			},
			args: args{
				accountID: "asdfasdfasd kekekrrrrr",
				region:    testRegion,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewRiotAPI(tt.fields.key)
			got, err := api.GetSummonerByAccountID(tt.args.accountID, tt.args.region)
			tt.want = got
			if (err != nil) != tt.wantErr {
				t.Errorf("RiotAPI.GetSummonerByAccountID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RiotAPI.GetSummonerByAccountID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRiotAPI_GetSummonerByPUUID(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		summonerID string
		region     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SummonerDTO
		wantErr bool
	}{
		{
			name: "GetSummonerByPUUID - Insulin Abuser",
			fields: fields{
				key: devKey,
			},
			args: args{
				summonerID: testPUUID,
				region:     testRegion,
			},
			want:    &SummonerDTO{},
			wantErr: false,
		},
		{
			name: "GetSummonerByPUUID - empty PUUID",
			fields: fields{
				key: devKey,
			},
			args: args{
				summonerID: "",
				region:     testRegion,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetSummonerByPUUID - invalid region",
			fields: fields{
				key: devKey,
			},
			args: args{
				summonerID: testPUUID,
				region:     "invalid",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetSummonerByPUUID - summoner not found",
			fields: fields{
				key: devKey,
			},
			args: args{
				summonerID: "asdfasdfasd kekekrrrrr",
				region:     testRegion,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewRiotAPI(tt.fields.key)
			got, err := api.GetSummonerByPUUID(tt.args.summonerID, tt.args.region)
			tt.want = got
			if (err != nil) != tt.wantErr {
				t.Errorf("RiotAPI.GetSummonerByPUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RiotAPI.GetSummonerByPUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
