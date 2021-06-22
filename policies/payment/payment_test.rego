package payment_test

import data.payment
import data.testdata

test_should_have_usd_tx_limit_1000 {
    1000 = payment.tx_limit with input as testdata.usdPayment
}

test_should_have_gbp_tx_limit_300 {
    300 = payment.tx_limit with input as testdata.gbpPayment
}

test_should_have_a_total_gbp_spend_of_200 {
    200 = payment.tx_total with input as testdata.gbpPayment
}

test_should_allow_gbp_payment {
    payment.allow_payment with input as testdata.gbpPayment
}
