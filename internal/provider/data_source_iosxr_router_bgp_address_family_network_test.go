// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterBGPAddressFamilyNetwork(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterBGPAddressFamilyNetworkConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDataSourceIosxrRouterBGPAddressFamilyNetworkConfig = `
resource "iosxr_router_bgp_address_family_network" "test" {
  as_number = "65001"
  af_name = "ipv4-unicast"
  address = "10.1.0.0"
  masklength = 16
}

data "iosxr_router_bgp_address_family_network" "test" {
  as_number = "65001"
  af_name = "ipv4-unicast"
  address = "10.1.0.0"
  masklength = 16
  depends_on = [iosxr_router_bgp_address_family_network.test]
}
`
