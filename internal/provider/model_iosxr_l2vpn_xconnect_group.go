// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type L2VPNXconnectGroup struct {
	Device    types.String `tfsdk:"device"`
	Id        types.String `tfsdk:"id"`
	GroupName types.String `tfsdk:"group_name"`
}

func (data L2VPNXconnectGroup) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=%s]", data.GroupName.Value)
}

func (data L2VPNXconnectGroup) toBody() string {
	body := "{}"

	return body
}

func (data *L2VPNXconnectGroup) fromBody(res []byte) {
}

func (data *L2VPNXconnectGroup) fromPlan(plan L2VPNXconnectGroup) {
	data.Device = plan.Device
	data.GroupName.Value = plan.GroupName.Value
}
