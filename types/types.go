package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"
)

type HexNumber big.Int

func (i HexNumber) MarshalJSON() ([]byte, error) {
	hexNumber := fmt.Sprintf("\"0x%x\"", (*big.Int)(&i))

	return []byte(hexNumber), nil
}

func (i *HexNumber) UnmarshalJSON(data []byte) error {
	var resultStr string
	err := json.Unmarshal(data, &resultStr)
	if err != nil {
		return err
	}

	var value *big.Int
	if resultStr == "0x" {
		value = new(big.Int)
	} else {
		hex := strings.Replace(resultStr, "0x", "", 1)

		var ok bool
		value, ok = new(big.Int).SetString(hex, 16)
		if !ok {
			return fmt.Errorf("could not parse hex value %v", resultStr)
		}
	}

	*i = HexNumber(*value)

	return nil
}

var Zero = big.NewInt(0)

// ToBig converts HexNumber to *big.Int.
// if HexNumber is nil, it returns a pointer to Zero.
func (i *HexNumber) ToBig() *big.Int {
	if i == nil {
		return Zero
	}

	return (*big.Int)(i)
}

// HexNumberToUnix converts HexNumber to time.Time by interpreting it as a Unix timestamp.
func HexNumberToUnix(n *HexNumber) time.Time {
	return time.Unix(n.ToBig().Int64(), 0)
}
