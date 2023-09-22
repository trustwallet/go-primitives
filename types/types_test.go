package types

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHex_UnmarshalAndMarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		result  string
		isError bool
	}{
		{
			name:    "value greater than 2^64 -1",
			input:   []byte(`{"value":"0x850a9af493d065b6c"}`),
			result:  "153386322112866048876",
			isError: false,
		},
		{
			name:    "value less than 2^64 -1",
			input:   []byte(`{"value":"0x746a528800"}`),
			result:  "500000000000",
			isError: false,
		},
		{
			name:    "error case: not hex (string)",
			input:   []byte(`{"value":"error_value"}`),
			result:  "",
			isError: true,
		},
		{
			name:    "error case: not hex (octal)",
			input:   []byte(`{"value":"20502515364447501455554"}`),
			result:  "",
			isError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			type req struct {
				Value *HexNumber `json:"value"`
			}

			var v req

			err := json.Unmarshal(tc.input, &v)
			if tc.isError {
				return
			}
			assert.NoError(t, err)

			output := (*big.Int)(v.Value)
			assert.Equal(t, tc.result, output.String())

			bytes, err := json.Marshal(v)
			assert.NoError(t, err)
			assert.Equal(t, tc.input, bytes)
		})
	}
}
