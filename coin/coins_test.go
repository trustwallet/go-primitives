package coin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCoinsBlockchain checks if all chains from coins.yml have blockchain field defined
func TestCoinsBlockchain(t *testing.T) {
	for _, c := range Coins {
		assert.Falsef(t, c.Blockchain == "", fmt.Sprintf("chain: %s", c.Handle))
	}
}
