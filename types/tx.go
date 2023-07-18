package types

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"

	"github.com/trustwallet/go-primitives/asset"
	"github.com/trustwallet/go-primitives/coin"
)

const (
	StatusCompleted Status = "completed"
	StatusPending   Status = "pending"
	StatusError     Status = "error"

	DirectionOutgoing Direction = "outgoing"
	DirectionIncoming Direction = "incoming"
	DirectionSelf     Direction = "yourself"

	TxTransfer          TransactionType = "transfer"
	TxSwap              TransactionType = "swap"
	TxContractCall      TransactionType = "contract_call"
	TxStakeClaimRewards TransactionType = "stake_claim_rewards"
	TxStakeDelegate     TransactionType = "stake_delegate"
	TxStakeUndelegate   TransactionType = "stake_undelegate"
	TxStakeRedelegate   TransactionType = "stake_redelegate"
	TxStakeCompound     TransactionType = "stake_compound"
	TxTransferNFT       TransactionType = "transfer_nft"
	TxTransferICS20     TransactionType = "transfer_ics20"
)

var SupportedTypes = []TransactionType{
	TxTransfer, TxSwap, TxContractCall, TxStakeClaimRewards, TxStakeDelegate, TxStakeUndelegate, TxStakeRedelegate,
	TxStakeCompound,
	TxTransferNFT,
	TxTransferICS20,
}

// Transaction fields
type (
	Direction       string
	Status          string
	TransactionType string
	KeyType         string
	KeyTitle        string

	// Amount is a positive decimal integer string.
	// It is written in the smallest possible unit (e.g. Wei, Satoshis)
	Amount string
)

// Data objects
type (
	Block struct {
		Number int64 `json:"number"`
		Txs    []Tx  `json:"txs"`
	}

	TxPage struct {
		Total int  `json:"total"`
		Docs  []Tx `json:"docs"`
	}

	// Tx describes an on-chain transaction generically
	Tx struct {
		// Unique identifier
		ID string `json:"id"`

		// Address of the transaction sender
		From string `json:"from"`

		// Address of the transaction recipient
		To string `json:"to"`

		// Unix timestamp of the block the transaction was included in
		BlockCreatedAt int64 `json:"block_created_at"`

		// Height of the block the transaction was included in
		Block uint64 `json:"block"`

		// Status of the transaction e.g: "completed", "pending", "error"
		Status Status `json:"status"`

		// Empty if the transaction "completed" or "pending", else error explaining why the transaction failed (optional)
		Error string `json:"error,omitempty"`

		// Transaction nonce or sequence
		Sequence uint64 `json:"sequence"`

		// Type of metadata
		Type TransactionType `json:"type"`

		// Transaction Direction
		Direction Direction `json:"direction,omitempty"`

		Inputs  []TxOutput `json:"inputs,omitempty"`
		Outputs []TxOutput `json:"outputs,omitempty"`

		Tokens []Asset `json:"tokens,omitempty"`

		Memo string `json:"memo,omitempty"`

		Fee Fee `json:"fee"`

		// Metadata data object
		Metadata interface{} `json:"metadata"`

		// Create At indicates transactions creation time in database, Unix
		CreatedAt int64 `json:"created_at"`
	}

	// Every transaction consumes some Fee
	Fee struct {
		Asset coin.AssetID `json:"asset"`
		Value Amount       `json:"value"`
	}

	// UTXO transactions consist of a set of inputs and a set of outputs
	// both represented by TxOutput model
	TxOutput struct {
		Address string `json:"address"`
		Value   Amount `json:"value"`
	}

	// Transfer describes the transfer of currency
	Transfer struct {
		Asset coin.AssetID `json:"asset"`
		Value Amount       `json:"value"`
	}

	// ICS20Transfer describes the transfer of the ICS20 token
	ICS20Transfer struct {
		Transfer

		SourcePort    string `json:"source_port"`
		SourceChannel string `json:"source_channel"`
	}

	// TransferNFT describes the transfer of the NFT token
	TransferNFT struct {
		Asset            coin.AssetID `json:"asset"`
		Collection       string       `json:"collection"`
		CollectibleID    string       `json:"collectible_id"`
		CollectionSymbol string       `json:"collection_symbol"`
		Value            Amount       `json:"value"`
	}

	Swap struct {
		From Transfer `json:"from"`
		To   Transfer `json:"to"`
	}

	// ContractCall describes a
	ContractCall struct {
		Asset coin.AssetID `json:"asset"`
		Value Amount       `json:"value"`
		Input string       `json:"input"`
	}

	Txs []Tx

	AssetHolder interface {
		GetAsset() coin.AssetID
	}

	Validator interface {
		Validate() error
	}
)

var (
	EmptyTxPage = TxPage{Total: 0, Docs: Txs{}}
)

