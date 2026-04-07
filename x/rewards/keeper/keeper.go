package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"rewardchain/x/rewards/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	bankKeeper   types.BankKeeper
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema  collections.Schema
	Params  collections.Item[types.Params]
	Partner collections.Map[string, types.Partner]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	bankKeeper types.BankKeeper,
	authority []byte,
) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		bankKeeper:   bankKeeper,
		authority:    authority,
		Params:       collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		Partner:      collections.NewMap(sb, types.PartnerKey, "partner", collections.StringKey, codec.CollValue[types.Partner](cdc)),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}
