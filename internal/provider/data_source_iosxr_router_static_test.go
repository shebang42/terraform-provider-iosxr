// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterStatic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterStaticConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_interfaces.0.interface_name", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_interfaces.0.description", "interface-description"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_interfaces.0.tag", "100"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_interfaces.0.distance_metric", "122"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_interface_addresses.0.interface_name", "GigabitEthernet0/0/0/2"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_interface_addresses.0.address", "11.11.11.1"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_interface_addresses.0.description", "interface-description"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_interface_addresses.0.tag", "103"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_interface_addresses.0.distance_metric", "144"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_addresses.0.address", "100.0.2.0"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_addresses.0.description", "ip-description"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_addresses.0.tag", "104"),
					resource.TestCheckResourceAttr("data.iosxr_router_static.test", "nexthop_addresses.0.distance_metric", "155"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterStaticConfig = `

resource "iosxr_router_static" "test" {
	delete_mode = "attributes"
	prefix_address = "100.0.1.0"
	prefix_length = 24
	nexthop_interfaces = [{
		interface_name = "GigabitEthernet0/0/0/1"
		description = "interface-description"
		tag = 100
		distance_metric = 122
	}]
	nexthop_interface_addresses = [{
		interface_name = "GigabitEthernet0/0/0/2"
		address = "11.11.11.1"
		description = "interface-description"
		tag = 103
		distance_metric = 144
	}]
	nexthop_addresses = [{
		address = "100.0.2.0"
		description = "ip-description"
		tag = 104
		distance_metric = 155
	}]
}

data "iosxr_router_static" "test" {
	prefix_address = "100.0.1.0"
	prefix_length = 24
	depends_on = [iosxr_router_static.test]
}
`
