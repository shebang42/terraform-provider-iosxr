// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrCDP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrCDPConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_cdp.test", "enable", "true"),
					resource.TestCheckResourceAttr("iosxr_cdp.test", "holdtime", "12"),
					resource.TestCheckResourceAttr("iosxr_cdp.test", "timer", "34"),
					resource.TestCheckResourceAttr("iosxr_cdp.test", "advertise_v1", "true"),
					resource.TestCheckResourceAttr("iosxr_cdp.test", "log_adjacency_changes", "true"),
				),
			},
			{
				ResourceName:  "iosxr_cdp.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-cdp-cfg:/cdp",
			},
		},
	})
}

func testAccIosxrCDPConfig_minimum() string {
	return `
	resource "iosxr_cdp" "test" {
	}
	`
}

func testAccIosxrCDPConfig_all() string {
	return `
	resource "iosxr_cdp" "test" {
		enable = true
		holdtime = 12
		timer = 34
		advertise_v1 = true
		log_adjacency_changes = true
	}
	`
}