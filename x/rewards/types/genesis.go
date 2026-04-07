package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:     DefaultParams(),
		PartnerMap: []Partner{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	partnerIndexMap := make(map[string]struct{})

	for _, elem := range gs.PartnerMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := partnerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for partner")
		}
		partnerIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}