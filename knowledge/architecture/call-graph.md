# Call Graph

<!-- sdd-knowledge-generated -->

> Deterministic call graph extracted via tree-sitter (no LLM). Direct calls are resolved by lexical scope + imports; **dynamic dispatch and ambiguous name matches are withheld** rather than guessed. The `## Calls` table is `EXTRACTED` (a single resolved target). Member calls (`obj.method()`) whose method name resolves to exactly one definition repo-wide are recovered separately under `## Inferred calls` (`INFERRED` — receiver type unverified, but a single plausible target).

## Calls

| Caller | Callee | Example args | Sites |
|--------|--------|--------------|-------|
| `EIP55Checksum` | `Remove0x(input: string): string` | `(…)` | 1 |
| `ToEIP55ByCoinID` | `EIP55Checksum(unchecksummed: string): (string, error)` | `(str)` | 1 |
| `FindCoinID` | `removeFirstChar(input: string): string` | `(w)` | 1 |
| `FindTokenID` | `removeFirstChar(input: string): string` | `(w)` | 1 |
| `ParseID` | `FindCoinID(words: []string): (uint, error)` | `(rawResult)` | 1 |
| `ParseID` | `FindTokenID(words: []string): string` | `(rawResult)` | 1 |
| `GetImageURL` | `ParseID(id: string): (uint, string, error)` | `(asset)` | 1 |
| `TokenAssetID` | `AssetID` | `(…)` | 1 |
| `AddAmount` | `ParseAmount(amount: string): int64` | `(left)` | 2 |
| `GetAmountValue` | `ParseAmount(amount: string): int64` | `(amount)` | 1 |
| `ParseAmount` | `ToSatoshi(amount: string): int64` | `(amount)` | 1 |
| `Float64toPrecision` | `Round(num: float64): int` | `(…)` | 1 |
| `FromDecimal` | `DecimalToSatoshis(dec: string): (string, error)` | `(dec)` | 1 |
| `FromDecimalExp` | `DecimalExp(dec: string, exp: int): string` | `(dec, exp)` | 1 |
| `GetChunks` | `GetInterfaceSlice(slice: interface{}): ([]interface{}, error)` | `(slice)` | 1 |
| `GetChunks` | `GetInterfaceSliceBatch(values: []interface{}, sizeUint: uint): (chunks [][]interface{})` | `(interfaceSlice, size)` | 1 |
| `GetChainFromAssetType` | `TokenType` | `(assetType)` | 1 |
| `MarshalJSON` | `wrappedTx` | `(t)` | 1 |
| `UnmarshalJSON` | `Tx` | `(wrapped)` | 1 |
| `AddressID` | `GetAddressID(coin: string): string` | `(…, v.Address)` | 1 |
| `GetTokenType` | `GetEthereumTokenTypeByIndex(coinIndex: uint): (TokenType, error)` | `(c)` | 1 |
| `GetTokenVersion` | `ParseTokenTypeFromString(t: string): (TokenType, error)` | `(tokenType)` | 1 |
| `ParseTokenTypeFromString` | `GetTokenTypes(): []TokenType` | `()` | 1 |
| `CleanMemos` | `cleanMemo(memo: string): string` | `(txs[i].Memo)` | 1 |
| `GetDirection` | `InferDirection(tx: *Tx, addressSet: mapset.Set): Direction` | `(t, addressSet)` | 1 |
| `GetUTXOValueFor` | `Amount` | `(…)` | 1 |
| `UnmarshalJSON` | `HexNumber` | `(…)` | 1 |

## Callers (reverse)

