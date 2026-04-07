package types

import (
	"fmt"
	"strings"
)

// NewParams creates a new Params instance.
func NewParams(adminAddress string, operatorAddresses []string) Params {
	return Params{
		AdminAddress:      adminAddress,
		OperatorAddresses: operatorAddresses,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams("", []string{})
}

// Validate validates the set of params.
func (p Params) Validate() error {
	seen := map[string]struct{}{}
	for _, addr := range p.OperatorAddresses {
		normalized := strings.TrimSpace(addr)
		if normalized == "" {
			return fmt.Errorf("operator address cannot be empty")
		}
		if _, ok := seen[normalized]; ok {
			return fmt.Errorf("duplicate operator address: %s", normalized)
		}
		seen[normalized] = struct{}{}
	}

	return nil
}
