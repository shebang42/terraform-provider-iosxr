// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterBGPNeighbor(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterBGPNeighborConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "remote_as", "65002"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "description", "My Neighbor Description"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "ignore_connected_check", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "ebgp_multihop_maximum_hop_count", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "bfd_minimum_interval", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "bfd_multiplier", "4"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "local_as", "65003"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "local_as_no_prepend", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "local_as_replace_as", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "local_as_dual_as", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "password", "12341C2713181F13253920"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "timers_keepalive_interval", "5"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "timers_holdtime", "20"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "update_source", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor.test", "ttl_security", "false"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterBGPNeighborConfig = `
resource "iosxr_router_bgp_neighbor" "test" {
  as_number = "65001"
  neighbor_address = "10.1.1.2"
  remote_as = "65002"
  description = "My Neighbor Description"
  ignore_connected_check = true
  ebgp_multihop_maximum_hop_count = 10
  bfd_minimum_interval = 10
  bfd_multiplier = 4
  local_as = "65003"
  local_as_no_prepend = true
  local_as_replace_as = true
  local_as_dual_as = true
  password = "12341C2713181F13253920"
  shutdown = false
  timers_keepalive_interval = 5
  timers_holdtime = "20"
  update_source = "GigabitEthernet0/0/0/1"
  ttl_security = false
}

data "iosxr_router_bgp_neighbor" "test" {
  as_number = "65001"
  neighbor_address = "10.1.1.2"
  depends_on = [iosxr_router_bgp_neighbor.test]
}
`
