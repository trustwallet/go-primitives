package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

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
	Coin            TokenType = "coin"
	Gas             TokenType = "gas"
	BRC20           TokenType = "BRC20"
	ERC20           TokenType = "ERC20"
	ERC721          TokenType = "ERC721"
	ERC1155         TokenType = "ERC1155"
	BEP2            TokenType = "BEP2"
	BEP8            TokenType = "BEP8"
	BEP20           TokenType = "BEP20"
	TRC10           TokenType = "TRC10"
	ETC20           TokenType = "ETC20"
	POA20           TokenType = "POA20"
	TRC20           TokenType = "TRC20"
	TRC21           TokenType = "TRC21"
	CLO20           TokenType = "CLO20"
	GO20            TokenType = "GO20"
	WAN20           TokenType = "WAN20"
	TT20            TokenType = "TT20"
	KAVA            TokenType = "KAVA"
	COSMOS          TokenType = "COSMOS"
	CRYPTOORG       TokenType = "CRYPTOORG"
	NATIVEEVMOS     TokenType = "NATIVEEVMOS"
	NATIVEINJECTIVE TokenType = "INJECTIVE"
	STARGAZE        TokenType = "STARGAZE"
	NEUTRON         TokenType = "NEUTRON"
	OSMOSIS         TokenType = "OSMOSIS"
	SPL             TokenType = "SPL"
	POLYGON         TokenType = "POLYGON"
	OPTIMISM        TokenType = "OPTIMISM"
	XDAI            TokenType = "XDAI"
	AVALANCHE       TokenType = "AVALANCHE"
	FANTOM          TokenType = "FANTOM"
	HRC20           TokenType = "HRC20"
	ARBITRUM        TokenType = "ARBITRUM"
	TERRA           TokenType = "TERRA"
	RONIN           TokenType = "RONIN"
	EOS             TokenType = "EOS"
	NEP5            TokenType = "NEP5"
	NRC20           TokenType = "NRC20"
	VET             TokenType = "VET"
	ONTOLOGY        TokenType = "ONTOLOGY"
	THETA           TokenType = "THETA"
	TOMO            TokenType = "TOMO"
	WAVES           TokenType = "WAVES"
	POA             TokenType = "POA"
	CELO            TokenType = "CELO"
	ESDT            TokenType = "ESDT"
	CW20            TokenType = "CW20"
	OASIS           TokenType = "OASIS"
	CRC20           TokenType = "CRC20"
	STELLAR         TokenType = "STELLAR"
	KRC20           TokenType = "KRC20"
	AURORA          TokenType = "AURORA"
	ALGORAND        TokenType = "ASA"
	KAVAEVM         TokenType = "KAVAEVM"
	METER           TokenType = "METER"
	EVMOS_ERC20     TokenType = "EVMOS_ERC20"
	KIP20           TokenType = "KIP20"
	APTOS           TokenType = "APTOS"
	APTOSFA         TokenType = "APTOSFA"
	MOONBEAM        TokenType = "MOONBEAM"
	KLAYTN          TokenType = "KAIA"
	METIS           TokenType = "METIS"
	MOONRIVER       TokenType = "MOONRIVER"
	BOBA            TokenType = "BOBA"
	JETTON          TokenType = "JETTON"
	POLYGONZKEVM    TokenType = "ZKEVM"
	ZKSYNC          TokenType = "ZKSYNC"
	SUI             TokenType = "SUI"
	STRIDE          TokenType = "STRIDE"
	FA2             TokenType = "FA2"
	CONFLUX         TokenType = "CONFLUX"
	ACA             TokenType = "ACA"
	ACALAEVM        TokenType = "ACALAEVM"
	BASE            TokenType = "BASE"
	AKASH           TokenType = "AKT"
	AGORIC          TokenType = "BLD"
	AXELAR          TokenType = "AXL"
	JUNO            TokenType = "JUNO"
	SEI             TokenType = "SEI"
	CARDANO         TokenType = "CARDANO"
	NEON            TokenType = "NEON"
	IOTEXEVM        TokenType = "XRC20"
	OPBNB           TokenType = "OPBNB"
	LINEA           TokenType = "LINEA"
	MANTLE          TokenType = "MANTLE"
	MANTA           TokenType = "MANTA"
	ZETACHAIN       TokenType = "ZETACHAIN"
	ZETAEVM         TokenType = "ZETAEVM"
	MERLIN          TokenType = "MERLIN"
	BLAST           TokenType = "BLAST"
	SCROLL          TokenType = "SCROLL"
	ICP             TokenType = "ICP"
	BOUNCEBIT       TokenType = "BOUNCEBIT"
	ZKLINKNOVA      TokenType = "ZKLINKNOVA"
	XRP             TokenType = "XRP"
	SONIC           TokenType = "SONIC"
	TIA             TokenType = "TIA"
	DYDX            TokenType = "DYDX"
	PLASMA          TokenType = "PLASMA"
	MONAD           TokenType = "MONAD"
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
	TokenVersionV12       TokenVersion = 12
	TokenVersionV13       TokenVersion = 13
	TokenVersionV14       TokenVersion = 14
	TokenVersionV15       TokenVersion = 15
	TokenVersionV16       TokenVersion = 16
	TokenVersionV17       TokenVersion = 17
	TokenVersionV18       TokenVersion = 18
	TokenVersionV19       TokenVersion = 19
	TokenVersionV20       TokenVersion = 20
	TokenVersionV21       TokenVersion = 21
	TokenVersionV22       TokenVersion = 22
	TokenVersionV23       TokenVersion = 23
	TokenVersionV24       TokenVersion = 24
	TokenVersionUndefined TokenVersion = -1
)

