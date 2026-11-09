package coin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinStringAndAssetID(t *testing.T) {
	tests := []struct {
		name       string
		coin       Coin
		token      string
		wantString string
		wantAsset  AssetID
		wantToken  AssetID
	}{
		{
			name:       "ethereum token",
			coin:       Ethereum(),
			token:      "0xabc",
			wantString: "[ETH] Ethereum (#60)",
			wantAsset:  "c60",
			wantToken:  "c60_t0xabc",
		},
		{
			name:       "empty token keeps the coin id",
			coin:       Ethereum(),
			token:      "",
			wantString: "[ETH] Ethereum (#60)",
			wantAsset:  "c60",
			wantToken:  "c60",
		},
		{
			name:       "blank token is still appended",
			coin:       Ethereum(),
			token:      " ",
			wantString: "[ETH] Ethereum (#60)",
			wantAsset:  "c60",
			wantToken:  "c60_t ",
		},
		{
			name:       "zero value",
			coin:       Coin{},
			token:      "",
			wantString: "[]  (#0)",
			wantAsset:  "c0",
			wantToken:  "c0",
		},
		{
			name:       "bitcoin id is zero and chain id is nil",
			coin:       Bitcoin(),
			token:      "ordi",
			wantString: "[BTC] Bitcoin (#0)",
			wantAsset:  "c0",
			wantToken:  "c0_tordi",
		},
		{
			name:       "large id",
			coin:       Arc(),
			token:      "usdc",
			wantString: "[USDC] Arc (#10005042)",
			wantAsset:  "c10005042",
			wantToken:  "c10005042_tusdc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantString, tt.coin.String())
			assert.Equal(t, tt.wantAsset, tt.coin.AssetID())
			assert.Equal(t, tt.wantToken, tt.coin.TokenAssetID(tt.token))
		})
	}
}

func TestCoinStringNilReceiver(t *testing.T) {
	var c *Coin
	assert.Panics(t, func() { _ = c.String() })
}

func TestChainIDOptional(t *testing.T) {
	assert.Nil(t, Bitcoin().ChainID)

	if assert.NotNil(t, Ethereum().ChainID) {
		assert.Equal(t, uint(1), *Ethereum().ChainID)
	}
}

func TestGetCoinForIdTable(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		want    Coin
		wantErr string
	}{
		{name: "ethereum", id: "ethereum", want: Ethereum()},
		{name: "empty", id: "", wantErr: "unknown id "},
		{name: "wrong case", id: "Ethereum", wantErr: "unknown id Ethereum"},
		{name: "unknown", id: "not-a-coin", wantErr: "unknown id not-a-coin"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCoinForId(tt.id)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				assert.Equal(t, Coin{}, got)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetCoinForIdRoundTrip(t *testing.T) {
	seen := map[string]uint{}
	for id, c := range Coins {
		if prev, ok := seen[c.Handle]; ok {
			t.Fatalf("handle %s is used by coin %d and coin %d", c.Handle, prev, id)
		}
		seen[c.Handle] = id

		got, err := GetCoinForId(c.Handle)
		assert.NoError(t, err)
		assert.Equal(t, c, got)
		assert.Equal(t, c, Chains[c.Handle])
	}
	assert.Equal(t, len(Coins), len(Chains))
}

func TestIsEVMTable(t *testing.T) {
	tests := []struct {
		name string
		id   uint
		want bool
	}{
		{name: "ethereum", id: ETHEREUM, want: true},
		{name: "bitcoin id zero", id: BITCOIN, want: false},
		{name: "greenfield", id: GBNB, want: false},
		{name: "missing id", id: ^uint(0), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsEVM(tt.id))
		})
	}
}

