
NOTE: *all bellow examples are just sample data to showcase fields used in data passed to hooks.*

## after_candidates_generated

This hook is capable of mutating data.
```json
{
  "cycle": 580,
  "candidates": [
    {
      "source": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "recipient": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "fee_rate": 5,
      "staked_balance": "1000000000",
      "delegated_balance": "1000000000",
      "is_invalid": true,
      "is_emptied": true,
      "is_baker_paying_tx_fee": true,
      "is_baker_paying_allocation_tx_fee": true,
      "invalid_because": "reason"
    }
  ]
}
```

## after_bonds_distributed

This hook is capable of mutating data.
```json
{
  "cycle": 580,
  "candidates": [
    {
      "source": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "recipient": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "fee_rate": 5,
      "staked_balance": "1000000000",
      "delegated_balance": "1000000000",
      "is_invalid": true,
      "is_emptied": true,
      "is_baker_paying_tx_fee": true,
      "is_baker_paying_allocation_tx_fee": true,
      "invalid_because": "reason",
      "bonds_amount": "1000000000",
      "tx_kind": "fa1",
      "fa_token_id": "10",
      "fa_contract": "KT18amZmM5W7qDWVt2pH6uj7sCEd3kbzLrHT"
    }
  ]
}
```

## check_balance

This hook is NOT capable of mutating data.
```json
{
  "skip_mav_check": true,
  "is_sufficient": true,
  "message": "This message is used to carry errors from hook to the caller.",
  "payouts": [
    {
      "source": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "recipient": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "fee_rate": 5,
      "staked_balance": "1000000000",
      "delegated_balance": "1000000000",
      "is_invalid": true,
      "is_emptied": true,
      "is_baker_paying_tx_fee": true,
      "is_baker_paying_allocation_tx_fee": true,
      "invalid_because": "reason",
      "bonds_amount": "1000000000",
      "tx_kind": "fa1",
      "fa_token_id": "10",
      "fa_contract": "KT18amZmM5W7qDWVt2pH6uj7sCEd3kbzLrHT",
      "fee": "1000000000"
    }
  ]
}
```

## on_fees_collection

This hook is capable of mutating data.
```json
{
  "cycle": 580,
  "candidates": [
    {
      "source": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "recipient": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "fee_rate": 5,
      "staked_balance": "1000000000",
      "delegated_balance": "1000000000",
      "is_invalid": true,
      "is_emptied": true,
      "is_baker_paying_tx_fee": true,
      "is_baker_paying_allocation_tx_fee": true,
      "invalid_because": "reason",
      "bonds_amount": "1000000000",
      "tx_kind": "fa1",
      "fa_token_id": "10",
      "fa_contract": "KT18amZmM5W7qDWVt2pH6uj7sCEd3kbzLrHT",
      "fee": "1000000000"
    }
  ]
}
```

## after_payouts_blueprint_generated

This hook is NOT capable of mutating data *currently*.
```json
{
  "cycles": 1,
  "payouts": [
    {
      "baker": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "delegator": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "cycle": 1,
      "recipient": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "kind": "invalid",
      "tx_kind": "fa1",
      "fa_token_id": "10",
      "fa_contract": "KT18amZmM5W7qDWVt2pH6uj7sCEd3kbzLrHT",
      "delegator_balance": "1000000000",
      "amount": "1000000000",
      "fee_rate": 5,
      "fee": "1000000000",
      "op_limits": {
        "transaction_fee": 1,
        "storage_limit": 1,
        "gas_limit": 1,
        "deserialization_gas_limit": 1,
        "allocation_burn": 1,
        "storage_burn": 1
      },
      "note": "reason"
    }
  ],
  "summary": {
    "cycle": 1,
    "delegators": 2,
    "paid_delegators": 1,
    "own_staked_balance": "1000000000",
    "own_delegated_balance": "0",
    "external_staked_balance": "0",
    "external_delegated_balance": "0",
    "cycle_fees": "1000000000",
    "cycle_rewards": "1000000000",
    "distributed_rewards": "1000000000",
    "bond_income": "1000000000",
    "fee_income": "1000000000",
    "total_income": "1000000000",
    "donated_bonds": "1000000000",
    "donated_fees": "1000000000",
    "donated_total": "1000000000",
    "timestamp": "2023-01-01T00:00:00Z"
  }
}
```

## after_payouts_prepared

This hook is capable of mutating data *currently*.
```json
{
  "recipes": null,
  "payouts": [
    {
      "baker": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "delegator": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "cycle": 1,
      "recipient": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "kind": "invalid",
      "tx_kind": "fa1",
      "fa_token_id": "10",
      "fa_contract": "KT18amZmM5W7qDWVt2pH6uj7sCEd3kbzLrHT",
      "delegator_balance": "1000000000",
      "amount": "1000000000",
      "fee_rate": 5,
      "fee": "1000000000",
      "op_limits": {
        "transaction_fee": 1,
        "storage_limit": 1,
        "gas_limit": 1,
        "deserialization_gas_limit": 1,
        "allocation_burn": 1,
        "storage_burn": 1
      },
      "note": "reason"
    }
  ],
  "invalid_payouts": null,
  "reports_of_past_succesful_payouts": [
    {
      "id": "fd8d0230c9d70458",
      "baker": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "timestamp": "2024-09-07T09:25:23.379448699Z",
      "cycle": 1,
      "kind": "invalid",
      "tx_kind": "fa1",
      "contract": "KT18amZmM5W7qDWVt2pH6uj7sCEd3kbzLrHT",
      "token_id": "10",
      "delegator": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "delegator_balance": "1000000000",
      "recipient": "mv1QiogZoD9f7o83b3BMWK977KHxw3zhN7cJ",
      "amount": "1000000000",
      "fee_rate": 5,
      "fee": "1000000000",
      "tx_fee": 1,
      "op_hash": "oneDGhZacw99EEFaYDTtWfz5QEhUW3PPVFsHa7GShnLPuDn7gSd",
      "success": true,
      "note": "reason"
    }
  ]
}
```

