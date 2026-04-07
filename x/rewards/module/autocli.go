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
					Short:     "List all partner",
				},
				{
					RpcMethod:      "GetPartner",
					Use:            "get-partner [id]",
					Short:          "Gets a partner",
					Alias:          []string{"show-partner"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
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
					Short:          "Update partner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "name"}, {ProtoField: "website"}, {ProtoField: "description"}, {ProtoField: "wallet"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "DeletePartner",
					Use:            "delete-partner [index]",
					Short:          "Delete partner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
			},
		},
	}
}
