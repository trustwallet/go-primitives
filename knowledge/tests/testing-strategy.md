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

**Table-driven tests** are the dominant pattern. Each test function defines a `tests`
slice of structs with `name`, `input`, and `expected` fields, then iterates with
`t.Run(tt.name, ...)`.

Example shape (from `asset/id_test.go`, `numbers/decimal_test.go`):
```go
tests := []struct {
    name     string
    input    string
    expected string
}{
    {"ethereum coin", "c60", "60 ''"},
    ...
}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        ...
        assert.Equal(t, tt.expected, result)
    })
}
```

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
