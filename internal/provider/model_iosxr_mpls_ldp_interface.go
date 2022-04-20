// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MPLSLDPInterface struct {
	Device        types.String `tfsdk:"device"`
	Id            types.String `tfsdk:"id"`
	InterfaceName types.String `tfsdk:"interface_name"`
}

func (data MPLSLDPInterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-mpls-ldp-cfg:/mpls/ldp/interfaces/interface[interface-name=%s]", data.InterfaceName.Value)
}

func (data MPLSLDPInterface) toBody() string {
	body := "{}"

	return body
}

func (data *MPLSLDPInterface) fromBody(res []byte) {
}

func (data *MPLSLDPInterface) fromPlan(plan MPLSLDPInterface) {
	data.Device = plan.Device
	data.InterfaceName.Value = plan.InterfaceName.Value
}
