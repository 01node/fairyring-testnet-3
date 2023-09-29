package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EncryptedTxArray:       []EncryptedTxArray{},
		AggregatedKeyShareList: []AggregatedKeyShare{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	encryptedTxArrIndexMap := make(map[string]struct{})
	for height, elem := range gs.EncryptedTxArray {
		for index, item := range elem.EncryptedTx {
			if item.Index != uint64(index) {
				return fmt.Errorf("encrypted tx index does not match")
			}

			if item.TargetCondition != string(height) {
				return fmt.Errorf("encrypted tx target height does not match")
			}
		}
		index := string(EncryptedTxAllFromCondition(string(height)))
		if _, ok := encryptedTxArrIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for encryptedTxArr")
		}
		encryptedTxArrIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in aggregatedKeyShare
	aggregatedKeyShareIndexMap := make(map[string]struct{})

	for _, elem := range gs.AggregatedKeyShareList {
		index := string(AggregatedKeyShareKey(elem.Condition))
		if _, ok := aggregatedKeyShareIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for aggregatedKeyShare")
		}
		aggregatedKeyShareIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	// Check for valid addresses in NONCE
	for _, elem := range gs.ConditionalencNonceList {
		_, err := sdk.AccAddressFromBech32(elem.Address)
		if err != nil {
			return err
		}
	}

	return gs.Params.Validate()
}
