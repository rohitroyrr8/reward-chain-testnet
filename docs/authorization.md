# Rewards Module Authorization

## Overview

The rewards module now uses a simple parameter-driven authorization model.

## Roles

### Admin
A single `admin_address` can be configured in module params.

### Operators
A list of `operator_addresses` can also be configured in module params.

## Current authorization rules
The following operations require the caller to be either:
- the configured admin, or
- one of the configured operators

Protected operations:
- create partner
- update partner
- delete partner
- issue reward
- burn reward

## Current design notes
- authorization is controlled through module params
- params are updated through the authority-gated `UpdateParams` flow
- this is a simple first pass and should evolve into a clearer role and policy model over time

## Future improvements
- separate permissions by action
- dedicated partner admins vs reward operators
- multisig/governance approval for higher-risk operations
- event emission for authorization changes
