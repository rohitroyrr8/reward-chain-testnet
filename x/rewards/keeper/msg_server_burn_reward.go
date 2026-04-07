package keeper

import (
	"context"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"rewardchain/x/rewards/types"
)

func (k msgServer) BurnReward(ctx context.Context, msg *types.MsgBurnReward) (*types.MsgBurnRewardResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	owner, err := k.addressCodec.StringToBytes(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}

	partner, err := k.Partner.Get(ctx, msg.PartnerIndex)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrPartnerNotFound, msg.PartnerIndex)
	}

	if strings.TrimSpace(strings.ToLower(partner.Status)) != "active" {
		return nil, errorsmod.Wrap(types.ErrInvalidPartnerStatus, partner.Status)
	}

	if strings.TrimSpace(msg.Reason) == "" {
		return nil, types.ErrEmptyReason
	}

	amount := sdk.NewCoins(msg.Amount...)
	if !amount.IsValid() || !amount.IsAllPositive() {
		return nil, types.ErrInvalidRewardAmount
	}

	partnerWallet, err := k.addressCodec.StringToBytes(partner.Wallet)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidPartnerWallet, "stored partner wallet is invalid")
	}

	if !sdk.AccAddress(partnerWallet).Equals(sdk.AccAddress(owner)) {
		return nil, errorsmod.Wrap(types.ErrInvalidPartnerWallet, "owner must match partner wallet")
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(owner), types.ModuleName, amount); err != nil {
		return nil, errorsmod.Wrap(err, "failed to move reward into module for burning")
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, amount); err != nil {
		return nil, errorsmod.Wrap(err, "failed to burn reward")
	}

	return &types.MsgBurnRewardResponse{}, nil
}
