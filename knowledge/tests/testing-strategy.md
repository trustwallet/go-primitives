---
title: "Testing Strategy"
category: tests
tags: [tests, go, unit-tests, table-driven]
confidence: high
source: "source analysis: *_test.go files across all packages"
updated: "2026-07-22"
---

# Testing Strategy

## Overview

All packages have accompanying `_test.go` files in the same package directory.
Tests are **unit tests only** — no integration tests, no mocks, no external dependencies.
The library has no network calls or I/O (except `coin/gen.go` which reads `coins.yml`
at codegen time, not at runtime).

## Test patterns

**Table-driven tests** are the dominant pattern. A test function defines a slice of
case structs (`name` plus per-case inputs and expected outputs), then iterates with
`t.Run(tt.name, ...)`.

Example shape (from `asset/id_test.go`, `TestParseID` — `ParseID` returns
`(uint, string, error)`):
```go
testStruct := []struct {
    name        string
    givenID     string
    wantedCoin  uint
    wantedToken string
    wantedError error
}{
    {name: "c714_tTWT-8C2", givenID: "c714_tTWT-8C2", wantedCoin: 714, wantedToken: "TWT-8C2", wantedError: nil},
    {name: "c714", givenID: "c714", wantedCoin: 714, wantedToken: "", wantedError: nil},
    // ...
}
for _, tt := range testStruct {
    t.Run(tt.name, func(t *testing.T) {
        coin, token, err := ParseID(tt.givenID)
        assert.Equal(t, tt.wantedCoin, coin)
        assert.Equal(t, tt.wantedToken, token)
        assert.Equal(t, tt.wantedError, err)
    })
}
```

Not every suite uses the struct-slice form: `numbers/decimal_test.go` instead defines
closure assertion helpers (e.g. `assertSatEquals(expected, input)`) and calls them with
a flat list of input/expected pairs — the same table-driven spirit without `t.Run`.

**Assertions** use `github.com/stretchr/testify/assert` throughout.

## Running tests

```bash
go test -v ./...
```

Tests are run by CI on every push/PR to `master` (see `build/overview.md`).

## What is tested

| Package | Coverage areas |
|---------|---------------|
| `coin` | `GetCoinForId`, `IsEVM`, `GetCoinExploreURL` for representative coins; codegen output correctness |
| `asset` | `ParseID` round-trips, `BuildID` format, `GetImageURL` URL construction |
| `address` | `Remove0x`, `EIP55Checksum` checksum correctness, `ToEIP55ByCoinID` EVM/Ronin |
| `numbers` | `DecimalToSatoshis`, `HexToDecimal`, `DecimalExp`, `ToDecimal`, `ParseAmount`, `ToSatoshi` |
| `slice` | `GetChunks` for various sizes and edge cases |
| `types` | `GetTokenType`, `GetTokenVersion`, `GetChainFromAssetType`, JSON marshal/unmarshal round-trips, subscription parsing |

## Dead-code in gen.go

`coin/gen.go` has a `main()` function and `getValidParameter()` that appear as
dead-code candidates because the build tag `//go:build coins` prevents them from
being compiled under normal `go test ./...`. They are not dead — they are the
code-generator entry point, invoked via `go run -tags=coins coin/gen.go`.

## See Also
- [build/overview.md](../build/overview.md)
- [code-conventions/go-conventions.md](../code-conventions/go-conventions.md)
- [call graph](../architecture/call-graph.md) <!-- rel:strong -->
- [god nodes](../architecture/god-nodes.md) <!-- rel:strong -->
- [domain overview](../architecture/domain-overview.md) <!-- rel:strong -->
