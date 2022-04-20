// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type L2VPNXconnectGroupP2PInterface struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	GroupName       types.String `tfsdk:"group_name"`
	P2pXconnectName types.String `tfsdk:"p2p_xconnect_name"`
	InterfaceName   types.String `tfsdk:"interface_name"`
}

func (data L2VPNXconnectGroupP2PInterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=%s]/p2ps/p2p[p2p-xconnect-name=%s]/interfaces/interface[interface-name=%s]", data.GroupName.Value, data.P2pXconnectName.Value, data.InterfaceName.Value)
}

func (data L2VPNXconnectGroupP2PInterface) toBody() string {
	body := "{}"

	return body
}

func (data *L2VPNXconnectGroupP2PInterface) fromBody(res []byte) {
}

func (data *L2VPNXconnectGroupP2PInterface) fromPlan(plan L2VPNXconnectGroupP2PInterface) {
	data.Device = plan.Device
	data.GroupName.Value = plan.GroupName.Value
	data.P2pXconnectName.Value = plan.P2pXconnectName.Value
	data.InterfaceName.Value = plan.InterfaceName.Value
}
