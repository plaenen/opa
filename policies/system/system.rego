package system  
  
activeSystems = ["payments", "onboarding"]


is_active_service(service) {
  activeSystems[_] = service 
}

is_in_maintenance(service) {
  not is_active_service(service)
}