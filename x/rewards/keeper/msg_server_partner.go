package keeper

import (
    "fmt"
	"context"

    "rewardchain/x/rewards/types"
    "cosmossdk.io/collections"
    errorsmod "cosmossdk.io/errors"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


func (k msgServer) CreatePartner(ctx context.Context,  msg *types.MsgCreatePartner) (*types.MsgCreatePartnerResponse, error) {
    if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
    }

    // Check if the value already exists
    ok, err := k.Partner.Has(ctx, msg.Index)
    if err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
    } else if ok {
        return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
    }

    var partner = types.Partner{
        Creator: msg.Creator,
        Index: msg.Index,
        Name: msg.Name,
        Website: msg.Website,
        Description: msg.Description,
        Wallet: msg.Wallet,
        Status: msg.Status,
        
    }

    if err := k.Partner.Set(ctx, partner.Index, partner); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
    }

    return &types.MsgCreatePartnerResponse{}, nil
}

func (k msgServer) UpdatePartner(ctx context.Context,  msg *types.MsgUpdatePartner) (*types.MsgUpdatePartnerResponse, error) {
    if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
    }

    // Check if the value exists
    val, err := k.Partner.Get(ctx, msg.Index)
    if err != nil {
        if errors.Is(err, collections.ErrNotFound) {
            return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
        }

        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
    }

    // Checks if the msg creator is the same as the current owner
    if msg.Creator != val.Creator {
        return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

    var partner = types.Partner{
		Creator: msg.Creator,
		Index: msg.Index,
        Name: msg.Name,
		Website: msg.Website,
		Description: msg.Description,
		Wallet: msg.Wallet,
		Status: msg.Status,
		
	}

    if err := k.Partner.Set(ctx, partner.Index, partner); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update partner")
    }

	return &types.MsgUpdatePartnerResponse{}, nil
}

func (k msgServer) DeletePartner(ctx context.Context,  msg *types.MsgDeletePartner) (*types.MsgDeletePartnerResponse, error) {
    if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
    }

    // Check if the value exists
    val, err := k.Partner.Get(ctx, msg.Index)
    if err != nil {
        if errors.Is(err, collections.ErrNotFound) {
            return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
        }

        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
    }

    // Checks if the msg creator is the same as the current owner
    if msg.Creator != val.Creator {
        return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	if err := k.Partner.Remove(ctx, msg.Index); err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove partner")
    }

	return &types.MsgDeletePartnerResponse{}, nil
}
