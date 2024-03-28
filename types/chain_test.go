package types

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/trustwallet/go-primitives/coin"
)

func TestGetChainsFromAssetType(t *testing.T) {
	type args struct {
		type_ string
	}

	tests := []struct {
		name    string
		args    args
		want    []coin.Coin
		wantErr bool
	}{
		{
			name: "Test ERC20",
			args: args{
				type_: "ERC20",
			},
			want:    []coin.Coin{coin.Ethereum()},
			wantErr: false,
		},
		{
			name: "Test custom chain type",
			args: args{
				type_: "UNKNOWN20",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test TRC20",
			args: args{
				type_: "TRC20",
			},
			want:    []coin.Coin{coin.Tron()},
			wantErr: false,
		},
		{
			name: "Test TRC10",
			args: args{
				type_: "TRC10",
			},
			want:    []coin.Coin{coin.Tron()},
			wantErr: false,
		},
		{
			name: "Test TRC10",
			args: args{
				type_: "TRC10",
			},
			want:    []coin.Coin{coin.Tron()},
			wantErr: false,
		},
		{
			name: "Test TOMO",
			args: args{
				type_: "TOMO",
			},
			want:    []coin.Coin{coin.Tomochain()},
			wantErr: false,
		},
		{
			name: "Test TRC21",
			args: args{
				type_: "TRC21",
			},
			want:    []coin.Coin{coin.Tomochain()},
			wantErr: false,
		},
		{
			name: "Test STELLAR",
			args: args{
				type_: "STELLAR",
			},
			want:    []coin.Coin{coin.Stellar()},
			wantErr: false,
		},
		{
			name: "Test Conflux eSpace",
			args: args{
				type_: "CONFLUX",
			},
			want:    []coin.Coin{coin.Cfxevm()},
			wantErr: false,
		},
		{
			name: "Test CW20",
			args: args{
				type_: "CW20",
			},
			want:    []coin.Coin{coin.Terra(), coin.Nativeinjective()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChainsFromAssetType(tt.args.type_)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoinForId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}

func TestGetChainsFromAssetTypeFullness(t *testing.T) {
	for _, tokenType := range GetTokenTypes() {
		if tokenType == ERC721 || tokenType == ERC1155 || tokenType == NATIVEINJECTIVE {
			continue
		}

		_, err := GetChainsFromAssetType(string(tokenType))
		if err != nil {
			t.Error(err)
		}
	}
}
