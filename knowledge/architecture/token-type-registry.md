---
title: "Token Type Registry — Blockchain Token Standards"
category: architecture
tags: [token-type, token-version, erc20, bep20, registry, versioning]
confidence: high
source: "source analysis: types/token.go, types/chain.go"
updated: "2026-07-22"
---

# Token Type Registry

## What it is

`types/token.go` is the canonical registry of all supported on-chain token standards.
It defines ~80 `TokenType` string constants (one per blockchain/standard combination)
and maps them to:

1. **Parent chain** — via `GetChainFromAssetType` in `types/chain.go`
2. **Version integer** — via `GetTokenVersion` (governs the mobile app schema format)
3. **EVM coin index** — via `GetEthereumTokenTypeByIndex` (for EVM chains that use
   different token standards per chain)

## `TokenType` constants (selected)

| Constant | Value | Blockchain |
|----------|-------|-----------|
| `ERC20` | `"ERC20"` | Ethereum mainnet |
| `BEP20` | `"BEP20"` | BNB Smart Chain |
| `SPL` | `"SPL"` | Solana |
| `TRC20` | `"TRC20"` | TRON |
| `JETTON` | `"JETTON"` | TON |
| `APTOSFA` | `"APTOSFA"` | Aptos fungible assets |
| `CW20` | `"CW20"` | CosmWasm chains (Terra, Osmosis…) |
| `ESDT` | `"ESDT"` | MultiversX (Elrond) |
| `POLYGON` | `"POLYGON"` | Polygon |
| `ARBITRUM` | `"ARBITRUM"` | Arbitrum One |

Full list: see `types/token.go` lines 36–150.

## `GetTokenType(coinID uint, tokenID string) (string, bool)`

Returns the dominant token type for a coin. Most coins have one token type; exceptions:

- **TRON (coin ID 195)**: `TRC10` if tokenID is numeric, `TRC20` otherwise
- **TERRA (coin ID 330)**: `CW20` if tokenID starts with `"terra"`, `TERRA` otherwise
- **APTOS (coin ID 637)**: `APTOSFA` if tokenID contains `"::"`, `APTOS` otherwise

Returns `("", false)` for coin IDs that have no associated token standard.

## `GetTokenVersion(tokenType string) (TokenVersion, error)`

Returns a monotonically increasing version integer assigned when each token type was
introduced. The version integers (V0–V28) are part of the mobile ↔ backend wire
contract. Consumers of `Token` must include the correct version for the mobile app to
render the token correctly.

**Never reuse a version integer**; always assign the next increment when adding a new
`TokenType`.

## `GetChainFromAssetType(assetType string) (coin.Coin, error)`

Maps a `TokenType` string back to the parent `coin.Coin`. It is a large `switch`
statement covering all ~80 token types. Used by services that receive a token type
and need the parent chain's metadata (e.g. RPC endpoint selection).

## See Also
- [architecture/coin-registry.md](coin-registry.md)
- [architecture/asset-id-encoding.md](asset-id-encoding.md)
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:strong -->
- [testing strategy](../tests/testing-strategy.md) <!-- rel:related -->
- [numbers](../libs/numbers.md) <!-- rel:weak -->
