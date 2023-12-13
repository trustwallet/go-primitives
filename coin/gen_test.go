package coin

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

const (
	coinFile = "coins.yml"
	filename = "coins.go"
)

type TestCoin struct {
	ID               uint   `yaml:"id"`
	Handle           string `yaml:"handle"`
	Symbol           string `yaml:"symbol"`
	Name             string `yaml:"name"`
	Decimals         uint   `yaml:"decimals"`
	BlockTime        int    `yaml:"blockTime"`
	MinConfirmations int64  `yaml:"minConfirmations"`
	IsTokenSupported bool   `yaml:"isTokenSupported"`
	SampleAddr       string `yaml:"sampleAddress"`
}

func TestFilesExists(t *testing.T) {
	assert.True(t, assert.FileExists(t, coinFile))
	assert.True(t, assert.FileExists(t, filename))
}

func TestCoinFile(t *testing.T) {
	var coinList []TestCoin
	coin, err := os.Open(coinFile)
	if err != nil {
		t.Error(err)
	}
	dec := yaml.NewDecoder(coin)
	err = dec.Decode(&coinList)
	if err != nil {
		t.Error(err)
	}

	f, err := os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	defer func() { _ = f.Close() }()

	b, err := io.ReadAll(f)
	if err != nil {
		t.Error(err)
	}

	r := regexp.MustCompile(`[ |\t]+`)
	code := string(r.ReplaceAll(b, []byte(" "))) // replace multi spaces and tabs with one space

	for _, want := range coinList {
		got, ok := Coins[want.ID]
		assert.True(t, ok)
		assert.Equal(t, got.ID, want.ID)
		assert.Equal(t, got.Handle, want.Handle)
		assert.Equal(t, got.Symbol, want.Symbol)
		assert.Equal(t, got.Name, want.Name)
		assert.Equal(t, got.Decimals, want.Decimals)
		assert.Equal(t, got.BlockTime, want.BlockTime)
		assert.Equal(t, got.MinConfirmations, want.MinConfirmations)
		assert.Equal(t, got.IsTokenSupported, want.IsTokenSupported)

		s := cases.Title(language.English).String(want.Handle)
		method := fmt.Sprintf("func %s() Coin", s)
		assert.True(t, strings.Contains(code, method), "Coin method not found")

		enum := fmt.Sprintf("%s = %d", strings.ToUpper(want.Handle), want.ID)
		assert.True(t, strings.Contains(code, enum), "Coin enum not found")

	}
}

func TestEthereum(t *testing.T) {

	c := Ethereum()

	assert.Equal(t, uint(60), c.ID)
	assert.Equal(t, "ethereum", c.Handle)
	assert.Equal(t, "ETH", c.Symbol)
	assert.Equal(t, "Ethereum", c.Name)
	assert.Equal(t, uint(18), c.Decimals)
	assert.Equal(t, 10000, c.BlockTime)
	assert.Equal(t, int64(12), c.MinConfirmations)
	assert.True(t, c.IsTokenSupported)
}

func TestBinance(t *testing.T) {

	c := Smartchain()

	assert.Equal(t, uint(20000714), c.ID)
	assert.Equal(t, "smartchain", c.Handle)
	assert.Equal(t, "BNB", c.Symbol)
	assert.Equal(t, "Smart Chain", c.Name)
	assert.Equal(t, uint(18), c.Decimals)
	assert.Equal(t, 3000, c.BlockTime)
	assert.Equal(t, int64(12), c.MinConfirmations)
	assert.True(t, c.IsTokenSupported)
}

func TestCosmos(t *testing.T) {

	c := Cosmos()

	assert.Equal(t, uint(118), c.ID)
	assert.Equal(t, "cosmos", c.Handle)
	assert.Equal(t, "ATOM", c.Symbol)
	assert.Equal(t, "Cosmos", c.Name)
	assert.Equal(t, uint(6), c.Decimals)
	assert.Equal(t, 5000, c.BlockTime)
	assert.Equal(t, int64(7), c.MinConfirmations)
	assert.True(t, c.IsTokenSupported)
}

func TestPublicVariables(t *testing.T) {
	want := []Coin{
		{
			ID:               20000714,
			Handle:           "smartchain",
			Symbol:           "BNB",
			Name:             "Smart Chain",
			Decimals:         18,
			BlockTime:        3000,
			MinConfirmations: 12,
			Blockchain:       "Ethereum",
			IsTokenSupported: true,
		},
		{
			ID:               60,
			Handle:           "ethereum",
			Symbol:           "ETH",
			Name:             "Ethereum",
			Decimals:         18,
			BlockTime:        10000,
			MinConfirmations: 12,
			Blockchain:       "Ethereum",
			IsTokenSupported: true,
		},
	}

	for _, c := range want {
		coin := Coins[c.ID]
		assert.Equal(t, c, coin)

		chain := Chains[c.Handle]
		assert.Equal(t, c, chain)
	}
}
