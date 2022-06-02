package types

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/trustwallet/go-primitives/coin"
)

func TestGetEthereumTokenTypeByIndex(t *testing.T) {
	type args struct {
		coinIndex uint
	}
	tests := []struct {
		name    string
		args    args
		want    TokenType
		wantErr bool
	}{
		{
			name: "Ethereum ERC20",
			args: args{coinIndex: coin.ETHEREUM},
			want: ERC20,
		},
		{
			name: "Ethereum Classic ETC20",
			args: args{coinIndex: coin.CLASSIC},
			want: ETC20,
		},
		{
			name: "Poa POA20",
			args: args{coinIndex: coin.POA},
			want: POA20,
		},
		{
			name: "Callisto CLO20",
			args: args{coinIndex: coin.CALLISTO},
			want: CLO20,
		},
		{
			name: "WanChain WAN20",
			args: args{coinIndex: coin.WANCHAIN},
			want: WAN20,
		},
		{
			name: "Thundertoken TT20",
			args: args{coinIndex: coin.THUNDERTOKEN},
			want: TT20,
		},
		{
			name: "GoChain GO20",
			args: args{coinIndex: coin.GOCHAIN},
			want: GO20,
		},
		{
			name: "TomoChain TRC21",
			args: args{coinIndex: coin.TOMOCHAIN},
			want: TRC21,
		},
		{
			name: "Smartchain BEP20",
			args: args{coinIndex: coin.SMARTCHAIN},
			want: BEP20,
		},
		{
			name:    "Solana SPL",
			args:    args{coinIndex: coin.SOLANA},
			want:    "",
			wantErr: true,
		},
		{
			name: "Polygon POLYGON",
			args: args{coinIndex: coin.POLYGON},
			want: POLYGON,
		},
		{
			name: "Optimism ERC20",
			args: args{coinIndex: coin.OPTIMISM},
			want: OPTIMISM,
		},
		{
			name: "xDAI ERC20",
			args: args{coinIndex: coin.XDAI},
			want: XDAI,
		},
		{
			name: "Avalanche ERC20",
			args: args{coinIndex: coin.AVALANCHEC},
			want: AVALANCHE,
		},
		{
			name: "Fantom ERC20",
			args: args{coinIndex: coin.FANTOM},
			want: FANTOM,
		},
		{
			name: "Heco ERC20",
			args: args{coinIndex: coin.HECO},
			want: HRC20,
		},
		{
			name: "Ronin ERC20",
			args: args{coinIndex: coin.RONIN},
			want: RONIN,
		},
		{
			name: "Cronos CRC20",
			args: args{coinIndex: coin.CRONOS},
			want: CRC20,
		},
		{
			name: "KuCoin KRC20",
			args: args{coinIndex: coin.KCC},
			want: KRC20,
		},
		{
			name: "Aurora AURORA",
			args: args{coinIndex: coin.AURORA},
			want: AURORA,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEthereumTokenTypeByIndex(tt.args.coinIndex)
			if tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.want, got)

		})
	}
}

func TestGetTokenType(t *testing.T) {
	type args struct {
		coinIndex uint
		tokenID   string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		wantBool bool
	}{
		{
			name:     "Ethereum",
			args:     args{coin.ETHEREUM, ""},
			want:     string(ERC20),
			wantBool: true,
		},
		{
			name:     "Cronos",
			args:     args{coin.CRONOS, ""},
			want:     string(CRC20),
			wantBool: true,
		},
		{
			name:     "KuCoin",
			args:     args{coin.KCC, ""},
			want:     string(KRC20),
			wantBool: true,
		},
		{
			name:     "Aurora",
			args:     args{coin.AURORA, ""},
			want:     string(AURORA),
			wantBool: true,
		},
		{
			name:     "Tron TRC20",
			args:     args{coin.TRON, "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"},
			want:     string(TRC20),
			wantBool: true,
		},
		{
			name:     "Tron TRC10",
			args:     args{coin.TRON, "1"},
			want:     string(TRC10),
			wantBool: true,
		},
		{
			name:     "Binance",
			args:     args{coin.BINANCE, ""},
			want:     string(BEP2),
			wantBool: true,
		},
		{
			name:     "Waves",
			args:     args{coin.WAVES, ""},
			want:     string(WAVES),
			wantBool: true,
		},
		{
			name:     "Theta",
			args:     args{coin.THETA, ""},
			want:     string(THETA),
			wantBool: true,
		},
		{
			name:     "Ontology",
			args:     args{coin.ONTOLOGY, ""},
			want:     string(ONTOLOGY),
			wantBool: true,
		},
		{
			name:     "Nuls",
			args:     args{coin.NULS, ""},
			want:     string(NRC20),
			wantBool: true,
		},
		{
			name:     "Vechain",
			args:     args{coin.VECHAIN, ""},
			want:     string(VET),
			wantBool: true,
		},
		{
			name:     "Vechain",
			args:     args{coin.NEO, ""},
			want:     string(NEP5),
			wantBool: true,
		},
		{
			name:     "Eos",
			args:     args{coin.EOS, ""},
			want:     string(EOS),
			wantBool: true,
		},
		{
			name:     "Solana",
			args:     args{coin.SOLANA, ""},
			want:     string(SPL),
			wantBool: true,
		},
		{
			name:     "Bitcoin",
			args:     args{coin.BITCOIN, ""},
			wantBool: false,
		},
		{
			name:     "TERRA CW20",
			args:     args{coin.TERRA, "terra14z56l0fp2lsf86zy3hty2z47ezkhnthtr9yq76"},
			want:     string(CW20),
			wantBool: true,
		},
		{
			name:     "OASIS",
			args:     args{coin.OASIS, ""},
			want:     string(OASIS),
			wantBool: true,
		},
		{
			name:     "Stellar",
			args:     args{coin.STELLAR, ""},
			want:     string(STELLAR),
			wantBool: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotBool := GetTokenType(tt.args.coinIndex, tt.args.tokenID)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantBool, gotBool)

		})
	}
}

func TestGetTokenVersion(t *testing.T) {
	type args struct {
		t TokenType
	}
	tests := []struct {
		name string
		args args
		want TokenVersion
	}{
		{
			"ERC20 token version",
			args{t: ERC20},
			TokenVersionV0,
		},
		{
			"SPL token version",
			args{t: SPL},
			TokenVersionV3,
		},
		{
			"Polygon token version",
			args{t: POLYGON},
			TokenVersionV4,
		},
		{
			"Fantom token version",
			args{t: FANTOM},
			TokenVersionV5,
		},
		{
			"Terra token version",
			args{t: "TERRA"},
			TokenVersionV6,
		},
		{
			"CELO token version",
			args{t: "CELO"},
			TokenVersionV7,
		},
		{
			"CW20 token version",
			args{t: "CW20"},
			TokenVersionV8,
		},
		{
			"CRC20 token version",
			args{t: "CRC20"},
			TokenVersionV9,
		},
		{
			"ESDT token version",
			args{t: "ESDT"},
			TokenVersionV9,
		},
		{
			"Random token version",
			args{t: "Random"},
			TokenVersionUndefined,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTokenVersion(tt.args.t); got != tt.want {
				t.Errorf("GetTokenVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTokenVersionImplementEverySupportedTokenTypes(t *testing.T) {
	for _, tokenType := range GetTokenTypes() {
		tokenVersion := GetTokenVersion(tokenType)
		assert.NotEqual(t, TokenVersionUndefined, tokenVersion, tokenType)
	}
}
