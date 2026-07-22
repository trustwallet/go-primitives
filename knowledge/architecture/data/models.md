---
category: architecture
subcategory: data
confidence: low
documentType: explanation
scope: repo
contentHash: d4c6d7e1e207
tags: [architecture, domain]
source: architecture/data/models.md
verified: 2026-07-22
splitPartIndex: 1
splitPartTotal: 2
canonical: true
---

## Data Models & Schemas
<!-- sdd-knowledge-generated -->

> Field-level shape of data models extracted via tree-sitter: TS interfaces / type aliases, Zod `z.object` schemas, Go/Rust/Swift structs, Kotlin data classes, and Python dataclasses. Scoped to domain data — UI views/props, view-models, design tokens (theme/style/colors), and constant/identifier namespaces are excluded. Deterministic, no LLM.

## Asset

_struct · `types/asset.go`:3_

| Field | Type | Optional |
|-------|------|----------|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Symbol` | `string` | no |
| `Type` | `TokenType` | no |
| `Decimals` | `uint` | no |

## Batch

_struct · `slice/batch.go`:8_

| Field | Type | Optional |
|-------|------|----------|
| `values` | `[]T` | no |

## Block

_struct · `types/tx.go`:57_

| Field | Type | Optional |
|-------|------|----------|
| `Number` | `int64` | no |
| `Txs` | `[]Tx` | no |

## Coin

_struct · `coin/coins.go`:17_

| Field | Type | Optional |
|-------|------|----------|
| `ID` | `uint` | no |
| `Handle` | `string` | no |
| `Symbol` | `string` | no |
| `Name` | `string` | no |
| `Decimals` | `uint` | no |
| `BlockTime` | `int` | no |
| `MinConfirmations` | `int64` | no |
| `Blockchain` | `string` | no |
| `ChainID` | `*uint` | no |

## Coin

_struct · `coin/gen.go`:125_

| Field | Type | Optional |
|-------|------|----------|
| `ID` | `uint` | no |
| `Handle` | `string` | no |
| `Symbol` | `string` | no |
| `Name` | `string` | no |
| `Decimals` | `uint` | no |
| `BlockTime` | `int` | no |
| `MinConfirmations` | `int64` | no |
| `Blockchain` | `string` | no |
| `Deprecated` | `bool` | no |
| `ChainID` | `*uint` | no |

## Collectible

_struct · `types/collectibles.go`:39_

| Field | Type | Optional |
|-------|------|----------|
| `ID` | `string` | no |
| `CollectionID` | `string` | no |
| `TokenID` | `string` | no |
| `ContractAddress` | `string` | no |
| `Category` | `string` | no |
| `ImageUrl` | `string` | no |
| `ExternalLink` | `string` | no |
| `ProviderLink` | `string` | no |
| `Type` | `string` | no |
| `Description` | `string` | no |
| `Coin` | `uint` | no |
| `Name` | `string` | no |
| `Version` | `string` | no |
| `TransferFee` | `*CollectibleTransferFee` | no |
| `PreviewImageURL` | `*CollectibleMedia` | no |
| `OriginalSourceURL` | `CollectibleMedia` | no |
| `Properties` | `[]CollectibleProperty` | no |
| `About` | `string` | no |
| `Balance` | `string` | no |

## CollectibleMedia

_struct · `types/collectibles.go`:34_

| Field | Type | Optional |
|-------|------|----------|
| `Mimetype` | `string` | no |
| `URL` | `string` | no |

## CollectibleProperty

_struct · `types/collectibles.go`:61_

| Field | Type | Optional |
|-------|------|----------|
| `Key` | `string` | no |
| `Value` | `string` | no |
| `Rarity` | `float32` | no |

## CollectibleTransferFee

_struct · `types/collectibles.go`:29_

| Field | Type | Optional |
|-------|------|----------|
| `Asset` | `string` | no |
| `Amount` | `string` | no |

## Collection

_struct · `types/collectibles.go`:13_

| Field | Type | Optional |
|-------|------|----------|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `ImageUrl` | `string` | no |
| `Description` | `string` | no |
| `ExternalLink` | `string` | no |
| `Total` | `int` | no |
| `Address` | `string` | no |
| `Coin` | `uint` | no |
| `Type` | `string` | no |
| `PreviewImageURL` | `*CollectibleMedia` | no |
| `OriginalSourceURL` | `CollectibleMedia` | no |

## ContractCall

_struct · `types/tx.go`:150_

| Field | Type | Optional |
|-------|------|----------|
| `Asset` | `coin.AssetID` | no |
| `Value` | `Amount` | no |
| `Input` | `string` | no |
