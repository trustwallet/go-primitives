---
title: "Go Code Conventions"
category: code-conventions
tags: [go, conventions, naming, errors, json]
confidence: high
source: "source analysis: all Go source files"
updated: "2026-07-22"
---

# Go Code Conventions

## Naming

- **Types** are `PascalCase` (`Tx`, `TokenType`, `HexNumber`).
- **Constants** for string enums use `ALLCAPS` for token types (`ERC20`, `BEP20`) and
  `PascalCase` for other string constants (`StatusCompleted`, `DirectionIncoming`).
- **Functions and methods** are `PascalCase` for exported, `camelCase` for unexported
  (`cleanMemo`, `removeFirstChar`, `determineTransactionDirection`).
- **Test files** follow `<file>_test.go` in the same package.

## JSON field names

All struct fields use `json:"snake_case"` tags. Optional fields use `json:"field,omitempty"`.
The `Tx` struct mixes required and optional fields — see `types/tx.go` for the exact tags.

## Error handling

- Sentinel errors are defined as `var Err... = errors.New(...)` at package level.
  Known sentinels: `asset.ErrBadAssetID`, `types.ErrUnknownTokenType`,
  `types.errTokenVersionNotImplemented` (unexported — internal).
- Functions return `(result, error)` per Go convention.
- `GetChainFromAssetType` returns `(coin.Coin, error)` with a formatted error for
  unknown types — downstream callers should handle this rather than ignoring it.

## Amount type

`Amount string` (in `types/tx.go`) is a **positive decimal integer string** in the
smallest unit (Wei, Satoshis, etc.). It is NEVER a floating-point representation.
The `numbers` package provides `ToDecimal` and `FromDecimal` to convert for display.

## HexNumber type

`HexNumber` (in `types/types.go`) wraps `*big.Int` and marshals/unmarshals as a
`"0x..."` hex string in JSON. Use it wherever large integer values must round-trip
through JSON as hex (e.g. gas prices, chain values).

## TokenType versioning

`GetTokenVersion(tokenType string) (TokenVersion, error)` maps a `TokenType` to a
versioned schema integer (`V0`–`V28`). These version integers are part of the mobile
app ↔ backend contract — changing them is a breaking change.

## TODO in GetTokenType (types/token.go:352)

A TODO comment `// TODO: improve this` exists inside `GetTokenType`. This function
disambiguates TRON, TERRA, and APTOS token types based on both the coin ID and the
token ID format. The comment indicates the branching logic is considered rough but
functional.

## See Also
- [code-conventions/code-style/anti-patterns-failed-approaches.md](code-style/anti-patterns-failed-approaches.md)
- [architecture/transaction-model.md](../architecture/transaction-model.md)
- [tests/testing-strategy.md](../tests/testing-strategy.md)
- [token type registry](../architecture/token-type-registry.md) <!-- rel:strong -->
- [god nodes](../architecture/god-nodes.md) <!-- rel:strong -->
