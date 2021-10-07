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

func GetCoinExploreURL(chain, addr string) (string, error) {
	switch strings.ToLower(chain) {
	case Coins[ETHEREUM].Handle:
		return fmt.Sprintf("https://etherscan.io/token/%s", addr), nil
	case Coins[TRON].Handle:
		if strings.HasPrefix(addr, "10") {
			return fmt.Sprintf("https://tronscan.io/#/token/%s", addr), nil
		}

		return fmt.Sprintf("https://tronscan.io/#/token20/%s", addr), nil
	case Coins[BINANCE].Handle:
		return fmt.Sprintf("https://explorer.binance.org/asset/%s", addr), nil
	case Coins[SMARTCHAIN].Handle:
		return fmt.Sprintf("https://bscscan.com/token/%s", addr), nil
	case Coins[EOS].Handle:
		return fmt.Sprintf("https://bloks.io/account/%s", addr), nil
	case Coins[NEO].Handle:
		return fmt.Sprintf("https://neo.tokenview.com/en/token/0x%s", addr), nil
	case Coins[NULS].Handle:
		return fmt.Sprintf("https://nulscan.io/token/info?contractAddress=%s", addr), nil
	case Coins[WANCHAIN].Handle:
		return fmt.Sprintf("https://www.wanscan.org/token/%s", addr), nil
	case Coins[SOLANA].Handle:
		return fmt.Sprintf("https://explorer.solana.com/address/%s", addr), nil
	case Coins[TOMOCHAIN].Handle:
		return fmt.Sprintf("https://scan.tomochain.com/address/%s", addr), nil
	case Coins[KAVA].Handle:
		return "https://www.mintscan.io/kava", nil
	case Coins[ONTOLOGY].Handle:
		return "https://explorer.ont.io", nil
	case Coins[GOCHAIN].Handle:
		return fmt.Sprintf("https://explorer.gochain.io/addr/%s", addr), nil
	case Coins[THETA].Handle:
		return "https://explorer.thetatoken.org/", nil
	case Coins[THUNDERTOKEN].Handle:
		return fmt.Sprintf("https://viewblock.io/thundercore/address/%s", addr), nil
	case Coins[CLASSIC].Handle:
		return fmt.Sprintf("https://blockscout.com/etc/mainnet/tokens/%s", addr), nil
	case Coins[VECHAIN].Handle:
		return fmt.Sprintf("https://explore.vechain.org/accounts/%s", addr), nil
	case Coins[WAVES].Handle:
		return fmt.Sprintf("https://wavesexplorer.com/assets/%s", addr), nil
	case Coins[XDAI].Handle:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/tokens/%s", addr), nil
	case Coins[POA].Handle:
		return fmt.Sprintf("https://blockscout.com/poa/core/tokens/%s", addr), nil
	case Coins[POLYGON].Handle:
		return fmt.Sprintf("https://polygonscan.com/token/%s", addr), nil
	case Coins[OPTIMISM].Handle:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/tokens/%s", addr), nil
	case Coins[AVALANCHEC].Handle:
		return fmt.Sprintf("https://cchain.explorer.avax.network/address/%s", addr), nil
	case Coins[ARBITRUM].Handle:
		return fmt.Sprintf("https://arbiscan.io/token/%s", addr), nil
	case Coins[FANTOM].Handle:
		return fmt.Sprintf("https://ftmscan.com/token/%s", addr), nil
	case Coins[TERRA].Handle:
		return fmt.Sprintf("https://finder.terra.money/columbus-4/%s", addr), nil
	}

	return "", errors.New("unknown chain: " + chain)
}
