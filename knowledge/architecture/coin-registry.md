---
title: "Coin Registry — Architecture and Code Generation"
category: architecture
tags: [coin, registry, code-generation, yaml, blockchain]
confidence: high
source: "source analysis: coin/coins.yml, coin/gen.go, coin/coins.go, coin/models.go"
updated: "2026-07-22"
---

# Coin Registry Architecture

## What it is

The `coin` package is the authoritative registry of all supported blockchains.
It is **code-generated from `coins.yml`** — a human-edited YAML file with 128+
coin definitions. The generated `coins.go` (~3368 lines) contains the full registry
as a `var Coins map[uint]Coin` lookup table plus one accessor function per coin
(e.g. `coin.Ethereum()`, `coin.Smartchain()`).

## How code generation works

```
coins.yml   → (go run -tags=coins coin/gen.go) → coins.go
```

- `coin/gen.go` is a program guarded by `//go:build coins` (never compiled into the
  library itself — only executed when you run `make generate-coins`).
- The Makefile target runs `goimports` after generation to keep the output tidy.
- **Never hand-edit `coins.go`** — changes will be overwritten by the next `make
  generate-coins` run.

## Adding a new coin

1. Add an entry to `coins.yml` with the required fields (`id`, `handle`, `symbol`,
   `name`, `decimals`, `blockTime`, `minConfirmations`, `blockchain`, optionally `chainID`).
2. Run `make generate-coins` to regenerate `coins.go`.
3. If it is an EVM chain with tokens, add a corresponding `TokenType` constant and a
   `case` in `GetTokenType()` inside `types/token.go`.
4. If it needs a block-explorer URL, add a `case` in `GetCoinExploreURL()` and
   `GetAddressExploreURL()` in `coin/models.go`.

## The `Coin` struct

| Field | Type | Meaning |
|-------|------|---------|
| `ID` | `uint` | Numeric coin ID (slip-044 derived; e.g. 60 = Ethereum) |
| `Handle` | `string` | Lowercase handle used in URLs / asset paths (e.g. `"ethereum"`) |
| `Symbol` | `string` | Ticker symbol (e.g. `"ETH"`) |
| `Name` | `string` | Display name (e.g. `"Ethereum"`) |
| `Decimals` | `uint` | Smallest unit decimals (e.g. 18 for ETH) |
| `BlockTime` | `int` | Milliseconds per block |
| `MinConfirmations` | `int64` | Confirmations needed to consider a tx finalized |
| `Blockchain` | `string` | Parent blockchain family (`"Ethereum"` for EVM chains, `"Bitcoin"`, etc.) |
| `ChainID` | `*uint` | EVM chain ID (nil for non-EVM chains) |
| `Deprecated` | `bool` | True for retired chains (only in `gen.go` template model) |

## IsEVM detection

`coin.IsEVM(coinID uint) bool` returns true when `Coins[coinID].Blockchain == "Ethereum"`.
This is the canonical way to check if a coin is an EVM chain — do not compare `ChainID`
or check `Blockchain` directly in calling code.

## Asset IDs from coin

`coin.AssetID` is a `string` type alias. Each `Coin` exposes:
- `AssetID() coin.AssetID` — returns the coin-only asset ID (`c<id>`)
- `TokenAssetID(tokenID string) coin.AssetID` — returns a token asset ID (`c<id>_t<tokenID>`)

## Key lookup functions

| Function | Lookup | Returns |
|----------|--------|---------|
| `coin.Coins[id]` | by numeric ID | `Coin` |
| `coin.Chains[handle]` | by handle string | `Coin` |
| `coin.GetCoinForId(id string)` | by handle string (linear scan over `coin.Coins`, matching `c.Handle == id`) | `(Coin, error)` — returns `errors.New("unknown id " + id)` when no match |
| `coin.Ethereum()` etc. | named accessor | `Coin` |

## See Also
- [features/coin.md](../features/coin.md)
- [architecture/domain-overview.md](domain-overview.md)
- [build/overview.md](../build/overview.md) — `make generate-coins` target
- [testing strategy](../tests/testing-strategy.md) <!-- rel:strong -->
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:strong -->