func NewTxPage(txs Txs) TxPage {
	if txs == nil {
		txs = Txs{}
	}
	return TxPage{
		Total: len(txs),
		Docs:  txs,
	}
}

func (txs Txs) FilterUniqueID() Txs {
	keys := make(map[string]bool)
	list := make(Txs, 0)
	for _, entry := range txs {
		if _, value := keys[entry.ID]; !value {
			keys[entry.ID] = true
			list = append(list, entry)
		}
	}
	return list
}

func (txs Txs) CleanMemos() {
	for i := range txs {
		txs[i].Memo = cleanMemo(txs[i].Memo)
	}
}

func (txs Txs) SortByBlockCreationTime() Txs {
	sort.Slice(txs, func(i, j int) bool {
		return txs[i].BlockCreatedAt > txs[j].BlockCreatedAt
	})
	return txs
}

func (txs Txs) FilterTransactionsByType(types []TransactionType) Txs {
	result := make(Txs, 0)
	for _, tx := range txs {
		for _, t := range types {
			if tx.Type == t {
				result = append(result, tx)
			}
		}
	}

	return result
}

func (t *Transfer) GetAsset() coin.AssetID {
	return t.Asset
}

func (t *Transfer) Validate() error {
	if t.Value == "" {
		return fmt.Errorf("emtpy transfer value")
	}

	if t.Asset == "" {
		return fmt.Errorf("empty transfer asset")
	}

	return nil
}

func (cc *ContractCall) GetAsset() coin.AssetID {
	return cc.Asset
}

func (cc *ContractCall) Validate() error {
	if cc.Value == "" {
		return fmt.Errorf("empty contract call value")
	}

	if cc.Asset == "" {
		return fmt.Errorf("empty contract call asset")
	}

	return nil
}

func (cc *TransferNFT) GetAsset() coin.AssetID {
	return cc.Asset
}

func (cc *TransferNFT) Validate() error {
	if cc.CollectibleID == "" {
		return fmt.Errorf("empty transfer NFT collectible ID value")
	}

	if cc.Asset == "" {
		return fmt.Errorf("empty transfer NFT asset")
	}

	return nil
}

func (s *Swap) GetAsset() coin.AssetID {
	return coin.AssetID(strings.Split(string(s.From.Asset), "_")[0])
}

func (i *ICS20Transfer) GetAsset() coin.AssetID {
	return i.Asset
}

func (i *ICS20Transfer) Validate() error {
	if i.Value == "" {
		return fmt.Errorf("empty ICS20 transfer value")
	}

	if i.Asset == "" {
		return fmt.Errorf("empty ICS20 transfer asset")
	}

	return nil
}

func cleanMemo(memo string) string {
	if len(memo) == 0 {
		return ""
	}

	_, err := strconv.ParseFloat(memo, 64)
	if err != nil {
		return ""
	}

	return memo
}

func (t *Tx) GetAddresses() []string {
	addresses := make([]string, 0)
	switch t.Type {
	case TxTransfer, TxTransferNFT:
		if len(t.Inputs) > 0 || len(t.Outputs) > 0 {
			uniqueAddresses := make(map[string]struct{})
			for _, input := range t.Inputs {
				uniqueAddresses[input.Address] = struct{}{}
			}

			for _, output := range t.Outputs {
				uniqueAddresses[output.Address] = struct{}{}
			}

			for address := range uniqueAddresses {
				addresses = append(addresses, address)
			}

			return addresses
		}

		return append(addresses, t.From, t.To)
	case TxContractCall:
		return append(addresses, t.From, t.To)
	case TxSwap:
		return append(addresses, t.From, t.To)
	case TxStakeDelegate, TxStakeRedelegate, TxStakeUndelegate, TxStakeClaimRewards, TxStakeCompound:
		return append(addresses, t.From)
	case TxTransferICS20:
		if t.inSameSubChain() {
			addresses = append(addresses, t.To)
		}
		return append(addresses, t.From)
	default:
		return addresses
	}
}

func (t *Tx) GetSubscriptionAddresses() ([]string, error) {
	coin, _, err := asset.ParseID(string(t.Metadata.(AssetHolder).GetAsset()))
	if err != nil {
		return nil, err
	}

	addresses := t.GetAddresses()
	result := make([]string, len(addresses))
	for i, a := range addresses {
		result[i] = fmt.Sprintf("%d_%s", coin, a)
	}

	return result, nil
}

func (t *Tx) GetDirection(address string) Direction {
	if len(t.Direction) > 0 {
		return t.Direction
	}

	if len(t.Inputs) > 0 && len(t.Outputs) > 0 {
		addressSet := mapset.NewSet(address)
		return InferDirection(t, addressSet)
	}

	return t.determineTransactionDirection(address, t.From, t.To)
}

