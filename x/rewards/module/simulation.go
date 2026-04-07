package rewards

import (
	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"

	rewardssimulation "rewardchain/x/rewards/simulation"
	"rewardchain/x/rewards/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rewardsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PartnerMap: []types.Partner{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rewardsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreatePartner          = "op_weight_msg_rewards"
		defaultWeightMsgCreatePartner int = 100
	)

	var weightMsgCreatePartner int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePartner, &weightMsgCreatePartner, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePartner = defaultWeightMsgCreatePartner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePartner,
		rewardssimulation.SimulateMsgCreatePartner(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdatePartner          = "op_weight_msg_rewards"
		defaultWeightMsgUpdatePartner int = 100
	)

	var weightMsgUpdatePartner int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePartner, &weightMsgUpdatePartner, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePartner = defaultWeightMsgUpdatePartner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePartner,
		rewardssimulation.SimulateMsgUpdatePartner(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeletePartner          = "op_weight_msg_rewards"
		defaultWeightMsgDeletePartner int = 100
	)

	var weightMsgDeletePartner int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePartner, &weightMsgDeletePartner, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePartner = defaultWeightMsgDeletePartner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePartner,
		rewardssimulation.SimulateMsgDeletePartner(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}