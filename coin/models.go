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
	return Coin{}, errors.New("unknown id: " + id)
}

func GetCoinExploreURL(c Coin, addr string) (string, error) {
	switch c.ID {
	case ETHEREUM:
		return fmt.Sprintf("https://etherscan.io/token/%s", addr), nil
	case TRON:
		if _, err := strconv.ParseUint(addr, 10, 64); err == nil {
			return fmt.Sprintf("https://tronscan.io/#/token/%s", addr), nil
		}

		return fmt.Sprintf("https://tronscan.io/#/token20/%s", addr), nil
	case BINANCE:
		return fmt.Sprintf("https://explorer.binance.org/asset/%s", addr), nil
	case SMARTCHAIN:
		return fmt.Sprintf("https://bscscan.com/token/%s", addr), nil
	case EOS:
		return fmt.Sprintf("https://bloks.io/account/%s", addr), nil
	case NEO:
		return fmt.Sprintf("https://neo.tokenview.com/en/token/0x%s", addr), nil
	case NULS:
		return fmt.Sprintf("https://nulscan.io/token/info?contractAddress=%s", addr), nil
	case WANCHAIN:
		return fmt.Sprintf("https://www.wanscan.org/token/%s", addr), nil
	case SOLANA:
		return fmt.Sprintf("https://explorer.solana.com/address/%s", addr), nil
	case TOMOCHAIN:
		return fmt.Sprintf("https://scan.tomochain.com/address/%s", addr), nil
	case KAVA:
		return "https://www.mintscan.io/kava", nil
	case ONTOLOGY:
		return "https://explorer.ont.io", nil
	case GOCHAIN:
		return fmt.Sprintf("https://explorer.gochain.io/addr/%s", addr), nil
	case THETA:
		return "https://explorer.thetatoken.org/", nil
	case THUNDERTOKEN:
		return fmt.Sprintf("https://viewblock.io/thundercore/address/%s", addr), nil
	case CLASSIC:
		return fmt.Sprintf("https://blockscout.com/etc/mainnet/tokens/%s", addr), nil
	case VECHAIN:
		return fmt.Sprintf("https://explore.vechain.org/accounts/%s", addr), nil
	case WAVES:
		return fmt.Sprintf("https://wavesexplorer.com/assets/%s", addr), nil
	case XDAI:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/tokens/%s", addr), nil
	case POA:
		return fmt.Sprintf("https://blockscout.com/poa/core/tokens/%s", addr), nil
	case POLYGON:
		return fmt.Sprintf("https://polygonscan.com/token/%s", addr), nil
	case OPTIMISM:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/tokens/%s", addr), nil
	case AVALANCHEC:
		return fmt.Sprintf("https://cchain.explorer.avax.network/address/%s", addr), nil
	case ARBITRUM:
		return fmt.Sprintf("https://arbiscan.io/token/%s", addr), nil
	case FANTOM:
		return fmt.Sprintf("https://ftmscan.com/token/%s", addr), nil
	case TERRA:
		return fmt.Sprintf("https://finder.terra.money/columbus-4/%s", addr), nil
	}

	return "", errors.New("no explorer for coin: " + c.Handle)
}
