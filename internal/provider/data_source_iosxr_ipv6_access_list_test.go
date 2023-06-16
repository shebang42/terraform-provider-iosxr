// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrIPv6AccessList(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrIPv6AccessListConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.sequence_number", "22"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.remark", "remark for access list"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_protocol", "tcp"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_source_address", "1::1"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_source_prefix_length", "64"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_source_port_range_start_value", "100"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_source_port_range_end_value", "200"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_destination_host", "2::1"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_destination_port_eq", "10"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_nexthop1_ipv6", "3::3"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_nexthop2_ipv6", "4::4"),
					resource.TestCheckResourceAttr("data.iosxr_ipv6_access_list.test", "sequences.0.permit_log", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrIPv6AccessListConfig = `

resource "iosxr_ipv6_access_list" "test" {
	access_list_name = "TEST1"
	sequences = [{
		sequence_number = 22
		remark = "remark for access list"
		permit_protocol = "tcp"
		permit_source_address = "1::1"
		permit_source_prefix_length = 64
		permit_source_port_range_start_value = "100"
		permit_source_port_range_end_value = "200"
		permit_destination_host = "2::1"
		permit_destination_port_eq = "10"
		permit_nexthop1_ipv6 = "3::3"
		permit_nexthop2_ipv6 = "4::4"
		permit_log = true
	}]
}

data "iosxr_ipv6_access_list" "test" {
	access_list_name = "TEST1"
	depends_on = [iosxr_ipv6_access_list.test]
}
`
