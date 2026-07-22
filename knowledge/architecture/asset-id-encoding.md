---
title: "Asset ID Encoding — AssetID vs AssetId"
category: architecture
tags: [asset-id, encoding, coin, token, naming]
confidence: high
source: "source analysis: asset/id.go, coin/coins.go, types/token.go, types/tx.go"
updated: "2026-07-22"
---

# Asset ID Encoding

## The `AssetID` / `AssetId` ambiguity — RESOLVED

The relation graph flagged `AssetID` and `AssetId` as possible duplicates. They are
**the same concept but two distinct types/contexts**:

| Name | Package | Kind | Role |
|------|---------|------|------|
| `coin.AssetID` | `coin/coins.go` | `type AssetID string` | The *type* for a formatted asset ID string (e.g. `"c60_t0xabc"`). Used as field types in `types/tx.go` (`Fee.Asset`, `TxOutput.Asset`, `Transfer.Asset`, etc.) |
| `Token.AssetId()` | `types/token.go` | method | Convenience method on `Token` that calls `asset.BuildID(t.Coin, t.TokenID)` and returns a `string`. |

They share the same encoding — `coin.AssetID` is just the typed alias for the string.
`Token.AssetId()` computes and returns a value of that same format.

**Recommendation:** When reading code, `coin.AssetID` (capital D, in `coin` pkg) is the
*type*, while `Token.AssetId()` (lowercase d, in `types` pkg) is a *method*. They will
not be confused in practice because Go's type system keeps them in separate namespaces.

## Encoding format

```
coin-only asset:   c<coinID>
                   e.g. "c60" — Ethereum (coin ID 60)

token asset:       c<coinID>_t<tokenID>
                   e.g. "c60_t0xdac17f958d2ee523a2206206994597c13d831ec7"
                        (Ethereum + USDT contract address)
```

- `c` and `t` are the fixed one-character prefixes (bytes, not runes — but
  `removeFirstChar` uses rune slicing for Unicode safety).
- `_` is the separator; a single `SplitN(..., 2)` is used so token IDs containing
  `_` are handled correctly.

## Key functions

| Function | Package | Purpose |
|----------|---------|---------|
| `asset.BuildID(coin uint, token string) string` | asset | Construct an asset ID string from components |
| `asset.ParseID(id string) (uint, string, error)` | asset | Decompose an asset ID string into components |
| `asset.FindCoinID(words []string) (uint, error)` | asset | Low-level: extract coin ID from split words |
| `asset.FindTokenID(words []string) string` | asset | Low-level: extract token ID from split words |
| `coin.AssetID.AssetID() coin.AssetID` | coin | Returns self (identity, for interface compliance) |
| `coin.AssetID.TokenAssetID(tokenID string) coin.AssetID` | coin | Returns a token-form asset ID for this coin |
| `Token.AssetId() string` | types | Computes `asset.BuildID(t.Coin, t.TokenID)` |

## See Also
- [features/asset.md](../features/asset.md)
- [architecture/god-nodes.md](god-nodes.md) — `ParseID` explained
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:strong -->
- [testing strategy](../tests/testing-strategy.md) <!-- rel:strong -->
- [numbers](../libs/numbers.md) <!-- rel:related -->
