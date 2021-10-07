package asset

import "errors"

func GetChainFromAssetType(type_ string) (string, error) {
	switch type_ {
	case "ERC20":
		return "ethereum", nil
	case "BEP2":
		return "binance", nil
	case "BEP20":
		return "smartchain", nil
	case "ETC20":
		return "classic", nil
	case "TRC10", "TRC20":
		return "tron", nil
	case "WAN20":
		return "wanchain", nil
	case "TRC21":
		return "tomochain", nil
	case "TT20":
		return "thundertoken", nil
	case "SPL":
		return "solana", nil
	case "EOS":
		return "eos", nil
	case "GO20":
		return "gochain", nil
	case "KAVA":
		return "kava", nil
	case "NEP5":
		return "neo", nil
	case "NRC20":
		return "nuls", nil
	case "VET":
		return "vechain", nil
	case "ONTOLOGY":
		return "ontology", nil
	case "THETA":
		return "theta", nil
	case "TOMO":
		return "tomochain", nil
	case "XDAI":
		return "xdai", nil
	case "WAVES":
		return "waves", nil
	case "POA":
		return "poa", nil
	case "POLYGON":
		return "polygon", nil
	case "OPTIMISM":
		return "optimism", nil
	case "AVALANCHE":
		return "avalanchec", nil
	case "ARBITRUM":
		return "arbitrum", nil
	case "FANTOM":
		return "fantom", nil
	case "TERRA":
		return "terra", nil
	}

	return "", errors.New("unknown asset type: " + type_)
}