func GetTokenTypes() []TokenType {
	return []TokenType{
		BRC20,
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
		COSMOS,
		KAVA,
		JUNO,
		AGORIC,
		AKASH,
		AXELAR,
		CRYPTOORG,
		NATIVEEVMOS,
		NATIVEINJECTIVE,
		OSMOSIS,
		STARGAZE,
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
		KAVAEVM,
		METER,
		EVMOS_ERC20,
		KIP20,
		APTOS,
		APTOSFA,
		MOONBEAM,
		KLAYTN,
		METIS,
		MOONRIVER,
		BOBA,
		JETTON,
		POLYGONZKEVM,
		ZKSYNC,
		SUI,
		STRIDE,
		NEUTRON,
		CONFLUX,
		ACA,
		BASE,
		SEI,
		CARDANO,
		NEON,
		OPBNB,
		LINEA,
		ACALAEVM,
		MANTLE,
		MANTA,
		ZETACHAIN,
		ZETAEVM,
		MERLIN,
		BLAST,
		SCROLL,
		ICP,
		BOUNCEBIT,
		ZKLINKNOVA,
		XRP,
		SONIC,
		TIA,
		DYDX,
		PLASMA,
		MONAD,
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
	case coin.BITCOIN:
		return string(BRC20), true
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
	case coin.KAVA:
		return string(KAVA), true
	case coin.COSMOS:
		return string(COSMOS), true
	case coin.CRYPTOORG:
		return string(CRYPTOORG), true
	case coin.NATIVEEVMOS:
		return string(NATIVEEVMOS), true
	case coin.NATIVEINJECTIVE:
		return string(NATIVEINJECTIVE), true
	case coin.STARGAZE:
		return string(STARGAZE), true
	case coin.NEUTRON:
		return string(NEUTRON), true
	case coin.OSMOSIS:
		return string(OSMOSIS), true
	case coin.CELO:
		return string(CELO), true
	case coin.ELROND:
		return string(ESDT), true
	case coin.EVMOS:
		return string(EVMOS_ERC20), true
	case coin.OKC:
		return string(KIP20), true
	case coin.APTOS:
		// TODO: improve this
		if strings.Contains(tokenID, "::") {
			return string(APTOS), true
		}
		return string(APTOSFA), true
	case coin.TON:
		return string(JETTON), true
	case coin.SUI:
		return string(SUI), true
	case coin.STRIDE:
		return string(STRIDE), true
	case coin.CFXEVM:
		return string(CONFLUX), true
	case coin.ACALA:
		return string(ACA), true
	case coin.AKASH:
		return string(AKASH), true
	case coin.AGORIC:
		return string(AGORIC), true
	case coin.AXELAR:
		return string(AXELAR), true
	case coin.JUNO:
		return string(JUNO), true
	case coin.SEI:
		return string(SEI), true
	case coin.CARDANO:
		return string(CARDANO), true
	case coin.NEON:
		return string(NEON), true
	case coin.MANTLE:
		return string(MANTLE), true
	case coin.MANTA:
		return string(MANTA), true
	case coin.ZETACHAIN:
		return string(ZETACHAIN), true
	case coin.ZETAEVM:
		return string(ZETAEVM), true
	case coin.BLAST:
		return string(BLAST), true
	case coin.SCROLL:
		return string(SCROLL), true
	case coin.INTERNET_COMPUTER:
		return string(ICP), true
	case coin.BOUNCEBIT:
		return string(BOUNCEBIT), true
	case coin.ZKLINKNOVA:
		return string(ZKLINKNOVA), true
	case coin.RIPPLE:
		return string(XRP), true
	case coin.TIA:
		return string(TIA), true
	case coin.DYDX:
		return string(DYDX), true
	case coin.PLASMA:
		return string(PLASMA), true
	case coin.MONAD:
		return string(MONAD), true
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
		WAVES,
		APTOS:
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
	case JETTON, POLYGONZKEVM, ZKSYNC, SUI:
		return TokenVersionV12, nil
	case BASE, AKASH, AGORIC, AXELAR, JUNO, SEI, OPBNB:
		return TokenVersionV13, nil
	case KAVAEVM, BOBA, METIS, NEON, LINEA, ACA, ACALAEVM, CONFLUX, IOTEXEVM, KLAYTN, MOONRIVER, MOONBEAM, MANTLE,
		NATIVEINJECTIVE, MANTA, ZETACHAIN, ZETAEVM:
		return TokenVersionV14, nil
	case BRC20:
		return TokenVersionV16, nil
	case MERLIN, ICP:
		return TokenVersionV17, nil
	case BLAST, SCROLL:
		return TokenVersionV18, nil
	case BOUNCEBIT:
		return TokenVersionV19, nil
	case ZKLINKNOVA:
		return TokenVersionV20, nil
	case ERC721, ERC1155, EOS, NEP5, VET, ONTOLOGY, THETA, TOMO, POA, OASIS, ALGORAND, METER, EVMOS_ERC20,
		KIP20, STRIDE, NEUTRON, FA2, CARDANO, NATIVEEVMOS, CRYPTOORG, COSMOS, OSMOSIS, STARGAZE, TIA, DYDX:
		return TokenVersionUndefined, nil
	case APTOSFA:
		return TokenVersionV21, nil
	case XRP, SONIC:
		return TokenVersionV22, nil
	case PLASMA:
		return TokenVersionV23, nil
	case MONAD:
		return TokenVersionV24, nil
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
	case coin.ARBITRUM:
		tokenType = ARBITRUM
	case coin.KAVAEVM:
		tokenType = KAVAEVM
	case coin.METER:
		tokenType = METER
	case coin.EVMOS:
		tokenType = EVMOS_ERC20
	case coin.OKC:
		tokenType = KIP20
	case coin.MOONBEAM:
		tokenType = MOONBEAM
	case coin.KLAYTN:
		tokenType = KLAYTN
	case coin.METIS:
		tokenType = METIS
	case coin.MOONRIVER:
		tokenType = MOONRIVER
	case coin.BOBA:
		tokenType = BOBA
	case coin.TON:
		tokenType = JETTON
	case coin.POLYGONZKEVM:
		tokenType = POLYGONZKEVM
	case coin.ZKSYNC:
		tokenType = ZKSYNC
	case coin.SUI:
		tokenType = SUI
	case coin.STRIDE:
		tokenType = STRIDE
	case coin.NEUTRON:
		tokenType = NEUTRON
	case coin.CFXEVM:
		tokenType = CONFLUX
	case coin.ACALAEVM:
		tokenType = ACALAEVM
	case coin.BASE:
		tokenType = BASE
	case coin.NEON:
		tokenType = NEON
	case coin.IOTEXEVM:
		tokenType = IOTEXEVM
	case coin.OPBNB:
		tokenType = OPBNB
	case coin.LINEA:
		tokenType = LINEA
	case coin.MANTLE:
		tokenType = MANTLE
	case coin.MANTA:
		tokenType = MANTA
	case coin.ZETAEVM:
		tokenType = ZETAEVM
	case coin.MERLIN:
		tokenType = MERLIN
	case coin.BLAST:
		tokenType = BLAST
	case coin.SCROLL:
		tokenType = SCROLL
	case coin.BOUNCEBIT:
		tokenType = BOUNCEBIT
	case coin.ZKLINKNOVA:
		tokenType = ZKLINKNOVA
	case coin.SONIC:
		tokenType = SONIC
	case coin.PLASMA:
		tokenType = PLASMA
	case coin.MONAD:
		tokenType = MONAD
	}

	if tokenType == "" {
		return "", fmt.Errorf("not evm coin %d", coinIndex)
	}

	return tokenType, nil
}

func (t Token) AssetId() string {
	return asset.BuildID(t.Coin, t.TokenID)
}
