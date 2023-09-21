package types

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHex_UnmarshalAndMarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		result string
	}{
		{
			input:  []byte(`{"value":"0x850a9af493d065b6c"}`),
			result: "153386322112866048876",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			type req struct {
				Value *HexNumber `json:"value"`
			}

			var v req

			err := json.Unmarshal(tc.input, &v)
			assert.NoError(t, err)

			output := (*big.Int)(v.Value)
			assert.Equal(t, tc.result, output.String())

			bytes, err := json.Marshal(v)
			assert.NoError(t, err)
			assert.Equal(t, tc.input, bytes)
		})
	}
}
