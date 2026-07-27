---
category: architecture
confidence: low
documentType: explanation
scope: repo
contentHash: e8b55b09a4ce
tags: [architecture, domain, dependency]
source: architecture/call-graph.md
verified: 2026-07-22
splitPartIndex: 2
splitPartTotal: 2
canonical: false
synthetic: split-part
---

## Call Graph (Part 2)

- Call edges: **27** across **248** symbols
- Reachable from entry points (exported symbols + routes): **244**
- Dependency cycles: **0**
- Cross-domain bridges: **0**
- Dead-code candidates (non-exported, zero callers, unreachable): **4**

### Most-coupled symbols (god-node ranking)

| Symbol | Fan-in | Fan-out | Degree |
|--------|--------|---------|--------|
| `ParseID` | 1 | 2 | 3 |
| `ParseAmount` | 2 | 1 | 3 |
| `EIP55Checksum` | 1 | 1 | 2 |
| `FindCoinID` | 1 | 1 | 2 |
| `FindTokenID` | 1 | 1 | 2 |
| `removeFirstChar` | 2 | 0 | 2 |
| `GetChunks` | 0 | 2 | 2 |
| `ParseTokenTypeFromString` | 1 | 1 | 2 |
| `Remove0x` | 1 | 0 | 1 |
| `ToEIP55ByCoinID` | 0 | 1 | 1 |
| `GetImageURL` | 0 | 1 | 1 |
| `AssetID` | 1 | 0 | 1 |
| `TokenAssetID` | 0 | 1 | 1 |
| `AddAmount` | 0 | 1 | 1 |
| `GetAmountValue` | 0 | 1 | 1 |

### Possible duplicate entities (name variants)

> Symbol names that normalize identically — likely the same entity spelled inconsistently. Unify or distinguish in the docs.

- `AssetID` / `AssetId`

### Dead-code candidates

> Non-exported symbols with no resolved callers, unreachable from any entry point. Static analysis cannot see dynamic dispatch — verify before removing.

- `ptr` (`coin/coins.go`)
- `getValidParameter` (`coin/gen.go`)
- `main` (`coin/gen.go`)
- `determineTransactionDirection` (`types/tx.go`)

## See Also
- [coin](../features/coin.md) <!-- rel:strong -->
- [testing strategy](../tests/testing-strategy.md) <!-- rel:related -->
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:weak -->
- [numbers](../libs/numbers.md) <!-- rel:weak -->
- [numbers](../features/numbers.md) <!-- rel:weak -->
