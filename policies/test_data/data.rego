package testdata 

usdPayment := { 
  "customer": {
      "id": "cus_12345678901234567890",
      "status": "active",
      "segment": "exclusive",
      "settings": {
          "daily_transaction_limits": {
              "gbp": 300,
              "usd": -1 # no limit set
          }
      }
  },
  "payment": {
      "from_account_id": "bac_12345678901234567890",
      "to_benificiary_id": "ben_12345678901234567890",
      "to_currency": "gbp",
      "amount": {
          "currency": "usd",
          "value": 100
      }
  },
  "transaction_info": {
      "spend_today": {
          "gbp": 100,
          "usd": 0
      },
  },
  "products": {
    "bank_accounts": {
      "active_deposit_accounts_count": 0,
      "active_ftd_accounts_count": 2
    }
  }
}

gbpPayment := { 
  "customer": {
      "id": "cus_12345678901234567890",
      "status": "suspended",
      "segment": "exclusive",
      "settings": {
          "daily_transaction_limits": {
              "gbp": 300,
              "usd": -1 # no limit set
          }
      }
  },
  "payment": {
      "from_account_id": "bac_12345678901234567890",
      "to_benificiary_id": "ben_12345678901234567890",
      "to_currency": "gbp",
      "amount": {
          "currency": "gbp",
          "value": 100
      }
  },
  "transaction_info": {
      "spend_today": {
          "gbp": 100,
          "usd": 0
      },
  },
  "user": {
      "id": "emp_12345678901234567890",
      "roles": ["operations"]
  }
}