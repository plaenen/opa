package customer.test

import data.customer

test_access_allowed {
    customer.allow_access with input as {
        "customer": {
            "status": "active"
        }
    } 
}

test_access_denied {
    not customer.allow_access with input as {
        "customer": {
            "status": "suspended"
        }
    } 
}


