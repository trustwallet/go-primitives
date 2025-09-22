package coin

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
				chain: Coin{
					ID:   1, // some id
					Name: "Custom Coin",
				},
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
			want:    "https://explorer.multiversx.com/tokens/EGLDUSDC-594e5e",
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
		{
			name: "Test Evmos",
			args: args{
				addr:      "test",
				tokenType: "EVMOS_ERC20",
				chain:     Evmos(),
			},
			want:    "https://evm.evmos.org/address/test",
			wantErr: false,
		},
		{
			name: "Test Okc",
			args: args{
				addr:      "test",
				tokenType: "KIP20",
				chain:     Okc(),
			},
			want:    "https://www.oklink.com/en/okc/address/test",
			wantErr: false,
		},
		{
			name: "Test Moonbeam",
			args: args{
				addr:      "test",
				tokenType: "MOONBEAM",
				chain:     Moonbeam(),
			},
			want:    "https://moonscan.io/token/test",
			wantErr: false,
		},
		{
			name: "Test Klaytn",
			args: args{
				addr:      "test",
				tokenType: "KLAYTN",
				chain:     Klaytn(),
			},
			want:    "https://kaiascan.io/token/test",
			wantErr: false,
		},
		{
			name: "Test Metis",
			args: args{
				addr:      "test",
				tokenType: "METIS",
				chain:     Metis(),
			},
			want:    "https://andromeda-explorer.metis.io/token/test",
			wantErr: false,
		},
		{
			name: "Test Moonriver",
			args: args{
				addr:      "test",
				tokenType: "MOONRIVER",
				chain:     Moonriver(),
			},
			want:    "https://moonriver.moonscan.io/token/test",
			wantErr: false,
		},
		{
			name: "Test Boba",
			args: args{
				addr:      "test",
				tokenType: "BOBA",
				chain:     Boba(),
			},
			want:    "https://bobascan.com/token/test",
			wantErr: false,
		},
		{
			name: "Test Ton",
			args: args{
				addr:      "test",
				tokenType: "TON",
				chain:     Ton(),
			},
			want:    "https://tonscan.org/address/test",
			wantErr: false,
		},
		{
			name: "Test ZKEVM",
			args: args{
				addr:      "test",
				tokenType: "ZKEVM",
				chain:     Polygonzkevm(),
			},
			want:    "https://explorer.public.zkevm-test.net/address/test",
			wantErr: false,
		},
		{
			name: "Test ZKSync",
			args: args{
				addr:      "test",
				tokenType: "ZKSYNC",
				chain:     Zksync(),
			},
			want:    "https://explorer.zksync.io/address/test",
			wantErr: false,
		},
		{
			name: "Test Sui",
			args: args{
				addr:      "test",
				tokenType: "Sui",
				chain:     Sui(),
			},
			want:    "https://explorer.sui.io/address/test",
			wantErr: false,
		},
		{
			name: "Test Stride",
			args: args{
				addr:      "test",
				tokenType: "Stride",
				chain:     Stride(),
			},
			want:    "https://www.mintscan.io/stride/account/test",
			wantErr: false,
		},
		{
			name: "Test Neutron",
			args: args{
				addr:      "test",
				tokenType: "Neutron",
				chain:     Neutron(),
			},
			want:    "https://www.mintscan.io/neutron/account/test",
			wantErr: false,
		},
		{
			name: "Test IoTex EVM",
			args: args{
				addr:      "test",
				tokenType: "",
				chain:     Iotexevm(),
			},
			want:    "https://iotexscan.io/address/test#transactions",
			wantErr: false,
		},
		{
			name: "Test CFXEVM",
			args: args{
				addr:      "test",
				tokenType: "",
				chain:     Cfxevm(),
			},
			want:    "https://evm.confluxscan.net/address/test",
			wantErr: false,
		},
		{
			name: "Test Acala system token",
			args: args{
				addr:      "test",
				tokenType: "",
				chain:     Acala(),
			},
			want:    "https://acala.subscan.io/system_token_detail?unique_id=test",
			wantErr: false,
		},
		{
			name: "Test Acala custom token",
			args: args{
				addr:      "test",
				tokenType: "custom_token",
				chain:     Acala(),
			},
			want:    "https://acala.subscan.io/custom_token?customTokenId=test",
			wantErr: false,
		},
		{
			name: "Test BASE20 token",
			args: args{
				addr:      "0x48bcf9455ba97cc439a2efbcfdf8f1afe692139b",
				tokenType: "BASE20",
				chain:     Base(),
			},
			want:    "https://basescan.org/token/0x48bcf9455ba97cc439a2efbcfdf8f1afe692139b",
			wantErr: false,
		},
		{
			name: "Test NEON token",
			args: args{
				addr:      "0x5f38248f339bf4e84a2caf4e4c0552862dc9f82a",
				tokenType: "NEON",
				chain:     Neon(),
			},
			want:    "https://neonscan.org/token/0x5f38248f339bf4e84a2caf4e4c0552862dc9f82a",
			wantErr: false,
		},
		{
			name: "Test Native ZETA",
			args: args{
				addr:      "zeta14py36sx57ud82t9yrks9z6hdsrpn5x6kmxs0ne",
				tokenType: "ZETA",
				chain:     Zetachain(),
			},
			want:    "https://explorer.zetachain.com/address/zeta14py36sx57ud82t9yrks9z6hdsrpn5x6kmxs0ne",
			wantErr: false,
		},
		{
			name: "Test ZETA EVM",
			args: args{
				addr:      "0x890a1b6dc3ca666eacda1c453115494291c6bc6a",
				tokenType: "ZETA",
				chain:     Zetaevm(),
			},
			want:    "https://explorer.zetachain.com/address/0x890a1b6dc3ca666eacda1c453115494291c6bc6a",
			wantErr: false,
		},
		{
			name: "Test Celo",
			args: args{
				addr:      "0x639A647fbe20b6c8ac19E48E2de44ea792c62c5C",
				tokenType: "CELO",
				chain:     Celo(),
			},
			want:    "https://explorer.celo.org/mainnet/address/0x639A647fbe20b6c8ac19E48E2de44ea792c62c5C",
			wantErr: false,
		},
		{
			name: "Test Blast",
			args: args{
				addr:      "0x4300000000000000000000000000000000000004",
				tokenType: "BLAST",
				chain:     Blast(),
			},
			want:    "https://blastscan.io/token/0x4300000000000000000000000000000000000004",
			wantErr: false,
		},
		{
			name: "Test Scroll",
			args: args{
				addr:      "0xf55bec9cafdbe8730f096aa55dad6d22d44099df",
				tokenType: "SCROLL",
				chain:     Scroll(),
			},
			want:    "https://scrollscan.com/token/0xf55bec9cafdbe8730f096aa55dad6d22d44099df",
			wantErr: false,
		},
		{
			name: "Test zkLink Nova - Dai",
			args: args{
				addr:      "0xF573fA04A73d5AC442F3DEa8741317fEaA3cDeab",
				tokenType: "ZKLINKNOVA",
				chain:     Zklinknova(),
			},
			want:    "https://explorer.zklink.io/address/0xF573fA04A73d5AC442F3DEa8741317fEaA3cDeab",
			wantErr: false,
		},
		{
			name: "Test Aptos (legacy)",
			args: args{
				addr:      "0xacd014e8bdf395fa8497b6d585b164547a9d45269377bdf67c96c541b7fec9ed::coin::T",
				tokenType: "APTOS",
				chain:     Aptos(),
			},
			want:    "https://explorer.aptoslabs.com/coin/0xacd014e8bdf395fa8497b6d585b164547a9d45269377bdf67c96c541b7fec9ed::coin::T?network=mainnet",
			wantErr: false,
		},
		{
			name: "Test Aptos (fungible asset)",
			args: args{
				addr:      "0x357b0b74bc833e95a115ad22604854d6b0fca151cecd94111770e5d6ffc9dc2b",
				tokenType: "APTOSFA",
				chain:     Aptos(),
			},
			want:    "https://explorer.aptoslabs.com/fungible_asset/0x357b0b74bc833e95a115ad22604854d6b0fca151cecd94111770e5d6ffc9dc2b?network=mainnet",
			wantErr: false,
		},
		{
			name: "Test XRP",
			args: args{
				addr:      "rMxCKbEDwqr76QuheSUMdEGf4B9xJ8m5De",
				tokenType: "XRP",
				chain:     Ripple(),
			},
			want:    "https://xrpscan.com/account/rMxCKbEDwqr76QuheSUMdEGf4B9xJ8m5De",
			wantErr: false,
		},
		{
			name: "Test Sonic",
			args: args{
				addr:      "0x29219dd400f2bf60e5a23d13be72b486d4038894",
				tokenType: "SONIC",
				chain:     Sonic(),
			},
			want:    "https://sonicscan.org/token/0x29219dd400f2bf60e5a23d13be72b486d4038894",
			wantErr: false,
		},
		{
			name: "Test Plasma",
			args: args{
				addr:      "0xB8CE59FC3717ada4C02eaDF9682A9e934F625ebb",
				tokenType: "PLASMA",
				chain:     Plasma(),
			},
			want:    "https://plasmascan.to/token/0xB8CE59FC3717ada4C02eaDF9682A9e934F625ebb",
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

var evmCoinsTestSet = map[uint]struct{}{
	ETHEREUM:     {},
	CLASSIC:      {},
	POA:          {},
	CALLISTO:     {},
	WANCHAIN:     {},
	THUNDERTOKEN: {},
	GOCHAIN:      {},
	TOMOCHAIN:    {},
	SMARTCHAIN:   {},
	POLYGON:      {},
	OPTIMISM:     {},
	XDAI:         {},
	AVALANCHEC:   {},
	FANTOM:       {},
	HECO:         {},
	RONIN:        {},
	CRONOS:       {},
	KCC:          {},
	AURORA:       {},
	ARBITRUM:     {},
	KAVAEVM:      {},
	METER:        {},
	EVMOS:        {},
	CELO:         {},
	OKC:          {},
	MOONBEAM:     {},
	KLAYTN:       {},
	METIS:        {},
	MOONRIVER:    {},
	BOBA:         {},
	POLYGONZKEVM: {},
	ZKSYNC:       {},
	CFXEVM:       {},
	ACALAEVM:     {},
	BASE:         {},
	NEON:         {},
	IOTEXEVM:     {},
	OPBNB:        {},
	LINEA:        {},
	MANTLE:       {},
	MANTA:        {},
	ZETAEVM:      {},
	MERLIN:       {},
	BLAST:        {},
	SCROLL:       {},
	BOUNCEBIT:    {},
	ZKLINKNOVA:   {},
	SONIC:        {},
	PLASMA:       {},
}

// TestEvmCoinsList This test will automatically fail when new EVM chain is added to coins.yml
// To fix it, extend evmCoinsTestSet with that new chain
func TestEvmCoinsList(t *testing.T) {
	for _, c := range Coins {
		_, ok := evmCoinsTestSet[c.ID]
		assert.Equalf(t, c.Blockchain == BlockchainEthereum, ok, fmt.Sprintf("chain: %s", c.Handle))
	}
}

func TestIsEVM(t *testing.T) {
	for _, c := range Coins {
		_, ok := evmCoinsTestSet[c.ID]
		if ok {
			assert.Truef(t, IsEVM(c.ID), fmt.Sprintf("chain: %s", c.Handle))
		}
	}
}
