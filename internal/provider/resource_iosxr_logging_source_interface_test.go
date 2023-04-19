// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrLoggingSourceInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrLoggingSourceInterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_logging_source_interface.test", "name", "Loopback0"),
					resource.TestCheckResourceAttr("iosxr_logging_source_interface.test", "vrfs.0.name", "VRF1"),
				),
			},
			{
				ResourceName:  "iosxr_logging_source_interface.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-logging-cfg:logging/source-interfaces/source-interface[source-interface-name=Loopback0]",
			},
		},
	})
}

func testAccIosxrLoggingSourceInterfaceConfig_minimum() string {
	return `
	resource "iosxr_logging_source_interface" "test" {
		name = "Loopback0"
	}
	`
}

func testAccIosxrLoggingSourceInterfaceConfig_all() string {
	return `
	resource "iosxr_logging_source_interface" "test" {
		name = "Loopback0"
		vrfs = [{
			name = "VRF1"
		}]
	}
	`
}