| Symbol | Called by |
|--------|-----------|
| `Amount` | `GetUTXOValueFor` |
| `AssetID` | `TokenAssetID` |
| `cleanMemo` | `CleanMemos` |
| `DecimalExp` | `FromDecimalExp` |
| `DecimalToSatoshis` | `FromDecimal` |
| `EIP55Checksum` | `ToEIP55ByCoinID` |
| `FindCoinID` | `ParseID` |
| `FindTokenID` | `ParseID` |
| `GetAddressID` | `AddressID` |
| `GetEthereumTokenTypeByIndex` | `GetTokenType` |
| `GetInterfaceSlice` | `GetChunks` |
| `GetInterfaceSliceBatch` | `GetChunks` |
| `GetTokenTypes` | `ParseTokenTypeFromString` |
| `HexNumber` | `UnmarshalJSON` |
| `InferDirection` | `GetDirection` |
| `ParseAmount` | `AddAmount`, `GetAmountValue` |
| `ParseID` | `GetImageURL` |
| `ParseTokenTypeFromString` | `GetTokenVersion` |
| `Remove0x` | `EIP55Checksum` |
| `removeFirstChar` | `FindCoinID`, `FindTokenID` |
| `Round` | `Float64toPrecision` |
| `TokenType` | `GetChainFromAssetType` |
| `ToSatoshi` | `ParseAmount` |
| `Tx` | `UnmarshalJSON` |
| `wrappedTx` | `MarshalJSON` |

## Inferred calls (member, unique name)

> `INFERRED` (confidence 0.9): `obj.method()` where `method` resolves to exactly one definition repo-wide. Receiver type is not resolved; common method names (init/update/onCreate) remain withheld as ambiguous. Use SCIP (`--scip`) for type-resolved member calls at full certainty.

| Caller | Callee | Example args | Sites |
|--------|--------|--------------|-------|
| `HexToDecimal` | `String(): string` | `()` | 1 |
| `ToDecimal` | `String(): string` | `()` | 1 |
| `GetInterfaceSlice` | `Len(): int` | `()` | 2 |
| `GetAssetsIds` | `AssetId(): string` | `()` | 1 |
| `GetChainFromAssetType` | `Acala(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Acalaevm(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Agoric(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Akash(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Algorand(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Aptos(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Arbitrum(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Aurora(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Avalanchec(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Axelar(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Base(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Binance(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Bitcoin(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Blast(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Boba(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Bouncebit(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Callisto(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Cardano(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Celo(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Cfxevm(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Classic(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Cosmos(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Cronos(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Cryptoorg(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Dydx(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Elrond(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Eos(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Ethereum(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Evmos(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Fantom(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Gochain(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Heco(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Hyperevm(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Internet_computer(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Juno(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Kava(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Kavaevm(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Kcc(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Klaytn(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Linea(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Manta(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Mantle(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Megaeth(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Merlin(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Meter(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Metis(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Monad(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Moonbeam(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Moonriver(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Nativeevmos(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Nativeinjective(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Neo(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Neon(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Neutron(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Nuls(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Oasis(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Okc(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Ontology(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Opbnb(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Optimism(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Osmosis(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Plasma(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Poa(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Polygon(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Polygonzkevm(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Ripple(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Robinhoodchain(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Ronin(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Scroll(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Sei(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Seievm(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Smartchain(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Solana(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Sonic(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Stargaze(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Stellar(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Stride(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Sui(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Terra(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Tezos(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Theta(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Thundertoken(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Tia(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Tomochain(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Ton(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Tron(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Vechain(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Wanchain(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Waves(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Xdai(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Zetachain(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Zetaevm(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Zklinknova(): Coin` | `()` | 1 |
| `GetChainFromAssetType` | `Zksync(): Coin` | `()` | 1 |
| `AssetId` | `BuildID(coin: uint, token: string): string` | `(t.Coin, t.TokenID)` | 1 |
| `GetAsset` | `AssetID` | `(…)` | 1 |
| `GetAssetID` | `GetAsset(): coin.AssetID` | `()` | 1 |
| `GetDirection` | `determineTransactionDirection(address: string): Direction` | `(address, t.From, t.To)` | 1 |
| `GetSubscriptionAddresses` | `ParseID(id: string): (uint, string, error)` | `(…)` | 1 |
| `GetSubscriptionAddresses` | `GetAddresses(): []string` | `()` | 1 |
| `GetSubscriptionAddresses` | `GetAsset(): coin.AssetID` | `()` | 1 |
| `IsEVM` | `ParseID(id: string): (uint, string, error)` | `(…)` | 1 |
| `IsEVM` | `GetAsset(): coin.AssetID` | `()` | 1 |
| `HexNumberToUnix` | `ToBig(): *big.Int` | `()` | 1 |

## Graph analytics

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

