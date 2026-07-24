---
title: "Transaction Model — Tx, Types, and Direction"
category: architecture
tags: [transaction, tx, direction, utxo, evm, types]
confidence: high
source: "source analysis: types/tx.go, types/marshal.go"
updated: "2026-07-22"
---

# Transaction Model

## Purpose

The `types` package defines the **canonical cross-chain transaction model** used by
Trust Wallet backend services. A single `Tx` struct represents any transaction on any
supported blockchain — UTXO (Bitcoin-style), account-based (Ethereum-style), or other.

## `Tx` struct

| Field | Type | Notes |
|-------|------|-------|
| `ID` | `string` | Transaction hash (chain-specific format) |
| `From` | `string` | Sender address |
| `To` | `string` | Recipient address |
| `Block` | `uint64` | Block height |
| `BlockCreatedAt` | `int64` | Unix timestamp of block inclusion |
| `Status` | `Status` | `"completed"` / `"pending"` / `"error"` |
| `Error` | `string` | Empty unless `Status == "error"` |
| `Sequence` | `uint64` | Nonce or sequence number |
| `Type` | `TransactionType` | Determines which `Metadata` sub-type is present |
| `Direction` | `Direction` | `"outgoing"` / `"incoming"` / `"yourself"` |
| `Inputs` | `[]TxOutput` | UTXO inputs (empty for account-based chains) |
| `Outputs` | `[]TxOutput` | UTXO outputs (empty for account-based chains) |
| `Tokens` | `[]Asset` | Token transfers (ERC-20 etc.) embedded in the tx |
| `Memo` | `string` | Optional memo (Cosmos, Stellar, etc.) |
| `Fee` | `Fee` | Transaction fee (asset ID + amount in smallest unit) |
| `Metadata` | `interface{}` | Type-discriminated payload: `Transfer`, `Swap`, `ContractCall`, `TransferNFT` |
| `CreatedAt` | `int64` | Database insertion timestamp |

## `TransactionType` and `Metadata` discriminated union

The `Type` field governs what is stored in `Metadata`:

| `Type` | `Metadata` struct | Contains |
|--------|-------------------|----------|
| `"transfer"` | `Transfer` | asset ID + amount |
| `"swap"` | `Swap` | from/to Transfer |
| `"contract_call"` | `ContractCall` | asset ID + amount + input data |
| `"transfer_nft"` | `TransferNFT` | asset ID + collection + collectible ID |
| staking types | `Transfer` | asset ID + staked amount |

`marshal.go` handles the JSON round-trip: `UnmarshalJSON` type-dispatches on `Type`
before deserializing `Metadata`, and `MarshalJSON` validates that `Type` matches the
actual `Metadata` type before serializing (defaults `Status` to `"completed"` if unset).

## Direction determination

`GetDirection(address string) Direction` computes whether the transaction is
incoming, outgoing, or a self-transfer relative to a given address.

If `t.Direction` is already set (non-empty) it is returned as-is. Otherwise the
branch is chosen structurally, **not** via `IsUTXO()`:

- **Input/output set present** (`len(t.Inputs) > 0 && len(t.Outputs) > 0`):
  `InferDirection(tx, addressSet)` using set arithmetic (via `golang-set`) on the
  input and output address sets. A tx where the address appears in both inputs and
  outputs is a self-send (`"yourself"`). This gate is independent of `IsUTXO()` —
  any tx carrying both inputs and outputs takes this path.
- **Otherwise** (account-based / no full in+out sets):
  `determineTransactionDirection(address, from, to)` — stake-undelegate and
  stake-claim-rewards are always `"incoming"`; else if `address == to`, `"yourself"`
  when `from == to`, otherwise `"incoming"`; else `"outgoing"`.

## UTXO vs account-based

`Tx.IsUTXO() bool` checks `t.Type == TxTransfer && len(t.Outputs) > 0` — not a coin
lookup. This means a UTXO check is structural (type + non-empty outputs), not
registry-based. Note it is **not** the same condition as the direction gate above,
which requires both inputs *and* outputs.

`Tx.IsEVM() (bool, error)` reads the asset ID from `t.Metadata` via the
`AssetHolder` interface (`t.Metadata.(AssetHolder).GetAsset()`), then
`asset.ParseID(...)` → `coin.IsEVM(coinID)`. It returns `(false, nil)` when
`Metadata` is nil or not an `AssetHolder`, and propagates the `ParseID` error.
It does **not** use `t.GetAssetID()`.

## `Txs` sort order

`Txs` implements `sort.Interface` via `Len`, `Less`, `Swap` in `marshal.go`.
`Less` sorts by `CreatedAt` **descending** (newest first).
`SortByBlockCreationTime()` sorts by `BlockCreatedAt` descending instead.

## Dead-code note

`determineTransactionDirection` (types/tx.go) appears as a dead-code candidate in
the relation graph because it is called via `GetDirection` with a method call that
the AST can't resolve at the type level. It is **not dead code** — it is actively
used inside `GetDirection`.

## See Also
- [architecture/god-nodes.md](god-nodes.md)
- [architecture/domain-overview.md](domain-overview.md)
- [go conventions](../code-conventions/go-conventions.md) <!-- rel:strong -->
- [numbers](../libs/numbers.md) <!-- rel:strong -->
- [testing strategy](../tests/testing-strategy.md) <!-- rel:strong -->
