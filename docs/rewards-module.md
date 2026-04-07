# Rewards Module

## Overview

The `rewards` module is the custom business module for Reward Chain Testnet.

Current responsibilities:
- partner record management
- reward issuance flow
- reward burn flow
- operator/admin-gated authorization for sensitive actions

## Partner model

A partner currently includes:
- `index`
- `name`
- `website`
- `description`
- `wallet`
- `status`
- `creator`

For reward operations, the current implementation assumes:
- the partner exists
- the partner status is `active`
- the partner wallet is the wallet used for reward operations

## Issue reward

`issue-reward` is intended to send reward-denominated coins to an active partner wallet.

Validation currently includes:
- valid creator address
- valid recipient address
- partner exists
- partner status is `active`
- non-empty reason
- valid positive amount
- only the `reward` denom is accepted
- recipient must match the stored partner wallet
- emits an `issue_reward` event on success

## Burn reward

`burn-reward` is intended to burn reward-denominated coins from an active partner wallet.

Validation currently includes:
- valid creator address
- valid owner address
- partner exists
- partner status is `active`
- non-empty reason
- valid positive amount
- only the `reward` denom is accepted
- owner must match the stored partner wallet
- emits a `burn_reward` event on success

Current burn flow:
1. transfer coins from the owner account into the module account
2. burn those coins from the module account

## Notes

This is an initial functional pass. Future improvements should include:
- richer partner status lifecycle
- explicit authorization model for operators/admins
- event emission for reward issue/burn actions
- better accounting and reward ledger tracking
- reward-specific denom constraints
- liquidity and swap integration
