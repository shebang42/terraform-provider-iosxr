// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
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
	if !data.SystemNetworkName.IsNull() && !data.SystemNetworkName.IsUnknown() {
		body, _ = sjson.Set(body, "system-network-name", data.SystemNetworkName.ValueString())
	}
	return body
}

func (data *Hostname) updateFromBody(res []byte) {
	if value := gjson.GetBytes(res, "system-network-name"); value.Exists() && !data.SystemNetworkName.IsNull() {
		data.SystemNetworkName = types.StringValue(value.String())
	} else {
		data.SystemNetworkName = types.StringNull()
	}
}

func (data *Hostname) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "system-network-name"); value.Exists() {
		data.SystemNetworkName = types.StringValue(value.String())
	}
}

func (data *Hostname) fromPlan(plan Hostname) {
	data.Device = plan.Device
}

func (data *Hostname) getDeletedListItems(state Hostname) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *Hostname) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
