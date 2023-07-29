// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterBGPNeighborGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "name", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "remote_as", "65001"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "update_source", "Loopback0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "bfd_minimum_interval", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "bfd_multiplier", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "bfd_fast_detect", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "bfd_fast_detect_strict_mode", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "bfd_fast_detect_inheritance_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "address_families.0.af_name", "ipv4-labeled-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "address_families.0.soft_reconfiguration_inbound_always", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "address_families.0.next_hop_self_inheritance_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_group.test", "address_families.0.route_reflector_client_inheritance_disable", "true"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrRouterBGPNeighborGroupPrerequisitesConfig + testAccIosxrRouterBGPNeighborGroupConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrRouterBGPNeighborGroupPrerequisitesConfig + testAccIosxrRouterBGPNeighborGroupConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_router_bgp_neighbor_group.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/neighbor-groups/neighbor-group[neighbor-group-name=GROUP1]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

const testAccIosxrRouterBGPNeighborGroupPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]"
	attributes = {
		"as-number" = "65001"
	}
}

`

func testAccIosxrRouterBGPNeighborGroupConfig_minimum() string {
	config := `resource "iosxr_router_bgp_neighbor_group" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	name = "GROUP1"` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterBGPNeighborGroupConfig_all() string {
	config := `resource "iosxr_router_bgp_neighbor_group" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	name = "GROUP1"` + "\n"
	config += `	remote_as = "65001"` + "\n"
	config += `	update_source = "Loopback0"` + "\n"
	config += `	bfd_minimum_interval = 3` + "\n"
	config += `	bfd_multiplier = 4` + "\n"
	config += `	bfd_fast_detect = true` + "\n"
	config += `	bfd_fast_detect_strict_mode = false` + "\n"
	config += `	bfd_fast_detect_inheritance_disable = false` + "\n"
	config += `	address_families = [{` + "\n"
	config += `		af_name = "ipv4-labeled-unicast"` + "\n"
	config += `		soft_reconfiguration_inbound_always = true` + "\n"
	config += `		next_hop_self_inheritance_disable = true` + "\n"
	config += `		route_reflector_client_inheritance_disable = true` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
