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
		addr      string
		tokenType string
		chain     Coin
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
				addr:      "token",
				tokenType: "",
				chain:     Ethereum(),
			},
			want:    "https://etherscan.io/token/token",
			wantErr: false,
		},
		{
			name: "Test custom chain",
			args: args{
				addr:      "token",
				tokenType: "",
				chain:     Coin{Name: "Custom Coin"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Tron TRC10",
			args: args{
				addr:      "10001",
				tokenType: "",
				chain:     Tron(),
			},
			want:    "https://tronscan.io/#/token/10001",
			wantErr: false,
		},
		{
			name: "Test Tron TRC20",
			args: args{
				addr:      "token",
				tokenType: "",
				chain:     Tron(),
			},
			want:    "https://tronscan.io/#/token20/token",
			wantErr: false,
		},
		{
			name: "Test Elrond ESDT",
			args: args{
				addr:      "EGLDUSDC-594e5e",
				tokenType: "ESDT",
				chain:     Elrond(),
			},
			want:    "https://explorer.elrond.com/tokens/EGLDUSDC-594e5e",
			wantErr: false,
		},
		{
			name: "Test STELLAR",
			args: args{
				addr:      "yXLM-GARDNV3Q7YGT4AKSDF25LT32YSCCW4EV22Y2TV3I2PU2MMXJTEDL5T55",
				tokenType: "STELLAR",
				chain:     Stellar(),
			},
			want:    "https://stellar.expert/explorer/public/asset/yXLM-GARDNV3Q7YGT4AKSDF25LT32YSCCW4EV22Y2TV3I2PU2MMXJTEDL5T55",
			wantErr: false,
		},
		{
			name: "Test CRONOS",
			args: args{
				addr:      "0x145677FC4d9b8F19B5D56d1820c48e0443049a30",
				tokenType: "CRC20",
				chain:     Cronos(),
			},
			want:    "https://cronos.org/explorer/address/0x145677FC4d9b8F19B5D56d1820c48e0443049a30/token-transfers",
			wantErr: false,
		},
		{
			name: "Test Aurora",
			args: args{
				addr:      "0x7b37ABAe99A560Aec9497DBbe1741204bd439AC0",
				tokenType: "AURORA",
				chain:     Aurora(),
			},
			want:    "https://aurorascan.dev/address/0x7b37ABAe99A560Aec9497DBbe1741204bd439AC0",
			wantErr: false,
		},
		{
			name: "Test KuCoin",
			args: args{
				addr:      "0x2cA48b4eeA5A731c2B54e7C3944DBDB87c0CFB6F",
				tokenType: "KRC20",
				chain:     Kcc(),
			},
			want:    "https://explorer.kcc.io/token/0x2cA48b4eeA5A731c2B54e7C3944DBDB87c0CFB6F",
			wantErr: false,
		},
		{
			name: "Test Algorand",
			args: args{
				addr:      "test",
				tokenType: "ALGORAND",
				chain:     Algorand(),
			},
			want:    "https://algoexplorer.io/asset/test",
			wantErr: false,
		},
		{
			name: "Test KavaEvm",
			args: args{
				addr:      "test",
				tokenType: "KAVA",
				chain:     Kavaevm(),
			},
			want:    "https://explorer.kava.io/token/test",
			wantErr: false,
		},
		{
			name: "Test Meter",
			args: args{
				addr:      "test",
				tokenType: "METER",
				chain:     Meter(),
			},
			want:    "https://scan.meter.io/address/test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCoinExploreURL(tt.args.chain, tt.args.addr, tt.args.tokenType)
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
