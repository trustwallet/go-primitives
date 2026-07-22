---
title: "Build and CI Overview"
category: build
tags: [build, ci, make, code-generation, release]
confidence: high
source: "source analysis: Makefile, .github/workflows/go.yml, .github/workflows/release.yml, .github/dependabot.yml"
updated: "2026-07-22"
---

# Build and CI Overview

## Makefile

```makefile
make generate-coins   # Regenerate coin/coins.go from coins.yml
```

This is the only non-trivial Makefile target. It runs:
```
go run -tags=coins coin/gen.go && goimports ./coin/...
```

`//go:build coins` guards the generator so it is never compiled into the library.

## CI: `.github/workflows/go.yml`

Triggers on push and pull_request to `master`.

Steps:
1. `actions/checkout`
2. `actions/setup-go` (uses `go-version-file: go.mod`)
3. `go get ./...`
4. `go test -v ./...`
5. `golangci-lint-action@v6` (version: v1.63, `only-new-issues: true`)

**Note:** The org-wide `golangci-lint-action` is proxied through
`trustwallet/github-actions-proxy` — see the github-actions-proxy knowledge base for
the replacement action path.

## CI: `.github/workflows/release.yml`

Triggers on push to `master` (separately from go.yml). Auto-bumps the semver tag using
`mathieudutour/github-tag-action@v6.2`. No build artifacts are produced — releases are
just tags on the module (consumers use `go get github.com/trustwallet/go-primitives@vX.Y.Z`).

## Dependabot

Weekly Go module updates and monthly GitHub Actions updates, both grouped into single PRs.

## Go module

```
module github.com/trustwallet/go-primitives
go 1.19
```

External dependencies:
| Dep | Version | Used for |
|-----|---------|---------|
| `github.com/deckarep/golang-set` | v1.7.1 | Set ops for UTXO direction inference |
| `github.com/shopspring/decimal` | v1.2.0 | Precise decimal arithmetic in `numbers.ToDecimal` |
| `github.com/stretchr/testify` | v1.7.0 | Test assertions |
| `golang.org/x/crypto` | v0.1.0 | Keccak-256 for EIP-55 |
| `gopkg.in/yaml.v2` | v2.4.0 | YAML parsing in code generator |

## See Also
- [architecture/coin-registry.md](../architecture/coin-registry.md) — code generation details
- [tests/testing-strategy.md](../tests/testing-strategy.md)
- [domain overview](../architecture/domain-overview.md) <!-- rel:strong -->
- [numbers](../libs/numbers.md) <!-- rel:related -->
- [numbers](../features/numbers.md) <!-- rel:related -->
