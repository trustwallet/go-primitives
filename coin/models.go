package coin

import (
	"errors"
	"fmt"
	"strconv"
)

func GetCoinForId(id string) (Coin, error) {
	for _, c := range Coins {
		if c.Handle == id {
			return c, nil
		}
	}
	return Coin{}, errors.New("unknown id " + id)
}

// todo test
func IsEVM(coinID uint) bool {
	switch coinID {
	case ETHEREUM,
		CLASSIC,
		POA,
		CALLISTO,
		WANCHAIN,
		THUNDERTOKEN,
		GOCHAIN,
		TOMOCHAIN,
		SMARTCHAIN,
		POLYGON,
		OPTIMISM,
		XDAI,
		AVALANCHEC,
		FANTOM,
		HECO,
		RONIN,
		CRONOS,
		KCC,
		AURORA,
		ARBITRUM,
		KAVAEVM,
		METER,
		EVMOS:
		return true
	}

	return false
}

// nolint:cyclop
func GetCoinExploreURL(c Coin, tokenID, tokenType string) (string, error) {
	switch c.ID {
	case ETHEREUM:
		return fmt.Sprintf("https://etherscan.io/token/%s", tokenID), nil
	case TRON:
		if _, err := strconv.ParseUint(tokenID, 10, 64); err == nil {
			return fmt.Sprintf("https://tronscan.io/#/token/%s", tokenID), nil
		}

		return fmt.Sprintf("https://tronscan.io/#/token20/%s", tokenID), nil
	case BINANCE:
		return fmt.Sprintf("https://explorer.binance.org/asset/%s", tokenID), nil
	case SMARTCHAIN:
		return fmt.Sprintf("https://bscscan.com/token/%s", tokenID), nil
	case EOS:
		return fmt.Sprintf("https://bloks.io/account/%s", tokenID), nil
	case NEO:
		return fmt.Sprintf("https://neo.tokenview.com/en/token/0x%s", tokenID), nil
	case NULS:
		return fmt.Sprintf("https://nulscan.io/token/info?contractAddress=%s", tokenID), nil
	case WANCHAIN:
		return fmt.Sprintf("https://www.wanscan.org/token/%s", tokenID), nil
	case SOLANA:
		return fmt.Sprintf("https://solscan.io/token/%s", tokenID), nil
	case TOMOCHAIN:
		return fmt.Sprintf("https://scan.tomochain.com/address/%s", tokenID), nil
	case KAVA:
		return "https://www.mintscan.io/kava", nil
	case ONTOLOGY:
		return "https://explorer.ont.io", nil
	case GOCHAIN:
		return fmt.Sprintf("https://explorer.gochain.io/addr/%s", tokenID), nil
	case THETA:
		return "https://explorer.thetatoken.org/", nil
	case THUNDERTOKEN:
		return fmt.Sprintf("https://viewblock.io/thundercore/address/%s", tokenID), nil
	case CLASSIC:
		return fmt.Sprintf("https://blockscout.com/etc/mainnet/tokens/%s", tokenID), nil
	case VECHAIN:
		return fmt.Sprintf("https://explore.vechain.org/accounts/%s", tokenID), nil
	case WAVES:
		return fmt.Sprintf("https://wavesexplorer.com/assets/%s", tokenID), nil
	case XDAI:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/tokens/%s", tokenID), nil
	case POA:
		return fmt.Sprintf("https://blockscout.com/poa/core/tokens/%s", tokenID), nil
	case POLYGON:
		return fmt.Sprintf("https://polygonscan.com/token/%s", tokenID), nil
	case OPTIMISM:
		return fmt.Sprintf("https://optimistic.etherscan.io/token/%s", tokenID), nil
	case AVALANCHEC:
		return fmt.Sprintf("https://snowtrace.io/address/%s", tokenID), nil
	case ARBITRUM:
		return fmt.Sprintf("https://arbiscan.io/token/%s", tokenID), nil
	case FANTOM:
		return fmt.Sprintf("https://ftmscan.com/token/%s", tokenID), nil
	case TERRA:
		return fmt.Sprintf("https://finder.terra.money/mainnet/address/%s", tokenID), nil
	case RONIN:
		return fmt.Sprintf("https://explorer.roninchain.com/token/%s", tokenID), nil
	case CELO:
		return fmt.Sprintf("https://explorer.bitquery.io/celo_rc1/token/%s", tokenID), nil
	case ELROND:
		if tokenType == "ESDT" {
			return fmt.Sprintf("https://explorer.elrond.com/tokens/%s", tokenID), nil
		}

		return fmt.Sprintf("https://explorer.elrond.com/collections/%s", tokenID), nil
	case HECO:
		return fmt.Sprintf("https://hecoinfo.com/token/%s", tokenID), nil
	case OASIS:
		return fmt.Sprintf("https://explorer.oasis.updev.si/token/%s", tokenID), nil
	case CRONOS:
		return fmt.Sprintf("https://cronos.org/explorer/address/%s/token-transfers", tokenID), nil
	case STELLAR:
		return fmt.Sprintf("https://stellar.expert/explorer/public/asset/%s", tokenID), nil
	case KCC:
		return fmt.Sprintf("https://explorer.kcc.io/token/%s", tokenID), nil
	case AURORA:
		return fmt.Sprintf("https://aurorascan.dev/address/%s", tokenID), nil
	case ALGORAND:
		return fmt.Sprintf("https://algoexplorer.io/asset/%s", tokenID), nil
	case KAVAEVM:
		return fmt.Sprintf("https://explorer.kava.io/token/%s", tokenID), nil
	case METER:
		return fmt.Sprintf("https://scan.meter.io/address/%s", tokenID), nil
	}

	return "", errors.New("no explorer for coin: " + c.Handle)
}
