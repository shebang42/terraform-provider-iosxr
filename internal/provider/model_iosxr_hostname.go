// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type Hostname struct {
	Device            types.String `tfsdk:"device"`
	Id                types.String `tfsdk:"id"`
	SystemNetworkName types.String `tfsdk:"system_network_name"`
}

func (data Hostname) getPath() string {
	return "Cisco-IOS-XR-um-hostname-cfg:/hostname"
}

func (data Hostname) toBody() string {
	body := "{}"

	if !data.SystemNetworkName.Null && !data.SystemNetworkName.Unknown {
		body, _ = sjson.Set(body, "system-network-name", data.SystemNetworkName.Value)
	}

	return body
}

func (data *Hostname) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "system-network-name"); value.Exists() {
		data.SystemNetworkName.Value = value.String()
	}
}

func (data *Hostname) fromPlan(plan Hostname) {
	data.Device = plan.Device
}
