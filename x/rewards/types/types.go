package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	PartnerStatusActive   = "active"
	PartnerStatusInactive = "inactive"
	PartnerStatusBlocked  = "blocked"
	RewardDenom           = "reward"

	EventTypeIssueReward = "issue_reward"
	EventTypeBurnReward  = "burn_reward"

	AttributeKeyPartnerIndex = "partner_index"
	AttributeKeyRecipient    = "recipient"
	AttributeKeyOwner        = "owner"
	AttributeKeyAmount       = "amount"
	AttributeKeyReason       = "reason"
)

func NormalizePartnerStatus(status string) string {
	return strings.TrimSpace(strings.ToLower(status))
}

func ValidatePartnerStatus(status string) error {
	switch NormalizePartnerStatus(status) {
	case PartnerStatusActive, PartnerStatusInactive, PartnerStatusBlocked:
		return nil
	default:
		return fmt.Errorf("invalid partner status: %s", status)
	}
}

func ValidateRewardCoins(amount sdk.Coins) error {
	if !amount.IsValid() || !amount.IsAllPositive() {
		return ErrInvalidRewardAmount
	}

	for _, coin := range amount {
		if coin.Denom != RewardDenom {
			return errorsmod.Wrapf(ErrInvalidRewardAmount, "only %s denom is supported", RewardDenom)
		}
	}

	return nil
}
