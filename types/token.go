package types

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/trustwallet/go-primitives/asset"
	"github.com/trustwallet/go-primitives/coin"
)

var ErrUnknownTokenType = errors.New("unknown token type")
var errTokenVersionNotImplemented = errors.New("token version not implemented")

type (
	TokenType    string
	TokenVersion int

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
	GO20      TokenType = "GO20"
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
	RONIN     TokenType = "RONIN"
	EOS       TokenType = "EOS"
	NEP5      TokenType = "NEP5"
	NRC20     TokenType = "NRC20"
	VET       TokenType = "VET"
	ONTOLOGY  TokenType = "ONTOLOGY"
	THETA     TokenType = "THETA"
	TOMO      TokenType = "TOMO"
	WAVES     TokenType = "WAVES"
	POA       TokenType = "POA"
	CELO      TokenType = "CELO"
	ESDT      TokenType = "ESDT"
	CW20      TokenType = "CW20"
	OASIS     TokenType = "OASIS"
	CRC20     TokenType = "CRC20"
	STELLAR   TokenType = "STELLAR"
	KRC20     TokenType = "KRC20"
	AURORA    TokenType = "AURORA"
	ALGORAND  TokenType = "ALGORAND"
)

const (
	TokenVersionV0        TokenVersion = 0
	TokenVersionV1        TokenVersion = 1
	TokenVersionV3        TokenVersion = 3
	TokenVersionV4        TokenVersion = 4
	TokenVersionV5        TokenVersion = 5
	TokenVersionV6        TokenVersion = 6
	TokenVersionV7        TokenVersion = 7
	TokenVersionV8        TokenVersion = 8
	TokenVersionV9        TokenVersion = 9
	TokenVersionV10       TokenVersion = 10
	TokenVersionV11       TokenVersion = 11
	TokenVersionUndefined TokenVersion = -1
)

func GetTokenTypes() []TokenType {
	return []TokenType{
		ERC20,
		ERC721,
		ERC1155,
		BEP2,
		BEP8,
		BEP20,
		TRC10,
		ETC20,
		POA20,
		TRC20,
		TRC21,
		CLO20,
		GO20,
		WAN20,
		TT20,
		CW20,
		KAVA,
		SPL,
		POLYGON,
		OPTIMISM,
		XDAI,
		AVALANCHE,
		FANTOM,
		HRC20,
		ARBITRUM,
		TERRA,
		RONIN,
		EOS,
		NEP5,
		NRC20,
		VET,
		ONTOLOGY,
		THETA,
		TOMO,
		WAVES,
		POA,
		CELO,
		ESDT,
		OASIS,
		CRC20,
		STELLAR,
		KRC20,
		AURORA,
		ALGORAND,
	}
}

func GetTokenType(c uint, tokenID string) (string, bool) {
	if coin.IsEVM(c) {
		tokenType, err := GetEthereumTokenTypeByIndex(c)
		if err != nil {
			return "", false
		}
		return string(tokenType), true
	}

	switch c {
	case coin.TRON:
		_, err := strconv.Atoi(tokenID)
		if err != nil {
			return string(TRC20), true
		}
		return string(TRC10), true
	case coin.TERRA:
		idSize := len(tokenID)
		if idSize == 44 {
			return string(CW20), true
		}
		return string(TERRA), true
	case coin.BINANCE:
		return string(BEP2), true
	case coin.WAVES:
		return string(WAVES), true
	case coin.THETA:
		return string(THETA), true
	case coin.ONTOLOGY:
		return string(ONTOLOGY), true
	case coin.NULS:
		return string(NRC20), true
	case coin.VECHAIN:
		return string(VET), true
	case coin.NEO:
		return string(NEP5), true
	case coin.EOS:
		return string(EOS), true
	case coin.SOLANA:
		return string(SPL), true
	case coin.HECO, coin.HARMONY:
		return string(HRC20), true
	case coin.OASIS:
		return string(OASIS), true
	case coin.STELLAR:
		return string(STELLAR), true
	case coin.ALGORAND:
		return string(ALGORAND), true
	default:
		return "", false
	}
}

func GetTokenVersion(tokenType string) (TokenVersion, error) {
	parsedTokenType, err := ParseTokenTypeFromString(tokenType)
	if err != nil {
		return TokenVersionUndefined, err
	}
	switch parsedTokenType {
	case ERC20,
		BEP2,
		BEP20,
		BEP8,
		ETC20,
		POA20,
		CLO20,
		TRC10,
		TRC21,
		WAN20,
		GO20,
		TT20,
		WAVES:
		return TokenVersionV0, nil
	case TRC20:
		return TokenVersionV1, nil
	case SPL, KAVA:
		return TokenVersionV3, nil
	case POLYGON:
		return TokenVersionV4, nil
	case AVALANCHE, ARBITRUM, FANTOM, HRC20, OPTIMISM, XDAI:
		return TokenVersionV5, nil
	case TERRA:
		return TokenVersionV6, nil
	case CELO, NRC20:
		return TokenVersionV7, nil
	case CW20:
		return TokenVersionV8, nil
	case ESDT, CRC20:
		return TokenVersionV9, nil
	case KRC20, STELLAR:
		return TokenVersionV10, nil
	case RONIN, AURORA:
		return TokenVersionV11, nil
	case ERC721, ERC1155, EOS, NEP5, VET, ONTOLOGY, THETA, TOMO, POA, OASIS, ALGORAND:
		return TokenVersionUndefined, nil
	default:
		// This should not happen, as it is guarded by TestGetTokenVersionImplementEverySupportedTokenTypes
		return TokenVersionUndefined, fmt.Errorf("tokenType %s: %w", parsedTokenType, errTokenVersionNotImplemented)
	}
}

func ParseTokenTypeFromString(t string) (TokenType, error) {
	for _, knownTokenType := range GetTokenTypes() {
		if string(knownTokenType) == t {
			return knownTokenType, nil
		}
	}
	return "", fmt.Errorf("%s: %w", t, ErrUnknownTokenType)
}

func GetEthereumTokenTypeByIndex(coinIndex uint) (TokenType, error) {
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
	case coin.RONIN:
		tokenType = RONIN
	case coin.CELO:
		tokenType = CELO
	case coin.CRONOS:
		tokenType = CRC20
	case coin.KCC:
		tokenType = KRC20
	case coin.AURORA:
		tokenType = AURORA
	}

	if tokenType == "" {
		return "", fmt.Errorf("not evm coin %d", coinIndex)
	}

	return tokenType, nil
}

func (t Token) AssetId() string {
	return asset.BuildID(t.Coin, t.TokenID)
}
