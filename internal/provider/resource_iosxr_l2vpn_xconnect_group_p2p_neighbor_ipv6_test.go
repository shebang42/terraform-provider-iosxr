//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrL2VPNXconnectGroupP2PNeighborIPv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrL2VPNXconnectGroupP2PNeighborIPv6Config_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv6.test", "address", "2001::2"),
					resource.TestCheckResourceAttr("iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv6.test", "pw_id", "2"),
					resource.TestCheckResourceAttr("iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv6.test", "pw_class", "PW_CLASS_1"),
				),
			},
			{
				ResourceName:  "iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv6.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=P2P]/p2ps/p2p[p2p-xconnect-name=XC]/neighbor/ipv6s/ipv6[address=2001::2][pw-id=2]",
			},
		},
	})
}

func testAccIosxrL2VPNXconnectGroupP2PNeighborIPv6Config_minimum() string {
	return `
	resource "iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv6" "test" {
		group_name = "P2P"
		p2p_xconnect_name = "XC"
		address = "2001::2"
		pw_id = 2
	}
	`
}

func testAccIosxrL2VPNXconnectGroupP2PNeighborIPv6Config_all() string {
	return `
	resource "iosxr_l2vpn_xconnect_group_p2p_neighbor_ipv6" "test" {
		group_name = "P2P"
		p2p_xconnect_name = "XC"
		address = "2001::2"
		pw_id = 2
		pw_class = "PW_CLASS_1"
	}
	`
}
