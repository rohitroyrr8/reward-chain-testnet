package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"rewardchain/x/rewards/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				PartnerMap: []types.Partner{{Index: "0"}, {Index: "1"}}},
			valid: true,
		}, {
			desc: "duplicated partner",
			genState: &types.GenesisState{
				PartnerMap: []types.Partner{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}