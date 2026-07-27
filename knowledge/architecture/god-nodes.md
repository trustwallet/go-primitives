---
title: "Central Symbols — God-Node Analysis"
category: architecture
tags: [god-nodes, central-symbols, contracts]
confidence: high
source: "knowledge/.relations.json, source analysis: asset/id.go, numbers/amount.go, address/address.go, types/token.go, slice/batch.go"
updated: "2026-07-22"
---

# Central Symbols — God-Node Analysis

The relation graph identifies symbols with the highest coupling. This doc explains
what each one does, its invariants, and why it is central.

## `ParseID` (asset/id.go) — fan-in 1, fan-out 2

**Contract:** Parses an asset ID string in the form `c<coinID>` or `c<coinID>_t<tokenID>`
and returns `(coinID uint, tokenID string, error)`.

- Input `"c60"` → `(60, "", nil)` (Ethereum coin)
- Input `"c60_t0x123..."` → `(60, "0x123...", nil)` (ERC-20 token on Ethereum)
- Bad input → `(0, "", ErrBadAssetID)`

**Why central:** Every place that receives an asset ID string (from JSON, from the mobile
app) must go through this to get structured data. It delegates to `FindCoinID` and
`FindTokenID`.

**Invariants:**
- `coinPrefix = 'c'`, `tokenPrefix = 't'` — hard-coded single-char delimiters
- Split on `_` (first `_` only — `SplitN(..., 2)`) to allow token IDs containing `_`
- Returns `ErrBadAssetID` (not raw errors) so callers can distinguish bad input from
  unexpected errors

**Callers:** `GetImageURL` (asset/image.go), `GetSubscriptionAddresses` (types/tx.go),
`IsEVM` (types/tx.go)

---

## `ParseAmount` (numbers/amount.go) — fan-in 2, fan-out 1

**Contract:** Parses a smallest-unit amount string to `int64`. First tries
`strconv.ParseInt`; on failure falls back to `ToSatoshi` (floating-point × 10^8),
which accommodates decimal input like `"1.5"` meaning 150_000_000 satoshis.

**Why central:** Both `AddAmount` and `GetAmountValue` depend on this to get a numeric
amount from a string field. The dual-path parsing (int first, satoshi fallback) is the
key semantic: callers do not need to know which representation was stored.

**Invariants:**
- Only the fallback path (`ToSatoshi`) can produce loss — `ParseInt` is exact
- Negative returns are possible when the int parsing path handles a negative integer

**Callers:** `AddAmount`, `GetAmountValue` (both in numbers/amount.go)

---

## `EIP55Checksum` (address/address.go) — fan-in 1, fan-out 1

**Contract:** Takes an Ethereum address string (with or without `0x` prefix),
strips the prefix via `Remove0x`, Keccak-256 hashes the lowercase hex, and returns
the mixed-case EIP-55 checksum address.

**Why central:** All EVM address normalization flows through this. Downstream
display and comparison both depend on the canonical checksum form.

**Callers:** `ToEIP55ByCoinID` (address/address.go)

---

## `removeFirstChar` (asset/id.go) — fan-in 2, fan-out 0

An internal helper called by both `FindCoinID` and `FindTokenID` to strip the
`c`/`t` prefix from the relevant segment. Uses `rune` slicing (not byte slicing) to
be Unicode-safe for any future prefix changes, though the current prefixes are ASCII.

Not exported — this is an implementation detail, not a public API.

**Dead-code note:** Listed as an entry-point reachable symbol; NOT dead code despite
zero exports.

---

## `GetChunks` (slice/batch.go) — fan-in 0, fan-out 2

**Contract:** Splits any slice into fixed-size chunks. Two implementations:

1. **Generic** (`Batch[T].GetChunks(size int) [][]T`) — type-safe, Go 1.18+
2. **Reflection-based** (`GetChunks(slice interface{}, size uint) ...`) — for
   pre-generics callers; delegates to `GetInterfaceSlice` + `GetInterfaceSliceBatch`

**Why central:** Batch processing of blockchain data (paginated RPC calls, DB writes)
requires chunking. This utility is the single place for that pattern.

---

## `ParseTokenTypeFromString` (types/token.go) — fan-in 1, fan-out 1

**Contract:** Takes a raw token type string (e.g. `"ERC20"`) and returns the
canonical `TokenType` value, or `ErrUnknownTokenType` if not in the registry.
Delegates to `GetTokenTypes()` to enumerate the known set.

**Called by:** `GetTokenVersion` — which needs a typed `TokenType` before mapping it
to the versioned schema integer (`TokenVersion`).

## See Also
- [architecture/domain-overview.md](domain-overview.md)
- [features/asset.md](../features/asset.md)
- [architecture/call-graph.md](call-graph.md)
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:strong -->
- [numbers](../libs/numbers.md) <!-- rel:strong -->