func (t *Tx) GetAssetID() *coin.AssetID {
	if t.Metadata == nil {
		return nil
	}

	assetHolder, ok := t.Metadata.(AssetHolder)
	if !ok {
		return nil
	}

	assetID := assetHolder.GetAsset()
	return &assetID
}

func (t *Tx) determineTransactionDirection(address, from, to string) Direction {
	if t.Type == TxStakeUndelegate || t.Type == TxStakeClaimRewards {
		return DirectionIncoming
	}

	if address == to {
		if from == to {
			return DirectionSelf
		}
		return DirectionIncoming
	}
	return DirectionOutgoing
}

func (t *Tx) IsUTXO() bool {
	return t.Type == TxTransfer && len(t.Outputs) > 0
}

func (t *Tx) IsEVM() (bool, error) {
	if t.Metadata == nil {
		return false, nil
	}

	assetHolder, ok := t.Metadata.(AssetHolder)
	if !ok {
		return false, nil
	}

	coinID, _, err := asset.ParseID(string(assetHolder.GetAsset()))
	if err != nil {
		return false, err
	}

	return coin.IsEVM(coinID), nil
}

func (t *Tx) GetUTXOValueFor(address string) (Amount, error) {
	isTransferOut := false
	isSelf := true

	totalInputValue := big.NewInt(0)
	addressInputValue := big.NewInt(0)
	for _, input := range t.Inputs {
		value, ok := big.NewInt(0).SetString(string(input.Value), 10)
		if !ok {
			return "0", fmt.Errorf("invalid input value for address %s: %s", input.Address, input.Value)
		}

		totalInputValue = totalInputValue.Add(totalInputValue, value)

		if input.Address == address {
			addressInputValue = value
			isTransferOut = true
		}
	}

	addressOutputValue := big.NewInt(0)
	totalOutputValue := big.NewInt(0)
	for _, output := range t.Outputs {
		value, ok := big.NewInt(0).SetString(string(output.Value), 10)
		if !ok {
			return "0", fmt.Errorf("invalid output value for address %s: %s", output.Address, output.Value)
		}
		totalOutputValue = totalOutputValue.Add(totalOutputValue, value)
		if output.Address == address {
			addressOutputValue = addressOutputValue.Add(addressOutputValue, value)
		} else {
			isSelf = false
		}
	}

	var result *big.Int
	if isTransferOut && !isSelf {
		if addressInputValue.Cmp(addressOutputValue) < 0 {
			result = big.NewInt(0) // address received more than sent, although it's an outgoing tx
		} else {
			//addressInputValue - (totalInputValue-totalOutputValue)/uint64(len(t.Inputs)) - addressOutputValue
			totalTransferred := totalInputValue.Sub(totalInputValue, totalOutputValue)
			avgSent := totalTransferred.Div(totalTransferred, big.NewInt(int64(len(t.Inputs))))
			output := addressInputValue.Sub(addressInputValue, addressOutputValue)

			// for utxo there is no way to define the exact amount sent
			// because there is many senders and many recipients
			result = output.Sub(output, avgSent)
		}
	} else {
		result = addressOutputValue
	}

	return Amount(fmt.Sprintf("%d", result)), nil
}

func (t *Tx) inSameSubChain() bool {
	//account address prefix coin.COSMOS:      "cosmos"
	//account address prefix coin.OSMOSIS:     "osmo"
	//account address prefix coin.TERRA:       "terra"
	//account address prefix coin.KAVA:        "kava"
	//account address prefix coin.NATIVEEVMOS: "evmos"
	//account address prefix coin.STRIDE:      "stride"
	//account address prefix coin.NEUTRON:     "neutron"
	minPrefixLen := 4
	if len(t.From) <= minPrefixLen || len(t.To) <= minPrefixLen {
		return false
	}
	return t.From[:minPrefixLen] == t.To[:minPrefixLen]
}

func InferDirection(tx *Tx, addressSet mapset.Set) Direction {
	inputSet := mapset.NewSet()
	for _, address := range tx.Inputs {
		inputSet.Add(address.Address)
	}
	outputSet := mapset.NewSet()
	for _, address := range tx.Outputs {
		outputSet.Add(address.Address)
	}
	intersect := addressSet.Intersect(inputSet)
	if intersect.Cardinality() == 0 {
		return DirectionIncoming
	}
	if outputSet.IsProperSubset(addressSet) || outputSet.Equal(inputSet) {
		return DirectionSelf
	}
	return DirectionOutgoing
}

func IsTxTypeAmong(txType TransactionType, types []TransactionType) bool {
	result := false
	for _, t := range types {
		if txType == t {
			result = true
			break
		}
	}

	return result
}
