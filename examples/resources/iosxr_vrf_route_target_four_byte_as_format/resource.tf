resource "iosxr_vrf_route_target_four_byte_as_format" "example" {
  vrf_name           = "VRF1"
  address_family     = "ipv4"
  sub_address_family = "unicast"
  direction          = "import"
  as_number          = 100000
  index              = 1
  stitching          = true
}
