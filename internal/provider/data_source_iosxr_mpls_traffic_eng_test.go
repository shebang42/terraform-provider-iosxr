// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrMPLSTrafficEng(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrMPLSTrafficEngConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_mpls_traffic_eng.test", "traffic_eng", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrMPLSTrafficEngConfig = `

resource "iosxr_mpls_traffic_eng" "test" {
	delete_mode = "attributes"
	traffic_eng = true
}

data "iosxr_mpls_traffic_eng" "test" {
	depends_on = [iosxr_mpls_traffic_eng.test]
}
`
