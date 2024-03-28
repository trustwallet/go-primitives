package types

import (
	"errors"

	"github.com/trustwallet/go-primitives/coin"
)

// nolint:cyclop
// Deprecated: Multiple chains can have the same type, use GetChainsFromAssetType instead
func GetChainFromAssetType(assetType string) (coin.Coin, error) {
	switch TokenType(assetType) {
	case BRC20:
		return coin.Bitcoin(), nil
	case ERC20:
		return coin.Ethereum(), nil
	case BEP2, BEP8:
		return coin.Binance(), nil
	case BEP20:
		return coin.Smartchain(), nil
	case ETC20:
		return coin.Classic(), nil
	case TRC10, TRC20:
		return coin.Tron(), nil
	case WAN20:
		return coin.Wanchain(), nil
	case TT20:
		return coin.Thundertoken(), nil
	case SPL:
		return coin.Solana(), nil
	case EOS:
		return coin.Eos(), nil
	case GO20:
		return coin.Gochain(), nil
	case KAVA:
		return coin.Kava(), nil
	case NEP5:
		return coin.Neo(), nil
	case NRC20:
		return coin.Nuls(), nil
	case VET:
		return coin.Vechain(), nil
	case ONTOLOGY:
		return coin.Ontology(), nil
	case THETA:
		return coin.Theta(), nil
	case TOMO, TRC21:
		return coin.Tomochain(), nil
	case XDAI:
		return coin.Xdai(), nil
	case WAVES:
		return coin.Waves(), nil
	case POA, POA20:
		return coin.Poa(), nil
	case POLYGON:
		return coin.Polygon(), nil
	case OPTIMISM:
		return coin.Optimism(), nil
	case AVALANCHE:
		return coin.Avalanchec(), nil
	case ARBITRUM:
		return coin.Arbitrum(), nil
	case FANTOM:
		return coin.Fantom(), nil
	case TERRA, CW20:
		return coin.Terra(), nil
	case RONIN:
		return coin.Ronin(), nil
	case CELO:
		return coin.Celo(), nil
	case HRC20:
		return coin.Heco(), nil
	case CLO20:
		return coin.Callisto(), nil
	case ESDT:
		return coin.Elrond(), nil
	case OASIS:
		return coin.Oasis(), nil
	case CRC20:
		return coin.Cronos(), nil
	case STELLAR:
		return coin.Stellar(), nil
	case KRC20:
		return coin.Kcc(), nil
	case AURORA:
		return coin.Aurora(), nil
	case ALGORAND:
		return coin.Algorand(), nil
	case KAVAEVM:
		return coin.Kavaevm(), nil
	case METER:
		return coin.Meter(), nil
	case EVMOS_ERC20:
		return coin.Evmos(), nil
	case KIP20:
		return coin.Okc(), nil
	case APTOS:
		return coin.Aptos(), nil
	case MOONBEAM:
		return coin.Moonbeam(), nil
	case KLAYTN:
		return coin.Klaytn(), nil
	case METIS:
		return coin.Metis(), nil
	case MOONRIVER:
		return coin.Moonriver(), nil
	case BOBA:
		return coin.Boba(), nil
	case TON:
		return coin.Ton(), nil
	case POLYGONZKEVM:
		return coin.Polygonzkevm(), nil
	case ZKSYNC:
		return coin.Zksync(), nil
	case SUI:
		return coin.Sui(), nil
	case STRIDE:
		return coin.Stride(), nil
	case NEUTRON:
		return coin.Neutron(), nil
	case FA2:
		return coin.Tezos(), nil
	case CONFLUX:
		return coin.Cfxevm(), nil
	case ACA:
		return coin.Acala(), nil
	case ACALAEVM:
		return coin.Acalaevm(), nil
	case BASE:
		return coin.Base(), nil
	case AKASH:
		return coin.Akash(), nil
	case AGORIC:
		return coin.Agoric(), nil
	case AXELAR:
		return coin.Axelar(), nil
	case JUNO:
		return coin.Juno(), nil
	case SEI:
		return coin.Sei(), nil
	case CARDANO:
		return coin.Cardano(), nil
	case NEON:
		return coin.Neon(), nil
	case OSMOSIS:
		return coin.Osmosis(), nil
	case NATIVEINJECTIVE:
		return coin.Nativeinjective(), nil
	case NATIVEEVMOS:
		return coin.Nativeevmos(), nil
	case CRYPTOORG:
		return coin.Cryptoorg(), nil
	case COSMOS:
		return coin.Cosmos(), nil
	case OPBNB:
		return coin.Opbnb(), nil
	case LINEA:
		return coin.Linea(), nil
	case STARGAZE:
		return coin.Stargaze(), nil
	case MANTLE:
		return coin.Mantle(), nil
	case MANTA:
		return coin.Manta(), nil
	case ZETACHAIN:
		return coin.Zetachain(), nil
	case ZETAEVM:
		return coin.Zetaevm(), nil
	}

	return coin.Coin{}, errors.New("unknown asset type: " + assetType)
}