func TestGetCoinExploreURLTable(t *testing.T) {
	const token = "tok"

	tests := []struct {
		name      string
		coin      Coin
		token     string
		tokenType string
		want      string
		wantErr   string
	}{
		{name: "empty ethereum token", coin: Ethereum(), want: "https://etherscan.io/token/"},
		{name: "binance", coin: Binance(), token: token, want: "https://explorer.binance.org/asset/tok"},
		{name: "smartchain", coin: Smartchain(), token: token, want: "https://bscscan.com/token/tok"},
		{name: "eos", coin: Eos(), token: token, want: "https://bloks.io/account/tok"},
		{name: "neo prefixes 0x", coin: Neo(), token: token, want: "https://neo.tokenview.com/en/token/0xtok"},
		{name: "nuls", coin: Nuls(), token: token, want: "https://nulscan.io/token/info?contractAddress=tok"},
		{name: "wanchain", coin: Wanchain(), token: token, want: "https://www.wanscan.org/token/tok"},
		{name: "solana", coin: Solana(), token: token, want: "https://solscan.io/token/tok"},
		{name: "tomochain", coin: Tomochain(), token: token, want: "https://tomoscan.io/token/tok"},
		{name: "kava ignores the token", coin: Kava(), token: "", want: "https://www.mintscan.io/kava"},
		{name: "ontology ignores the token", coin: Ontology(), token: token, want: "https://explorer.ont.io"},
		{name: "gochain", coin: Gochain(), token: token, want: "https://explorer.gochain.io/addr/tok"},
		{name: "theta ignores the token", coin: Theta(), want: "https://explorer.thetatoken.org/"},
		{name: "thundercore", coin: Thundertoken(), token: token, want: "https://viewblock.io/thundercore/address/tok"},
		{name: "classic", coin: Classic(), token: token, want: "https://blockscout.com/etc/mainnet/tokens/tok"},
		{name: "vechain", coin: Vechain(), token: token, want: "https://explore.vechain.org/accounts/tok"},
		{name: "waves", coin: Waves(), token: token, want: "https://wavesexplorer.com/assets/tok"},
		{name: "xdai", coin: Xdai(), token: token, want: "https://blockscout.com/xdai/mainnet/tokens/tok"},
		{name: "poa", coin: Poa(), token: token, want: "https://blockscout.com/poa/core/tokens/tok"},
		{name: "polygon", coin: Polygon(), token: token, want: "https://polygonscan.com/token/tok"},
		{name: "optimism", coin: Optimism(), token: token, want: "https://optimistic.etherscan.io/token/tok"},
		{name: "avalanche", coin: Avalanchec(), token: token, want: "https://snowtrace.io/token/tok"},
		{name: "arbitrum", coin: Arbitrum(), token: token, want: "https://arbiscan.io/token/tok"},
		{name: "fantom", coin: Fantom(), token: token, want: "https://ftmscan.com/token/tok"},
		{name: "terra", coin: Terra(), token: token, want: "https://finder.terra.money/mainnet/address/tok"},
		{name: "ronin", coin: Ronin(), token: token, want: "https://explorer.roninchain.com/token/tok"},
		{name: "elrond collection", coin: Elrond(), token: token, tokenType: "NFT", want: "https://explorer.multiversx.com/collections/tok"},
		{name: "elrond type is case sensitive", coin: Elrond(), token: token, tokenType: "esdt", want: "https://explorer.multiversx.com/collections/tok"},
		{name: "heco", coin: Heco(), token: token, want: "https://hecoinfo.com/token/tok"},
		{name: "oasis", coin: Oasis(), token: token, want: "https://explorer.oasis.updev.si/token/tok"},
		{name: "cryptoorg", coin: Cryptoorg(), token: token, want: "https://crypto.org/explorer/account/tok"},
		{name: "tezos", coin: Tezos(), token: token, want: "https://tzstats.com/tok"},
		{name: "acala evm", coin: Acalaevm(), token: token, want: "https://blockscout.acala.network/token/tok"},
		{name: "cardano", coin: Cardano(), token: token, want: "https://cexplorer.io/asset/tok"},
		{name: "mantle", coin: Mantle(), token: token, want: "https://explorer.mantle.xyz/address/tok"},
		{name: "linea", coin: Linea(), token: token, want: "https://explorer.linea.build/token/tok"},
		{name: "opbnb", coin: Opbnb(), token: token, want: "https://opbnbscan.com/token/tok"},
		{name: "manta", coin: Manta(), token: token, want: "https://pacific-explorer.manta.network/token/tok"},
		{name: "bitcoin brc20", coin: Bitcoin(), token: token, want: "https://unisat.io/brc20/tok"},
		{name: "celestia ignores the token", coin: Tia(), want: "https://www.mintscan.io/celestia"},
		{name: "dydx ignores the token", coin: Dydx(), token: token, want: "https://www.mintscan.io/dydx"},
		{name: "monad", coin: Monad(), token: token, want: "https://explorer.monad.xyz/token/tok"},
		{name: "megaeth", coin: Megaeth(), token: token, want: "https://mega.etherscan.com/token/tok"},
		{name: "sei evm", coin: Seievm(), token: token, want: "https://seitrace.com/token/tok"},
		{name: "hyperevm", coin: Hyperevm(), token: token, want: "https://hyperevmscan.io/token/tok"},
		{name: "aptos empty type uses the coin url", coin: Aptos(), token: token, want: "https://explorer.aptoslabs.com/coin/tok?network=mainnet"},
		{
			name:  "tron max uint64 is trc10",
			coin:  Tron(),
			token: "18446744073709551615",
			want:  "https://tronscan.io/#/token/18446744073709551615",
		},
		{
			name:  "tron past uint64 is trc20",
			coin:  Tron(),
			token: "18446744073709551616",
			want:  "https://tronscan.io/#/token20/18446744073709551616",
		},
		{name: "tron zero is trc10", coin: Tron(), token: "0", want: "https://tronscan.io/#/token/0"},
		{name: "tron empty is trc20", coin: Tron(), want: "https://tronscan.io/#/token20/"},
		{name: "tron negative is trc20", coin: Tron(), token: "-1", want: "https://tronscan.io/#/token20/-1"},
		{name: "greenfield has no token explorer", coin: Gbnb(), token: token, wantErr: "no explorer for coin: gbnb"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCoinExploreURL(tt.coin, tt.token, tt.tokenType)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				assert.Empty(t, got)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetAddressExploreURLTable(t *testing.T) {
	const addr = "addr"

	tests := []struct {
		name    string
		coin    Coin
		address string
		want    string
		wantErr string
	}{
		{name: "empty ethereum address", coin: Ethereum(), want: "https://etherscan.io/address/"},
		{name: "binance", coin: Binance(), address: addr, want: "https://explorer.binance.org/address/addr"},
		{name: "eos", coin: Eos(), address: addr, want: "https://bloks.io/account/addr"},
		{name: "nuls", coin: Nuls(), address: addr, want: "https://nulscan.io/token/info?address=addr"},
		{name: "wanchain", coin: Wanchain(), address: addr, want: "https://www.wanscan.org/address/addr"},
		{name: "solana", coin: Solana(), address: addr, want: "https://solscan.io/account/addr"},
		{name: "tomochain", coin: Tomochain(), address: addr, want: "https://tomoscan.io/address/addr"},
		{name: "kava ignores the address", coin: Kava(), want: "https://www.mintscan.io/kava"},
		{name: "ontology ignores the address", coin: Ontology(), address: addr, want: "https://explorer.ont.io"},
		{name: "gochain", coin: Gochain(), address: addr, want: "https://explorer.gochain.io/addr/addr"},
		{name: "theta ignores the address", coin: Theta(), want: "https://explorer.thetatoken.org/"},
		{name: "thundercore", coin: Thundertoken(), address: addr, want: "https://explorer-mainnet.thundercore.com/address/addr"},
		{name: "classic", coin: Classic(), address: addr, want: "https://blockscout.com/etc/mainnet/address/addr"},
		{name: "vechain", coin: Vechain(), address: addr, want: "https://explore.vechain.org/accounts/addr"},
		{name: "waves", coin: Waves(), address: addr, want: "https://wavesexplorer.com/addresses/addr"},
		{name: "xdai", coin: Xdai(), address: addr, want: "https://blockscout.com/xdai/mainnet/address/addr"},
		{name: "polygon", coin: Polygon(), address: addr, want: "https://polygonscan.com/address/addr"},
		{name: "optimism", coin: Optimism(), address: addr, want: "https://optimistic.etherscan.io/address/addr"},
		{name: "avalanche", coin: Avalanchec(), address: addr, want: "https://snowtrace.io/address/addr"},
		{name: "arbitrum", coin: Arbitrum(), address: addr, want: "https://arbiscan.io/address/addr"},
		{name: "fantom", coin: Fantom(), address: addr, want: "https://explorer.fantom.network/address/addr"},
		{name: "terra", coin: Terra(), address: addr, want: "https://finder.terra.money/mainnet/address/addr"},
		{name: "ronin", coin: Ronin(), address: addr, want: "https://explorer.roninchain.com/address/addr"},
		{name: "celo", coin: Celo(), address: addr, want: "https://explorer.celo.org/mainnet/address/addr"},
		{name: "elrond", coin: Elrond(), address: addr, want: "https://explorer.multiversx.com/accounts/addr"},
		{name: "cronos", coin: Cronos(), address: addr, want: "https://explorer.cronos.org/address/addr"},
		{name: "stellar", coin: Stellar(), address: addr, want: "https://stellar.expert/explorer/public/account/addr"},
		{name: "kcc", coin: Kcc(), address: addr, want: "https://explorer.kcc.io/address/addr"},
		{name: "aurora", coin: Aurora(), address: addr, want: "https://aurorascan.dev/address/addr"},
		{name: "kava evm", coin: Kavaevm(), address: addr, want: "https://explorer.kava.io/address/addr"},
		{name: "meter", coin: Meter(), address: addr, want: "https://scan.meter.io/address/addr"},
		{name: "aptos", coin: Aptos(), address: addr, want: "https://explorer.aptoslabs.com/account/addr?network=mainnet"},
		{name: "moonbeam", coin: Moonbeam(), address: addr, want: "https://moonscan.io/address/addr"},
		{name: "klaytn", coin: Klaytn(), address: addr, want: "https://kaiascan.io/address/addr"},
		{name: "metis", coin: Metis(), address: addr, want: "https://andromeda-explorer.metis.io/address/addr"},
		{name: "moonriver", coin: Moonriver(), address: addr, want: "https://moonriver.moonscan.io/address/addr"},
		{name: "boba", coin: Boba(), address: addr, want: "https://bobascan.com/address/addr"},
		{name: "ton", coin: Ton(), address: addr, want: "https://tonscan.org/address/addr"},
		{name: "polygon zkevm", coin: Polygonzkevm(), address: addr, want: "https://explorer.public.zkevm-test.net/address/addr"},
		{name: "zksync", coin: Zksync(), address: addr, want: "https://explorer.zksync.io/address/addr"},
		{name: "sui", coin: Sui(), address: addr, want: "https://explorer.sui.io/address/addr"},
		{name: "stride", coin: Stride(), address: addr, want: "https://www.mintscan.io/stride/account/addr"},
		{name: "neutron", coin: Neutron(), address: addr, want: "https://www.mintscan.io/neutron/account/addr"},
		{name: "iotex evm", coin: Iotexevm(), address: addr, want: "https://iotexscan.io/address/addr#transactions"},
		{name: "cryptoorg", coin: Cryptoorg(), address: addr, want: "https://crypto.org/explorer/account/addr"},
		{name: "tezos", coin: Tezos(), address: addr, want: "https://tzstats.com/addr"},
		{name: "conflux evm", coin: Cfxevm(), address: addr, want: "https://evm.confluxscan.net/address/addr"},
		{name: "acala", coin: Acala(), address: addr, want: "https://acala.subscan.io/account/addr"},
		{name: "acala evm", coin: Acalaevm(), address: addr, want: "https://blockscout.acala.network/address/addr"},
		{name: "cardano", coin: Cardano(), address: addr, want: "https://cexplorer.io/address/addr"},
		{name: "neon", coin: Neon(), address: addr, want: "https://neonscan.org/address/addr"},
		{name: "mantle", coin: Mantle(), address: addr, want: "https://explorer.mantle.xyz/address/addr"},
		{name: "linea", coin: Linea(), address: addr, want: "https://explorer.linea.build/address/addr"},
		{name: "opbnb", coin: Opbnb(), address: addr, want: "https://opbnbscan.com/address/addr"},
		{name: "manta", coin: Manta(), address: addr, want: "https://pacific-explorer.manta.network/address/addr"},
		{name: "zetachain", coin: Zetachain(), address: addr, want: "https://explorer.zetachain.com/address/addr"},
		{name: "zeta evm", coin: Zetaevm(), address: addr, want: "https://explorer.zetachain.com/address/addr"},
		{name: "blast", coin: Blast(), address: addr, want: "https://blastscan.io/address/addr"},
		{name: "scroll", coin: Scroll(), address: addr, want: "https://scrollscan.com/address/addr"},
		{name: "zklink nova", coin: Zklinknova(), address: addr, want: "https://explorer.zklink.io/address/addr"},
		{name: "ripple", coin: Ripple(), address: addr, want: "https://xrpscan.com/account/addr"},
		{name: "sonic", coin: Sonic(), address: addr, want: "https://sonicscan.org/address/addr"},
		{name: "celestia ignores the address", coin: Tia(), want: "https://www.mintscan.io/celestia"},
		{name: "dydx ignores the address", coin: Dydx(), address: addr, want: "https://www.mintscan.io/dydx"},
		{name: "plasma", coin: Plasma(), address: addr, want: "https://plasmascan.to/address/addr"},
		{name: "sei evm", coin: Seievm(), address: addr, want: "https://seitrace.com/address/addr"},
		{name: "robinhood", coin: Robinhoodchain(), address: addr, want: "https://robinhoodchain.blockscout.com/address/addr"},
		{name: "bitcoin has no address explorer", coin: Bitcoin(), address: addr, wantErr: "no explorer for coin: bitcoin"},
		{name: "greenfield has no address explorer", coin: Gbnb(), address: addr, wantErr: "no explorer for coin: gbnb"},
		{name: "zero coin has no address explorer", coin: Coin{}, wantErr: "no explorer for coin: "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAddressExploreURL(tt.coin, tt.address)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				assert.Empty(t, got)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
