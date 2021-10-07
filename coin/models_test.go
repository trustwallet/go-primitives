package coin

import (
	"reflect"
	"testing"
)

func TestGetCoinForId(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Coin
		wantErr bool
	}{
		{
			"Test ethereum",
			args{
				id: "ethereum",
			},
			Ethereum(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCoinForId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoinForId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoinForId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCoinExploreURL(t *testing.T) {
	type args struct {
		addr  string
		chain string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test ethereum",
			args: args{
				addr:  "token",
				chain: "ethereum",
			},
			want:    "https://etherscan.io/token/token",
			wantErr: false,
		},
		{
			name: "Test custom chain",
			args: args{
				addr:  "token",
				chain: "customchain",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCoinExploreURL(tt.args.chain, tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoinForId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoinForId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
