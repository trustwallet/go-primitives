package types

import (
	"strconv"

	"github.com/trustwallet/go-primitives/asset"
	"github.com/trustwallet/go-primitives/coin"
)

type (
	TokenType string

	// Token describes the non-native tokens.
	// Examples: ERC-20, TRC-20, BEP-2
	Token struct {
		Name     string    `json:"name"`
		Symbol   string    `json:"symbol"`
		Decimals uint      `json:"decimals"`
		TokenID  string    `json:"token_id"`
		Coin     uint      `json:"coin"`
		Type     TokenType `json:"type"`
	}
)

const (
	Coin      TokenType = "coin"
	Gas       TokenType = "gas"
	ERC20     TokenType = "ERC20"
	ERC721    TokenType = "ERC721"
	ERC1155   TokenType = "ERC1155"
	BEP2      TokenType = "BEP2"
	BEP8      TokenType = "BEP8"
	BEP20     TokenType = "BEP20"
	TRC10     TokenType = "TRC10"
	ETC20     TokenType = "ETC20"
	POA20     TokenType = "POA20"
	TRC20     TokenType = "TRC20"
	TRC21     TokenType = "TRC21"
	CLO20     TokenType = "CLO20"
	GO20      TokenType = "G020"
	WAN20     TokenType = "WAN20"
	TT20      TokenType = "TT20"
	KAVA      TokenType = "KAVA"
	SPL       TokenType = "SPL"
	POLYGON   TokenType = "POLYGON"
	OPTIMISM  TokenType = "OPTIMISM"
	XDAI      TokenType = "XDAI"
	AVALANCHE TokenType = "AVALANCHE"
	FANTOM    TokenType = "FANTOM"
	HRC20     TokenType = "HRC20"
	ARBITRUM  TokenType = "ARBITRUM"
	TERRA     TokenType = "TERRA"
)

// todo test
func GetTokenType(c uint, tokenID string) (string, bool) {
	switch c {
	case coin.ETHEREUM,
		coin.CLASSIC,
		coin.POA,
		coin.CALLISTO,
		coin.WANCHAIN,
		coin.THUNDERTOKEN,
		coin.GOCHAIN,
		coin.TOMOCHAIN,
		coin.SMARTCHAIN,
		coin.SOLANA,
		coin.POLYGON,
		coin.OPTIMISM,
		coin.XDAI,
		coin.AVALANCHEC,
		coin.FANTOM,
		coin.HECO:
		return string(GetEthereumTokenTypeByIndex(c)), true
	case coin.TRON:
		_, err := strconv.Atoi(tokenID)
		if err != nil {
			return string(TRC20), true
		}
		return string(TRC10), true
	case coin.BINANCE:
		return string(BEP2), true
	default:
		return "", false
	}
}

// todo test
func GetEthereumTokenTypeByIndex(coinIndex uint) TokenType {
	var tokenType TokenType
	switch coinIndex {
	case coin.ETHEREUM:
		tokenType = ERC20
	case coin.CLASSIC:
		tokenType = ETC20
	case coin.POA:
		tokenType = POA20
	case coin.CALLISTO:
		tokenType = CLO20
	case coin.WANCHAIN:
		tokenType = WAN20
	case coin.THUNDERTOKEN:
		tokenType = TT20
	case coin.GOCHAIN:
		tokenType = GO20
	case coin.TOMOCHAIN:
		tokenType = TRC21
	case coin.SMARTCHAIN:
		tokenType = BEP20
	case coin.SOLANA:
		tokenType = SPL
	case coin.POLYGON:
		tokenType = POLYGON
	case coin.OPTIMISM:
		tokenType = OPTIMISM
	case coin.XDAI:
		tokenType = XDAI
	case coin.AVALANCHEC:
		tokenType = AVALANCHE
	case coin.FANTOM:
		tokenType = FANTOM
	case coin.HECO:
		tokenType = HRC20
	default:
		tokenType = ERC20
	}
	return tokenType
}

func (t Token) AssetId() string {
	return asset.BuildID(t.Coin, t.TokenID)
}
