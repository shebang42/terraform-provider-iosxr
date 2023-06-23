// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrSNMPServerView(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrSNMPServerViewConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_snmp_server_view.test", "mib_view_families.0.name", "iso"),
					resource.TestCheckResourceAttr("data.iosxr_snmp_server_view.test", "mib_view_families.0.included", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrSNMPServerViewConfig = `

resource "iosxr_snmp_server_view" "test" {
	delete_mode = "attributes"
	view_name = "VIEW12"
	mib_view_families = [{
		name = "iso"
		included = true
	}]
}

data "iosxr_snmp_server_view" "test" {
	view_name = "VIEW12"
	depends_on = [iosxr_snmp_server_view.test]
}
`
