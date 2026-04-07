package keeper

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"rewardchain/x/rewards/types"
)

func (k msgServer) CreatePartner(ctx context.Context, msg *types.MsgCreatePartner) (*types.MsgCreatePartnerResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}
	if err := k.isAuthorizedOperator(ctx, msg.Creator); err != nil {
		return nil, err
	}
	if _, err := k.addressCodec.StringToBytes(msg.Wallet); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid wallet address: %s", err))
	}
	if strings.TrimSpace(msg.Name) == "" {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "partner name cannot be empty")
	}
	if err := types.ValidatePartnerStatus(msg.Status); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ok, err := k.Partner.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	partner := types.Partner{
		Creator:     msg.Creator,
		Index:       msg.Index,
		Name:        strings.TrimSpace(msg.Name),
		Website:     strings.TrimSpace(msg.Website),
		Description: strings.TrimSpace(msg.Description),
		Wallet:      msg.Wallet,
		Status:      types.NormalizePartnerStatus(msg.Status),
	}

	if err := k.Partner.Set(ctx, partner.Index, partner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreatePartnerResponse{}, nil
}

func (k msgServer) UpdatePartner(ctx context.Context, msg *types.MsgUpdatePartner) (*types.MsgUpdatePartnerResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}
	if err := k.isAuthorizedOperator(ctx, msg.Creator); err != nil {
		return nil, err
	}
	if _, err := k.addressCodec.StringToBytes(msg.Wallet); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid wallet address: %s", err))
	}
	if strings.TrimSpace(msg.Name) == "" {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "partner name cannot be empty")
	}
	if err := types.ValidatePartnerStatus(msg.Status); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	_, err := k.Partner.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	partner := types.Partner{
		Creator:     msg.Creator,
		Index:       msg.Index,
		Name:        strings.TrimSpace(msg.Name),
		Website:     strings.TrimSpace(msg.Website),
		Description: strings.TrimSpace(msg.Description),
		Wallet:      msg.Wallet,
		Status:      types.NormalizePartnerStatus(msg.Status),
	}

	if err := k.Partner.Set(ctx, partner.Index, partner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update partner")
	}

	return &types.MsgUpdatePartnerResponse{}, nil
}

func (k msgServer) DeletePartner(ctx context.Context, msg *types.MsgDeletePartner) (*types.MsgDeletePartnerResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}
	if err := k.isAuthorizedOperator(ctx, msg.Creator); err != nil {
		return nil, err
	}

	_, err := k.Partner.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	if err := k.Partner.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove partner")
	}

	return &types.MsgDeletePartnerResponse{}, nil
}
