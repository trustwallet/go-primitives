---
title: "Domain Overview — go-primitives"
category: architecture
tags: [domains, overview, blockchain, primitives]
confidence: high
source: "source analysis: README.md, go.mod, all package dirs"
updated: "2026-07-22"
---

# Domain Overview — go-primitives

`go-primitives` is a pure Go library (`github.com/trustwallet/go-primitives`) that provides
two orthogonal sets of primitives for Trust Wallet backends:

1. **Blockchain primitives** — canonical types and helpers tied to specific blockchains
   (`coin`, `types`, `address`, `asset`)
2. **General-purpose Go utilities** — chain-agnostic helpers (`numbers`, `slice`)

There are no entry points, no HTTP servers, no CLI programs. The module ships as a library
only; all packages are imported by downstream services.

## Package map

| Package | Domain | Responsibility |
|---------|--------|----------------|
| `coin` | Blockchain identities | Registry of 128+ supported blockchains: IDs, handles, symbols, decimals, explorer URLs. Code-generated from `coins.yml`. |
| `types` | Transaction & token model | The canonical `Tx`, `Token`, `Collection`, `Subscription`, `HexNumber` types used across TW backends. |
| `asset` | Asset ID encoding | Encodes/decodes asset IDs in the `c<coinID>_t<tokenID>` format shared between the mobile app and backend services. |
| `address` | EVM address utilities | EIP-55 checksum, 0x prefix helpers, Ronin address handling. |
| `numbers` | Numeric conversions | Decimal ↔ satoshi conversions, big-integer decimal formatting, hex-to-decimal. |
| `slice` | Batch processing | Generic and reflection-based slice chunking. |

## Cross-package dependencies

```
types  ──imports──▶  asset (BuildID, ParseID)
types  ──imports──▶  coin  (AssetID, Coin accessors)
asset  ──(none to other packages)
coin   ──(none to other packages)
numbers──(none to other packages)
slice  ──(none to other packages)
address──(none to other packages)
```

`types` is the only package that depends on the others. All other packages are
independent of each other.

## Domain sizing

| Domain | Go files | Symbols |
|--------|----------|---------|
| coin | 6 | 141 |
| types | 15 | 80 |
| numbers | 6 | 18 |
| asset | 4 | 7 |
| slice | 2 | 6 |
| address | 2 | 3 |

## See Also
- [architecture/god-nodes.md](god-nodes.md) — central symbols explained
- [features/coin.md](../features/coin.md)
- [architecture/dependency-graph.md](dependency-graph.md)
- [testing strategy](../tests/testing-strategy.md) <!-- rel:strong -->
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:strong -->
