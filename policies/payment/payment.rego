package payment

# Defines if the customer is allowed to make a payment
default can_initiate_payment = false

import data.customer
import data.system

default_daily_tx_limits := {
  "gbp": 500,
  "usd": 1000
}

# Get the default transaction limit when no customer limit is set
tx_limit = x {
  x := default_daily_tx_limits[input.payment.amount.currency]
  input.customer.settings.daily_transaction_limits[input.payment.amount.currency] < 0
}

# Get the customer transaction limit when a limit is set
tx_limit = x {
  x := input.customer.settings.daily_transaction_limits[input.payment.amount.currency]
  input.customer.settings.daily_transaction_limits[input.payment.amount.currency] > -1
}

# Take total spend (daily spend + amount to send)
tx_total = x {
  spend_today := input.transaction_info.spend_today[input.payment.amount.currency]
  amount_to_spend := input.payment.amount.value 
  x := spend_today + amount_to_spend
}

allow_payment {
  system.is_active_service("payments")
  customer.allow_access
  tx_total < tx_limit
}