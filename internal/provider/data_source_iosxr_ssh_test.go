//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrSSH(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrSSHConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_ssh.test", "server_dscp", "48"),
					resource.TestCheckResourceAttr("data.iosxr_ssh.test", "server_logging", "true"),
					resource.TestCheckResourceAttr("data.iosxr_ssh.test", "server_rate_limit", "60"),
					resource.TestCheckResourceAttr("data.iosxr_ssh.test", "server_session_limit", "10"),
					resource.TestCheckResourceAttr("data.iosxr_ssh.test", "server_v2", "true"),
					resource.TestCheckResourceAttr("data.iosxr_ssh.test", "server_vrfs.0.vrf_name", "VRF1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrSSHConfig = `

resource "iosxr_ssh" "test" {
	delete_mode = "attributes"
	server_dscp = 48
	server_logging = true
	server_rate_limit = 60
	server_session_limit = 10
	server_v2 = true
	server_vrfs = [{
		vrf_name = "VRF1"
	}]
}

data "iosxr_ssh" "test" {
	depends_on = [iosxr_ssh.test]
}
`
