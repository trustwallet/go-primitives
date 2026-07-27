---
category: architecture
subcategory: data
confidence: low
documentType: explanation
scope: repo
contentHash: e4d43195cce5
tags: [architecture]
source: architecture/data/models.md
verified: 2026-07-22
splitPartIndex: 2
splitPartTotal: 2
canonical: false
synthetic: split-part
---

## Data Models & Schemas (Part 2)

_struct · `types/tx.go`:116_

| Field | Type | Optional |
|-------|------|----------|
| `Asset` | `coin.AssetID` | no |
| `Value` | `Amount` | no |

## Subscription

_struct · `types/subscription.go`:17_

| Field | Type | Optional |
|-------|------|----------|
| `Coin` | `uint` | no |
| `Address` | `string` | no |

## SubscriptionEvent

_struct · `types/subscription.go`:12_

| Field | Type | Optional |
|-------|------|----------|
| `Subscriptions` | `Subscriptions` | no |
| `Operation` | `SubscriptionOperation` | no |

## Swap

_struct · `types/tx.go`:144_

| Field | Type | Optional |
|-------|------|----------|
| `From` | `Transfer` | no |
| `To` | `Transfer` | no |

## Token

_struct · `types/token.go`:22_

| Field | Type | Optional |
|-------|------|----------|
| `Name` | `string` | no |
| `Symbol` | `string` | no |
| `Decimals` | `uint` | no |
| `TokenID` | `string` | no |
| `Coin` | `uint` | no |
| `Type` | `TokenType` | no |

## TransactionNotification

_struct · `types/subscription.go`:22_

| Field | Type | Optional |
|-------|------|----------|
| `Action` | `TransactionType` | no |
| `Result` | `Tx` | no |

## Transfer

_struct · `types/tx.go`:130_

| Field | Type | Optional |
|-------|------|----------|
| `Asset` | `coin.AssetID` | no |
| `Value` | `Amount` | no |

## TransferNFT

_struct · `types/tx.go`:136_

| Field | Type | Optional |
|-------|------|----------|
| `Asset` | `coin.AssetID` | no |
| `Collection` | `string` | no |
| `CollectibleID` | `string` | no |
| `CollectionSymbol` | `string` | no |
| `Value` | `Amount` | no |

## Tx

_struct · `types/tx.go`:68_

| Field | Type | Optional |
|-------|------|----------|
| `ID` | `string` | no |
| `From` | `string` | no |
| `To` | `string` | no |
| `BlockCreatedAt` | `int64` | no |
| `Block` | `uint64` | no |
| `Status` | `Status` | no |
| `Error` | `string` | no |
| `Sequence` | `uint64` | no |
| `Type` | `TransactionType` | no |
| `Direction` | `Direction` | no |
| `Inputs` | `[]TxOutput` | no |
| `Outputs` | `[]TxOutput` | no |
| `Tokens` | `[]Asset` | no |
| `Memo` | `string` | no |
| `Fee` | `Fee` | no |
| `Metadata` | `interface{}` | no |
| `CreatedAt` | `int64` | no |

## TxOutput

_struct · `types/tx.go`:123_

| Field | Type | Optional |
|-------|------|----------|
| `Address` | `string` | no |
| `Value` | `Amount` | no |
| `Asset` | `coin.AssetID` | no |

## TxPage

_struct · `types/tx.go`:62_

| Field | Type | Optional |
|-------|------|----------|
| `Total` | `int` | no |
| `Docs` | `[]Tx` | no |

## See Also
- [go conventions](../../code-conventions/go-conventions.md) <!-- rel:strong -->
- [testing strategy](../../tests/testing-strategy.md) <!-- rel:strong -->
- [numbers](../../libs/numbers.md) <!-- rel:strong -->
- [overview](../../build/overview.md) <!-- rel:related -->
- [slice](../../features/slice.md) <!-- rel:related -->
