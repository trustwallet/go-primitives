package types

import (
	"errors"

	"github.com/trustwallet/go-primitives/coin"
)

func GetChainFromAssetType(assetType string) (coin.Coin, error) {
	switch TokenType(assetType) {
	case ERC20:
		return coin.Ethereum(), nil
	case BEP2:
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
	case POA:
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
	case TERRA:
		return coin.Terra(), nil
	case RONIN:
		return coin.Coin{Handle: "ronin"}, nil
	}

	return coin.Coin{}, errors.New("unknown asset type: " + assetType)
}
