package coin

import "errors"

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
		HECO:
		return true
	}

	return false
}
