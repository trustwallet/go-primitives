package types

import (
	"testing"

	"github.com/trustwallet/go-primitives/coin"
)

func TestGetEthereumTokenTypeByIndex(t *testing.T) {
	type args struct {
		coinIndex uint
	}
	tests := []struct {
		name string
		args args
		want TokenType
	}{
		{
			"Ethereum ERC20",
			args{coinIndex: coin.ETHEREUM},
			ERC20,
		},
		{
			"Ethereum Classic ETC20",
			args{coinIndex: coin.CLASSIC},
			ETC20,
		},
		{
			"Poa POA20",
			args{coinIndex: coin.POA},
			POA20,
		},
		{
			"Callisto CLO20",
			args{coinIndex: coin.CALLISTO},
			CLO20,
		},
		{
			"WanChain WAN20",
			args{coinIndex: coin.WANCHAIN},
			WAN20,
		},
		{
			"Thundertoken TT20",
			args{coinIndex: coin.THUNDERTOKEN},
			TT20,
		},
		{
			"GoChain GO20",
			args{coinIndex: coin.GOCHAIN},
			GO20,
		},
		{
			"TomoChain TRC21",
			args{coinIndex: coin.TOMOCHAIN},
			TRC21,
		},
		{
			"Smartchain BEP20",
			args{coinIndex: coin.SMARTCHAIN},
			BEP20,
		},
		{
			"Solana SPL",
			args{coinIndex: coin.SOLANA},
			SPL,
		},
		{
			"Polygon POLYGON",
			args{coinIndex: coin.POLYGON},
			POLYGON,
		},
		{
			"Optimism ERC20",
			args{coinIndex: coin.OPTIMISM},
			OPTIMISM,
		},
		{
			"xDAI ERC20",
			args{coinIndex: coin.XDAI},
			XDAI,
		},
		{
			"Avalanche ERC20",
			args{coinIndex: coin.AVALANCHEC},
			AVALANCHE,
		},
		{
			"Fantom ERC20",
			args{coinIndex: coin.FANTOM},
			FANTOM,
		},
		{
			"Heco ERC20",
			args{coinIndex: coin.HECO},
			HRC20,
		},
		{
			"Ronin ERC20",
			args{coinIndex: coin.RONIN},
			RONIN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEthereumTokenTypeByIndex(tt.args.coinIndex); got != tt.want {
				t.Errorf("GetEthereumTokenTypeByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