// nolint:cyclop
func GetChainsFromAssetType(assetType string) ([]coin.Coin, error) {
	switch TokenType(assetType) {
	case BRC20:
		return []coin.Coin{coin.Bitcoin()}, nil
	case ERC20:
		return []coin.Coin{coin.Ethereum()}, nil
	case BEP2, BEP8:
		return []coin.Coin{coin.Binance()}, nil
	case BEP20:
		return []coin.Coin{coin.Smartchain()}, nil
	case ETC20:
		return []coin.Coin{coin.Classic()}, nil
	case TRC10, TRC20:
		return []coin.Coin{coin.Tron()}, nil
	case WAN20:
		return []coin.Coin{coin.Wanchain()}, nil
	case TT20:
		return []coin.Coin{coin.Thundertoken()}, nil
	case SPL:
		return []coin.Coin{coin.Solana()}, nil
	case EOS:
		return []coin.Coin{coin.Eos()}, nil
	case GO20:
		return []coin.Coin{coin.Gochain()}, nil
	case KAVA:
		return []coin.Coin{coin.Kava()}, nil
	case NEP5:
		return []coin.Coin{coin.Neo()}, nil
	case NRC20:
		return []coin.Coin{coin.Nuls()}, nil
	case VET:
		return []coin.Coin{coin.Vechain()}, nil
	case ONTOLOGY:
		return []coin.Coin{coin.Ontology()}, nil
	case THETA:
		return []coin.Coin{coin.Theta()}, nil
	case TOMO, TRC21:
		return []coin.Coin{coin.Tomochain()}, nil
	case XDAI:
		return []coin.Coin{coin.Xdai()}, nil
	case WAVES:
		return []coin.Coin{coin.Waves()}, nil
	case POA, POA20:
		return []coin.Coin{coin.Poa()}, nil
	case POLYGON:
		return []coin.Coin{coin.Polygon()}, nil
	case OPTIMISM:
		return []coin.Coin{coin.Optimism()}, nil
	case AVALANCHE:
		return []coin.Coin{coin.Avalanchec()}, nil
	case ARBITRUM:
		return []coin.Coin{coin.Arbitrum()}, nil
	case FANTOM:
		return []coin.Coin{coin.Fantom()}, nil
	case TERRA:
		return []coin.Coin{coin.Terra()}, nil
	case CW20:
		return []coin.Coin{coin.Terra(), coin.Nativeinjective()}, nil
	case RONIN:
		return []coin.Coin{coin.Ronin()}, nil
	case CELO:
		return []coin.Coin{coin.Celo()}, nil
	case HRC20:
		return []coin.Coin{coin.Heco()}, nil
	case CLO20:
		return []coin.Coin{coin.Callisto()}, nil
	case ESDT:
		return []coin.Coin{coin.Elrond()}, nil
	case OASIS:
		return []coin.Coin{coin.Oasis()}, nil
	case CRC20:
		return []coin.Coin{coin.Cronos()}, nil
	case STELLAR:
		return []coin.Coin{coin.Stellar()}, nil
	case KRC20:
		return []coin.Coin{coin.Kcc()}, nil
	case AURORA:
		return []coin.Coin{coin.Aurora()}, nil
	case ALGORAND:
		return []coin.Coin{coin.Algorand()}, nil
	case KAVAEVM:
		return []coin.Coin{coin.Kavaevm()}, nil
	case METER:
		return []coin.Coin{coin.Meter()}, nil
	case EVMOS_ERC20:
		return []coin.Coin{coin.Evmos()}, nil
	case KIP20:
		return []coin.Coin{coin.Okc()}, nil
	case APTOS:
		return []coin.Coin{coin.Aptos()}, nil
	case MOONBEAM:
		return []coin.Coin{coin.Moonbeam()}, nil
	case KLAYTN:
		return []coin.Coin{coin.Klaytn()}, nil
	case METIS:
		return []coin.Coin{coin.Metis()}, nil
	case MOONRIVER:
		return []coin.Coin{coin.Moonriver()}, nil
	case BOBA:
		return []coin.Coin{coin.Boba()}, nil
	case TON:
		return []coin.Coin{coin.Ton()}, nil
	case POLYGONZKEVM:
		return []coin.Coin{coin.Polygonzkevm()}, nil
	case ZKSYNC:
		return []coin.Coin{coin.Zksync()}, nil
	case SUI:
		return []coin.Coin{coin.Sui()}, nil
	case STRIDE:
		return []coin.Coin{coin.Stride()}, nil
	case NEUTRON:
		return []coin.Coin{coin.Neutron()}, nil
	case FA2:
		return []coin.Coin{coin.Tezos()}, nil
	case CONFLUX:
		return []coin.Coin{coin.Cfxevm()}, nil
	case ACA:
		return []coin.Coin{coin.Acala()}, nil
	case ACALAEVM:
		return []coin.Coin{coin.Acalaevm()}, nil
	case BASE:
		return []coin.Coin{coin.Base()}, nil
	case AKASH:
		return []coin.Coin{coin.Akash()}, nil
	case AGORIC:
		return []coin.Coin{coin.Agoric()}, nil
	case AXELAR:
		return []coin.Coin{coin.Axelar()}, nil
	case JUNO:
		return []coin.Coin{coin.Juno()}, nil
	case SEI:
		return []coin.Coin{coin.Sei()}, nil
	case CARDANO:
		return []coin.Coin{coin.Cardano()}, nil
	case NEON:
		return []coin.Coin{coin.Neon()}, nil
	case OSMOSIS:
		return []coin.Coin{coin.Osmosis()}, nil
	case NATIVEEVMOS:
		return []coin.Coin{coin.Nativeevmos()}, nil
	case CRYPTOORG:
		return []coin.Coin{coin.Cryptoorg()}, nil
	case COSMOS:
		return []coin.Coin{coin.Cosmos()}, nil
	case OPBNB:
		return []coin.Coin{coin.Opbnb()}, nil
	case LINEA:
		return []coin.Coin{coin.Linea()}, nil
	case STARGAZE:
		return []coin.Coin{coin.Stargaze()}, nil
	case MANTLE:
		return []coin.Coin{coin.Mantle()}, nil
	case MANTA:
		return []coin.Coin{coin.Manta()}, nil
	case ZETACHAIN:
		return []coin.Coin{coin.Zetachain()}, nil
	case ZETAEVM:
		return []coin.Coin{coin.Zetaevm()}, nil
	}

	return nil, errors.New("unknown asset type: " + assetType)
}
