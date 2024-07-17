package address

import (
	"encoding/hex"
	"errors"
	"strings"

	"golang.org/x/crypto/sha3"

	"github.com/trustwallet/go-primitives/coin"
)

var ErrInvalidInput = errors.New("invalid input")

// Decode decodes a hex string with 0x prefix.
func Remove0x(input string) string {
	if strings.HasPrefix(input, "0x") {
		return input[2:]
	}
	return input
}

// Hex returns an EIP55-compliant hex string representation of the address.
func EIP55Checksum(unchecksummed string) (string, error) {
	v := []byte(Remove0x(strings.ToLower(unchecksummed)))

	_, err := hex.DecodeString(string(v))
	if err != nil {
		return "", err
	}

	sha := sha3.NewLegacyKeccak256()
	_, err = sha.Write(v)
	if err != nil {
		return "", err
	}
	hash := sha.Sum(nil)

	result := v
	if (len(result)-1)/2 >= len(hash) {
		return "", ErrInvalidInput
	}

	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	val := string(result)
	return "0x" + val, nil
}

func ToEIP55ByCoinID(str string, coinID uint) (eip55Addr string, err error) {
	if !coin.IsEVM(coinID) {
		return str, nil
	}

	// special case for ronin addresses
	const roninPrefix, hexPrefix = "ronin:", "0x"
	if coinID == coin.RONIN && strings.HasPrefix(str, roninPrefix) {
		str = hexPrefix + str[len(roninPrefix):]
		defer func() {
			// remove 0x prefix, then add roninPrefix
			eip55Addr = roninPrefix + eip55Addr[len(hexPrefix):]
		}()
	}

	eip55Addr, err = EIP55Checksum(str)
	if err != nil {
		return "", err
	}

	return
}
