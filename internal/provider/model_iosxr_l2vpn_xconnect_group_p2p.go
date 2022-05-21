// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type L2VPNXconnectGroupP2P struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	GroupName       types.String `tfsdk:"group_name"`
	P2pXconnectName types.String `tfsdk:"p2p_xconnect_name"`
	Description     types.String `tfsdk:"description"`
}

func (data L2VPNXconnectGroupP2P) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=%s]/p2ps/p2p[p2p-xconnect-name=%s]", data.GroupName.Value, data.P2pXconnectName.Value)
}

func (data L2VPNXconnectGroupP2P) toBody() string {
	body := "{}"
	if !data.Description.Null && !data.Description.Unknown {
		body, _ = sjson.Set(body, "description", data.Description.Value)
	}
	return body
}

func (data *L2VPNXconnectGroupP2P) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() {
		data.Description.Value = value.String()
	}
}

func (data *L2VPNXconnectGroupP2P) fromPlan(plan L2VPNXconnectGroupP2P) {
	data.Device = plan.Device
	data.GroupName.Value = plan.GroupName.Value
	data.P2pXconnectName.Value = plan.P2pXconnectName.Value
}
