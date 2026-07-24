---
title: "Numbers Package — Decimal and Satoshi Conversions"
category: libs
tags: [numbers, decimal, satoshi, big-int, precision]
confidence: high
source: "source analysis: numbers/amount.go, numbers/decimal.go, numbers/number.go"
updated: "2026-07-22"
---

# Numbers Package

## Purpose

`package numbers` provides string-based arithmetic helpers for blockchain amounts.
Where possible it avoids floating-point arithmetic — the `decimal.go` / `number.go`
helpers operate on string representations using `math/big` or `shopspring/decimal`
for precision. Note that this is **not** universal: `ParseAmount` falls back to
`ToSatoshi`, which uses `strconv.ParseFloat` + `float64` (× 10^8), so parsing a
non-integer amount string does go through float64 and is subject to its precision
limits.

## Three files, three concerns

### `amount.go` — Amount parsing

| Function | Input | Output | Notes |
|----------|-------|--------|-------|
| `ParseAmount(amount string) int64` | decimal or float string | `int64` smallest-unit | Tries `ParseInt` first; falls back to `ToSatoshi` (float×10^8). Central symbol. |
| `ToSatoshi(amount string) int64` | float string | `int64` | `parseFloat * 1e8`. Use only for Bitcoin-scale values (8 decimal places). |
| `AddAmount(left, right string) string` | two amount strings | sum as string (no error) | Parses both via `ParseAmount`, adds as `int64`, formats back to string |
| `GetAmountValue(amount string) string` | amount string | `string` | Delegates to `ParseAmount`, then `strconv.FormatInt` |

### `decimal.go` — String-based decimal operations

| Function | Purpose |
|----------|---------|
| `DecimalToSatoshis(dec string) (string, error)` | Removes the decimal point: `"12.345"` → `"12345"`. Returns error if more than one `.` |
| `DecimalExp(dec string, exp int) string` | Shifts decimal point by `exp` places (right = multiply, left = divide). Pure string manipulation — no float involved. |
| `HexToDecimal(hex string) (string, error)` | Converts `"0x1a"` → `"26"` using `math/big` |
| `CutZeroFractional(dec string) (string, bool)` | Strips `.000...` only if fractional part is all zeros. Returns `(result, wasStripped bool)`. |

### `number.go` — General numeric utilities

| Function | Purpose |
|----------|---------|
| `ToDecimal(value string, exp int) string` | Converts smallest-unit integer string to decimal string using `shopspring/decimal`. E.g. `("1000000", 6)` → `"1.000000"` |
| `FromDecimal(dec string) string` | Removes decimal point (delegates to `DecimalToSatoshis`, ignores error) |
| `FromDecimalExp(dec string, exp int) string` | Shifts decimal point left by `exp` (delegates to `DecimalExp`) |
| `Min(a, b int) int`, `Max` | Integer min/max |
| `Round(num float64) int` | Rounds float to nearest int (`math.Floor(num + 0.5)`) |
| `Float64toPrecision(num float64, precision int) float64` | Rounds to N decimal places using `Round` |
| `SliceAtoi(sa []string) ([]int, error)` | Converts `[]string` of integers to `[]int` |

## Key invariant

`Amount string` (in `types/tx.go`) is always in the **smallest unit** (Wei, Satoshi,
uATOM, etc.). Use `numbers.ToDecimal(amount, coin.Decimals)` to convert to a
human-readable decimal string for display.

## See Also
- [features/numbers.md](../features/numbers.md)
- [architecture/transaction-model.md](../architecture/transaction-model.md)
- [code-conventions/go-conventions.md](../code-conventions/go-conventions.md)
- [god nodes](../architecture/god-nodes.md) <!-- rel:strong -->
- [domain overview](../architecture/domain-overview.md) <!-- rel:strong -->
