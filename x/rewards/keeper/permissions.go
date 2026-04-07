package keeper

import (
	"context"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) isAuthorizedOperator(ctx context.Context, actor string) error {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "module params not configured")
	}

	actor = strings.TrimSpace(actor)
	if actor == "" {
		return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "empty actor")
	}

	if strings.TrimSpace(params.AdminAddress) == actor {
		return nil
	}

	for _, operator := range params.OperatorAddresses {
		if strings.TrimSpace(operator) == actor {
			return nil
		}
	}

	return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "actor is not an authorized rewards operator")
}
