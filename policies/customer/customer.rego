package customer


status_is_active {
  input.customer.status == "active" 
}

is_operations_user {
  input.user.roles[_] == "operations"
}

# the customer can only access if 
#  * has an active status 
allow_access {
  status_is_active 
}

# allow power users to overrule access
allow_access {
  is_operations_user
}
