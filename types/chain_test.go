package types

import (
	"reflect"
	"testing"

	"github.com/trustwallet/go-primitives/coin"
)

func TestGetChainFromAssetType(t *testing.T) {
	type args struct {
		type_ string
	}

	tests := []struct {
		name    string
		args    args
		want    coin.Coin
		wantErr bool
	}{
		{
			name: "Test ERC20",
			args: args{
				type_: "ERC20",
			},
			want:    coin.Ethereum(),
			wantErr: false,
		},
		{
			name: "Test custom chain type",
			args: args{
				type_: "UNKNOWN20",
			},
			want:    coin.Coin{},
			wantErr: true,
		},
		{
			name: "Test TRC20",
			args: args{
				type_: "TRC20",
			},
			want:    coin.Tron(),
			wantErr: false,
		},
		{
			name: "Test TRC10",
			args: args{
				type_: "TRC10",
			},
			want:    coin.Tron(),
			wantErr: false,
		},
		{
			name: "Test TRC10",
			args: args{
				type_: "TRC10",
			},
			want:    coin.Tron(),
			wantErr: false,
		},
		{
			name: "Test TOMO",
			args: args{
				type_: "TOMO",
			},
			want:    coin.Tomochain(),
			wantErr: false,
		},
		{
			name: "Test TRC21",
			args: args{
				type_: "TRC21",
			},
			want:    coin.Tomochain(),
			wantErr: false,
		},
		{
			name: "Test STELLAR",
			args: args{
				type_: "STELLAR",
			},
			want:    coin.Stellar(),
			wantErr: false,
		},
		{
			name: "Test Conflux eSpace",
			args: args{
				type_: "CFXEVM",
			},
			want:    coin.Cfxevm(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChainFromAssetType(tt.args.type_)
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

func TestGetChainFromAssetTypeFullness(t *testing.T) {
	for _, tokenType := range GetTokenTypes() {
		if tokenType == ERC721 || tokenType == ERC1155 {
			continue
		}

		_, err := GetChainFromAssetType(string(tokenType))
		if err != nil {
			t.Error(err)
		}
	}
}
