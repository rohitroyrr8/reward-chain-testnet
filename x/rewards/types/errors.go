package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/rewards module sentinel errors
var (
	ErrInvalidSigner        = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrPartnerNotFound      = errors.Register(ModuleName, 1101, "partner not found")
	ErrInvalidPartnerWallet = errors.Register(ModuleName, 1102, "partner wallet does not match stored partner wallet")
	ErrEmptyReason          = errors.Register(ModuleName, 1103, "reason cannot be empty")
	ErrInvalidRewardAmount  = errors.Register(ModuleName, 1104, "amount must be valid and positive")
	ErrInvalidPartnerStatus = errors.Register(ModuleName, 1105, "partner is not active")
)
