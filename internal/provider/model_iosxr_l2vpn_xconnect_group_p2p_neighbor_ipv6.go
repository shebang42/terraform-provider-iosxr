// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type L2VPNXconnectGroupP2PNeighborIPv6 struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	GroupName       types.String `tfsdk:"group_name"`
	P2pXconnectName types.String `tfsdk:"p2p_xconnect_name"`
	Address         types.String `tfsdk:"address"`
	PwId            types.Int64  `tfsdk:"pw_id"`
	PwClass         types.String `tfsdk:"pw_class"`
}

func (data L2VPNXconnectGroupP2PNeighborIPv6) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=%s]/p2ps/p2p[p2p-xconnect-name=%s]/neighbor/ipv6s/ipv6[address=%s][pw-id=%v]", data.GroupName.Value, data.P2pXconnectName.Value, data.Address.Value, data.PwId.Value)
}

func (data L2VPNXconnectGroupP2PNeighborIPv6) toBody() string {
	body := "{}"

	if !data.PwClass.Null && !data.PwClass.Unknown {
		body, _ = sjson.Set(body, "pw-class", data.PwClass.Value)
	}

	return body
}

func (data *L2VPNXconnectGroupP2PNeighborIPv6) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "pw-class"); value.Exists() {
		data.PwClass.Value = value.String()
	}
}

func (data *L2VPNXconnectGroupP2PNeighborIPv6) fromPlan(plan L2VPNXconnectGroupP2PNeighborIPv6) {
	data.Device = plan.Device
	data.GroupName.Value = plan.GroupName.Value
	data.P2pXconnectName.Value = plan.P2pXconnectName.Value
	data.Address.Value = plan.Address.Value
	data.PwId.Value = plan.PwId.Value
}
