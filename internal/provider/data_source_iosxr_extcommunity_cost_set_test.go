// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrExtcommunityCostSet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrExtcommunityCostSetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_extcommunity_cost_set.test", "rpl", "extcommunity-set cost COST2\nend-set\n"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrExtcommunityCostSetConfig = `

resource "iosxr_extcommunity_cost_set" "test" {
	set_name = "COST2"
	rpl = "extcommunity-set cost COST2\nend-set\n"
}

data "iosxr_extcommunity_cost_set" "test" {
	set_name = "COST2"
	depends_on = [iosxr_extcommunity_cost_set.test]
}
`
