package rewards

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"rewardchain/x/rewards/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListPartner",
					Use:       "list-partner",
					Short:     "List all partner records",
				},
				{
					RpcMethod:      "GetPartner",
					Use:            "get-partner [index]",
					Short:          "Get a partner by index",
					Alias:          []string{"show-partner"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true,
				},
				{
					RpcMethod:      "CreatePartner",
					Use:            "create-partner [index] [name] [website] [description] [wallet] [status]",
					Short:          "Create a new partner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "name"}, {ProtoField: "website"}, {ProtoField: "description"}, {ProtoField: "wallet"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "UpdatePartner",
					Use:            "update-partner [index] [name] [website] [description] [wallet] [status]",
					Short:          "Update a partner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "name"}, {ProtoField: "website"}, {ProtoField: "description"}, {ProtoField: "wallet"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "DeletePartner",
					Use:            "delete-partner [index]",
					Short:          "Delete a partner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "IssueReward",
					Use:            "issue-reward [partner-index] [recipient] [amount] [reason]",
					Short:          "Issue a reward to an active partner wallet",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "partner_index"}, {ProtoField: "recipient"}, {ProtoField: "amount"}, {ProtoField: "reason"}},
				},
				{
					RpcMethod:      "BurnReward",
					Use:            "burn-reward [partner-index] [owner] [amount] [reason]",
					Short:          "Burn reward tokens from an active partner wallet",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "partner_index"}, {ProtoField: "owner"}, {ProtoField: "amount"}, {ProtoField: "reason"}},
				},
			},
		},
	}
}
