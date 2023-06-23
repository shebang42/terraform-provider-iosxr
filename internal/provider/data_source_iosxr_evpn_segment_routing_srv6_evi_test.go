// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrEVPNSegmentRoutingSRv6EVI(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrEVPNSegmentRoutingSRv6EVIConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_evpn_segment_routing_srv6_evi.test", "description", "My Description"),
					resource.TestCheckResourceAttr("data.iosxr_evpn_segment_routing_srv6_evi.test", "bgp_route_target_import_two_byte_as_format.0.as_number", "1"),
					resource.TestCheckResourceAttr("data.iosxr_evpn_segment_routing_srv6_evi.test", "bgp_route_target_import_two_byte_as_format.0.assigned_number", "1"),
					resource.TestCheckResourceAttr("data.iosxr_evpn_segment_routing_srv6_evi.test", "bgp_route_target_export_ipv4_address_format.0.ipv4_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("data.iosxr_evpn_segment_routing_srv6_evi.test", "bgp_route_target_export_ipv4_address_format.0.assigned_number", "1"),
					resource.TestCheckResourceAttr("data.iosxr_evpn_segment_routing_srv6_evi.test", "advertise_mac", "true"),
					resource.TestCheckResourceAttr("data.iosxr_evpn_segment_routing_srv6_evi.test", "locator", "LOC12"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrEVPNSegmentRoutingSRv6EVIConfig = `

resource "iosxr_evpn_segment_routing_srv6_evi" "test" {
	vpn_id = 1235
	description = "My Description"
	bgp_route_target_import_two_byte_as_format = [{
		as_number = 1
		assigned_number = 1
	}]
	bgp_route_target_export_ipv4_address_format = [{
		ipv4_address = "1.1.1.1"
		assigned_number = 1
	}]
	advertise_mac = true
	locator = "LOC12"
}

data "iosxr_evpn_segment_routing_srv6_evi" "test" {
	vpn_id = 1235
	depends_on = [iosxr_evpn_segment_routing_srv6_evi.test]
}
`
