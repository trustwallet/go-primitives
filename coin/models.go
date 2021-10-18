package coin

import (
	"errors"
	"fmt"
	"strings"
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
	switch c {
	case Coins[ETHEREUM]:
		return fmt.Sprintf("https://etherscan.io/token/%s", addr), nil
	case Coins[TRON]:
		if strings.HasPrefix(addr, "10") {
			return fmt.Sprintf("https://tronscan.io/#/token/%s", addr), nil
		}

		return fmt.Sprintf("https://tronscan.io/#/token20/%s", addr), nil
	case Coins[BINANCE]:
		return fmt.Sprintf("https://explorer.binance.org/asset/%s", addr), nil
	case Coins[SMARTCHAIN]:
		return fmt.Sprintf("https://bscscan.com/token/%s", addr), nil
	case Coins[EOS]:
		return fmt.Sprintf("https://bloks.io/account/%s", addr), nil
	case Coins[NEO]:
		return fmt.Sprintf("https://neo.tokenview.com/en/token/0x%s", addr), nil
	case Coins[NULS]:
		return fmt.Sprintf("https://nulscan.io/token/info?contractAddress=%s", addr), nil
	case Coins[WANCHAIN]:
		return fmt.Sprintf("https://www.wanscan.org/token/%s", addr), nil
	case Coins[SOLANA]:
		return fmt.Sprintf("https://explorer.solana.com/address/%s", addr), nil
	case Coins[TOMOCHAIN]:
		return fmt.Sprintf("https://scan.tomochain.com/address/%s", addr), nil
	case Coins[KAVA]:
		return "https://www.mintscan.io/kava", nil
	case Coins[ONTOLOGY]:
		return "https://explorer.ont.io", nil
	case Coins[GOCHAIN]:
		return fmt.Sprintf("https://explorer.gochain.io/addr/%s", addr), nil
	case Coins[THETA]:
		return "https://explorer.thetatoken.org/", nil
	case Coins[THUNDERTOKEN]:
		return fmt.Sprintf("https://viewblock.io/thundercore/address/%s", addr), nil
	case Coins[CLASSIC]:
		return fmt.Sprintf("https://blockscout.com/etc/mainnet/tokens/%s", addr), nil
	case Coins[VECHAIN]:
		return fmt.Sprintf("https://explore.vechain.org/accounts/%s", addr), nil
	case Coins[WAVES]:
		return fmt.Sprintf("https://wavesexplorer.com/assets/%s", addr), nil
	case Coins[XDAI]:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/tokens/%s", addr), nil
	case Coins[POA]:
		return fmt.Sprintf("https://blockscout.com/poa/core/tokens/%s", addr), nil
	case Coins[POLYGON]:
		return fmt.Sprintf("https://polygonscan.com/token/%s", addr), nil
	case Coins[OPTIMISM]:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/tokens/%s", addr), nil
	case Coins[AVALANCHEC]:
		return fmt.Sprintf("https://cchain.explorer.avax.network/address/%s", addr), nil
	case Coins[ARBITRUM]:
		return fmt.Sprintf("https://arbiscan.io/token/%s", addr), nil
	case Coins[FANTOM]:
		return fmt.Sprintf("https://ftmscan.com/token/%s", addr), nil
	case Coins[TERRA]:
		return fmt.Sprintf("https://finder.terra.money/columbus-4/%s", addr), nil
	}

	return "", errors.New("no explorer for coin: " + c.Handle)
}
