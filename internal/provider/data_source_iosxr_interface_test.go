// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "l2transport", "false"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "point_to_point", "false"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "multipoint", "false"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "shutdown", "true"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "mtu", "9000"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "bandwidth", "100000"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "description", "My Interface Description"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "vrf", "VRF1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrInterfaceConfig = `
resource "iosxr_interface" "test" {
  interface_name = "GigabitEthernet0/0/0/1"
  l2transport = false
  point_to_point = false
  multipoint = false
  shutdown = true
  mtu = 9000
  bandwidth = 100000
  description = "My Interface Description"
  vrf = "VRF1"
}

data "iosxr_interface" "test" {
  interface_name = "GigabitEthernet0/0/0/1"
  depends_on = [iosxr_interface.test]
}
`
