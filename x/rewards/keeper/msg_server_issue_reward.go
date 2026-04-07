package keeper

import (
	"context"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"rewardchain/x/rewards/types"
)

func (k msgServer) IssueReward(ctx context.Context, msg *types.MsgIssueReward) (*types.MsgIssueRewardResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	recipient, err := k.addressCodec.StringToBytes(msg.Recipient)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
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

	if !sdk.AccAddress(partnerWallet).Equals(sdk.AccAddress(recipient)) {
		return nil, errorsmod.Wrap(types.ErrInvalidPartnerWallet, "recipient must match partner wallet")
	}

	if err := k.bankKeeper.SendCoins(ctx, sdk.AccAddress(partnerWallet), sdk.AccAddress(recipient), amount); err != nil {
		return nil, errorsmod.Wrap(err, "failed to issue reward")
	}

	return &types.MsgIssueRewardResponse{}, nil
}
