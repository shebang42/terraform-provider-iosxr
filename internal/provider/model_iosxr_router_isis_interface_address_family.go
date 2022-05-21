// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RouterISISInterfaceAddressFamily struct {
	Device        types.String `tfsdk:"device"`
	Id            types.String `tfsdk:"id"`
	ProcessId     types.String `tfsdk:"process_id"`
	InterfaceName types.String `tfsdk:"interface_name"`
	AfName        types.String `tfsdk:"af_name"`
	SafName       types.String `tfsdk:"saf_name"`
}

func (data RouterISISInterfaceAddressFamily) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=%s]/interfaces/interface[interface-name=%s]/address-families/address-family[af-name=%s][saf-name=%s]", data.ProcessId.Value, data.InterfaceName.Value, data.AfName.Value, data.SafName.Value)
}

func (data RouterISISInterfaceAddressFamily) toBody() string {
	body := "{}"
	return body
}

func (data *RouterISISInterfaceAddressFamily) fromBody(res []byte) {
}

func (data *RouterISISInterfaceAddressFamily) fromPlan(plan RouterISISInterfaceAddressFamily) {
	data.Device = plan.Device
	data.ProcessId.Value = plan.ProcessId.Value
	data.InterfaceName.Value = plan.InterfaceName.Value
	data.AfName.Value = plan.AfName.Value
	data.SafName.Value = plan.SafName.Value
}
