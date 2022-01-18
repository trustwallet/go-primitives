//go:build coins
// +build coins

package main

import (
	"html/template"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	coinFile     = "coin/coins.yml"
	filename     = "coin/coins.go"
	templateFile = `// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
// using data from coins.yml
package coin

import (
	"fmt"
)

const (
	coinPrefix  = "c"
	tokenPrefix = "t"
)

// Coin is the native currency of a blockchain
type Coin struct {
	ID               uint
	Handle           string
	Symbol           string
	Name             string
	Decimals         uint
	BlockTime        int
	MinConfirmations int64
}

type AssetID string

func (c *Coin) String() string {
	return fmt.Sprintf("[%s] %s (#%d)", c.Symbol, c.Name, c.ID)
}

func (c Coin) AssetID() AssetID {
	return AssetID(coinPrefix + fmt.Sprint(c.ID))
}

func (c Coin) TokenAssetID(t string) AssetID {
	result := c.AssetID()
	if len(t) > 0 {
		result += AssetID("_" + tokenPrefix + t)
	}

	return result
}

const (
{{- range .Coins }}
	{{ .Handle | ToUpper }} = {{ .ID }}
{{- end }}
)

var Coins = map[uint]Coin{
{{- range .Coins }}
	{{ .Handle | ToUpper }}: {
		ID:               {{.ID}},
		Handle:           "{{.Handle}}",
		Symbol:           "{{.Symbol}}",
		Name:             "{{.Name}}",
		Decimals:         {{.Decimals}},
		BlockTime:        {{.BlockTime}},
		MinConfirmations: {{.MinConfirmations}},
	},
{{- end }}
}

var Chains = map[string]Coin{
{{- range .Coins }}
	{{ .Handle | Capitalize }}().Handle: {
		ID:               {{.ID}},
		Handle:           "{{.Handle}}",
		Symbol:           "{{.Symbol}}",
		Name:             "{{.Name}}",
		Decimals:         {{.Decimals}},
		BlockTime:        {{.BlockTime}},
		MinConfirmations: {{.MinConfirmations}},
	},
{{- end }}
}

{{- range .Coins }}

func {{ .Handle | Capitalize }}() Coin {
	return Coins[{{ .Handle | ToUpper }}]
}
{{- end }}

`
)

type Coin struct {
	ID               uint   `yaml:"id"`
	Handle           string `yaml:"handle"`
	Symbol           string `yaml:"symbol"`
	Name             string `yaml:"name"`
	Decimals         uint   `yaml:"decimals"`
	BlockTime        int    `yaml:"blockTime"`
	MinConfirmations int64  `yaml:"minConfirmations"`
}

func main() {
	var coinList []Coin
	coin, err := os.Open(coinFile)
	dec := yaml.NewDecoder(coin)
	err = dec.Decode(&coinList)
	if err != nil {
		log.Panic(err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	funcMap := template.FuncMap{
		"Capitalize": strings.Title,
		"ToUpper":    strings.ToUpper,
	}

	coinsTemplate := template.Must(template.New("").Funcs(funcMap).Parse(templateFile))
	err = coinsTemplate.Execute(f, map[string]interface{}{
		"Timestamp": time.Now(),
		"Coins":     coinList,
	})
	if err != nil {
		log.Panic(err)
	}
}

func getValidParameter(env, variable string) string {
	e, ok := os.LookupEnv(env)
	if ok {
		return e
	}
	return variable
}
