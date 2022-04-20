// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrMPLSLDP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrMPLSLDPConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_mpls_ldp.test", "router_id", "1.2.3.4"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrMPLSLDPConfig = `
resource "iosxr_mpls_ldp" "test" {
  router_id = "1.2.3.4"
}

data "iosxr_mpls_ldp" "test" {
  depends_on = [iosxr_mpls_ldp.test]
}
`
