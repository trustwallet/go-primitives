package numbers

import (
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func Float64toPrecision(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

// 0.1010 => "0.101"
func Float64toString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

// "0.00037500" => 0.000375, "non-string-number" => 0
func StringNumberToFloat64(str string) (float64, error) {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	} else {
		return value, nil
	}
}

func FromDecimal(dec string) string {
	v, err := DecimalToSatoshis(dec)
	if err != nil {
		return "0"
	}
	return v
}

func ToDecimal(value string, exp int) string {
	num, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "0"
	}
	denom := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(exp)), nil)
	rat := new(big.Rat).SetFrac(num, denom)
	f, err := decimal.NewFromString(rat.FloatString(10))
	if err != nil {
		return "0"
	}
	return f.String()
}

func FromDecimalExp(dec string, exp int) string {
	return strings.Split(DecimalExp(dec, exp), ".")[0]
}

func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
