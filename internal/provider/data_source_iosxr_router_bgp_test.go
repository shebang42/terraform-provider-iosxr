// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterBGP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterBGPConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "default_information_originate", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "default_metric", "125"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "nsr_disable", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "bgp_redistribute_internal", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "segment_routing_srv6_locator", "locator11"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "timers_bgp_keepalive_interval", "5"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "timers_bgp_holdtime", "20"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "bgp_router_id", "22.22.22.22"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "bgp_graceful_restart_graceful_reset", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "ibgp_policy_out_enforce_modifications", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "bgp_log_neighbor_changes_detail", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "bfd_minimum_interval", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "bfd_multiplier", "4"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "nexthop_validation_color_extcomm_sr_policy", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "nexthop_validation_color_extcomm_disable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.neighbor_address", "10.1.1.2"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.remote_as", "65002"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.description", "My Neighbor Description"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.use_neighbor_group", "GROUP1"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.ignore_connected_check", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.ebgp_multihop_maximum_hop_count", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.bfd_minimum_interval", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.bfd_multiplier", "4"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.local_as", "65003"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.local_as_no_prepend", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.local_as_replace_as", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.local_as_dual_as", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.password", "12341C2713181F13253920"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.shutdown", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.timers_keepalive_interval", "5"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.timers_holdtime", "20"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.update_source", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbors.0.ttl_security", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbor_groups.0.name", "GROUP1"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbor_groups.0.remote_as", "65001"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbor_groups.0.update_source", "Loopback0"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp.test", "neighbor_groups.0.bfd_minimum_interval", "3"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterBGPConfig = `

resource "iosxr_router_bgp" "test" {
	delete_mode = "attributes"
	as_number = "65001"
	default_information_originate = true
	default_metric = 125
	nsr_disable = false
	bgp_redistribute_internal = true
	segment_routing_srv6_locator = "locator11"
	timers_bgp_keepalive_interval = 5
	timers_bgp_holdtime = "20"
	bgp_router_id = "22.22.22.22"
	bgp_graceful_restart_graceful_reset = true
	ibgp_policy_out_enforce_modifications = true
	bgp_log_neighbor_changes_detail = true
	bfd_minimum_interval = 10
	bfd_multiplier = 4
	nexthop_validation_color_extcomm_sr_policy = true
	nexthop_validation_color_extcomm_disable = true
	neighbors = [{
		neighbor_address = "10.1.1.2"
		remote_as = "65002"
		description = "My Neighbor Description"
		use_neighbor_group = "GROUP1"
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
	}]
	neighbor_groups = [{
		name = "GROUP1"
		remote_as = "65001"
		update_source = "Loopback0"
		bfd_minimum_interval = 3
	}]
}

data "iosxr_router_bgp" "test" {
	as_number = "65001"
	depends_on = [iosxr_router_bgp.test]
}
`
