resource "iosxr_router_ospf_redistribute_ospf" "example" {
  process_name        = "OSPF1"
  instance_name       = "OSPF2"
  match_internal      = true
  match_external      = false
  match_nssa_external = false
  tag                 = 4
  metric_type         = "1"
}
