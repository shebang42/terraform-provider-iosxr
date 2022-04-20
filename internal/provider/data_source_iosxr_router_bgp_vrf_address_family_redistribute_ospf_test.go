// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterBGPVRFAddressFamilyRedistributeOSPF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterBGPVRFAddressFamilyRedistributeOSPFPrerequisitesConfig + testAccDataSourceIosxrRouterBGPVRFAddressFamilyRedistributeOSPFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family_redistribute_ospf.test", "match_internal", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family_redistribute_ospf.test", "match_internal_external", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family_redistribute_ospf.test", "match_internal_nssa_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family_redistribute_ospf.test", "match_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family_redistribute_ospf.test", "match_external_nssa_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family_redistribute_ospf.test", "match_nssa_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family_redistribute_ospf.test", "metric", "100"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterBGPVRFAddressFamilyRedistributeOSPFPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
  path = "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]"
  attributes = {
  }
}

resource "iosxr_gnmi" "PreReq1" {
  path = "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]/Cisco-IOS-XR-um-router-bgp-cfg:rd/Cisco-IOS-XR-um-router-bgp-cfg:two-byte-as"
  attributes = {
      as-number = "1"
      index = "1"
  }
  depends_on = [iosxr_gnmi.PreReq0, ]
}

resource "iosxr_gnmi" "PreReq2" {
  path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]"
  attributes = {
  }
}

resource "iosxr_gnmi" "PreReq3" {
  path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/address-families/address-family[af-name=vpnv4-unicast]"
  attributes = {
  }
  depends_on = [iosxr_gnmi.PreReq1, iosxr_gnmi.PreReq2, ]
}

resource "iosxr_gnmi" "PreReq4" {
  path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/vrfs/vrf[vrf-name=VRF1]/address-families/address-family[af-name=ipv4-unicast]"
  attributes = {
  }
  depends_on = [iosxr_gnmi.PreReq3, ]
}

`

const testAccDataSourceIosxrRouterBGPVRFAddressFamilyRedistributeOSPFConfig = `
resource "iosxr_router_bgp_vrf_address_family_redistribute_ospf" "test" {
  as_number = "65001"
  vrf_name = "VRF1"
  af_name = "ipv4-unicast"
  router_tag = "OSPF1"
  match_internal = true
  match_internal_external = true
  match_internal_nssa_external = false
  match_external = false
  match_external_nssa_external = false
  match_nssa_external = false
  metric = 100
  depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, iosxr_gnmi.PreReq2, iosxr_gnmi.PreReq3, iosxr_gnmi.PreReq4, ]
}

data "iosxr_router_bgp_vrf_address_family_redistribute_ospf" "test" {
  as_number = "65001"
  vrf_name = "VRF1"
  af_name = "ipv4-unicast"
  router_tag = "OSPF1"
  depends_on = [iosxr_router_bgp_vrf_address_family_redistribute_ospf.test]
}
`
