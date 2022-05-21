// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type RouterBGPVRFAddressFamily struct {
	Device                      types.String `tfsdk:"device"`
	Id                          types.String `tfsdk:"id"`
	AsNumber                    types.String `tfsdk:"as_number"`
	VrfName                     types.String `tfsdk:"vrf_name"`
	AfName                      types.String `tfsdk:"af_name"`
	MaximumPathsEbgpMultipath   types.Int64  `tfsdk:"maximum_paths_ebgp_multipath"`
	MaximumPathsEibgpMultipath  types.Int64  `tfsdk:"maximum_paths_eibgp_multipath"`
	MaximumPathsIbgpMultipath   types.Int64  `tfsdk:"maximum_paths_ibgp_multipath"`
	LabelModePerCe              types.Bool   `tfsdk:"label_mode_per_ce"`
	LabelModePerVrf             types.Bool   `tfsdk:"label_mode_per_vrf"`
	RedistributeConnected       types.Bool   `tfsdk:"redistribute_connected"`
	RedistributeConnectedMetric types.Int64  `tfsdk:"redistribute_connected_metric"`
	RedistributeStatic          types.Bool   `tfsdk:"redistribute_static"`
	RedistributeStaticMetric    types.Int64  `tfsdk:"redistribute_static_metric"`
}

func (data RouterBGPVRFAddressFamily) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]/vrfs/vrf[vrf-name=%s]/address-families/address-family[af-name=%s]", data.AsNumber.Value, data.VrfName.Value, data.AfName.Value)
}

func (data RouterBGPVRFAddressFamily) toBody() string {
	body := "{}"
	if !data.MaximumPathsEbgpMultipath.Null && !data.MaximumPathsEbgpMultipath.Unknown {
		body, _ = sjson.Set(body, "maximum-paths.ebgp.multipath", strconv.FormatInt(data.MaximumPathsEbgpMultipath.Value, 10))
	}
	if !data.MaximumPathsEibgpMultipath.Null && !data.MaximumPathsEibgpMultipath.Unknown {
		body, _ = sjson.Set(body, "maximum-paths.eibgp.multipath", strconv.FormatInt(data.MaximumPathsEibgpMultipath.Value, 10))
	}
	if !data.MaximumPathsIbgpMultipath.Null && !data.MaximumPathsIbgpMultipath.Unknown {
		body, _ = sjson.Set(body, "maximum-paths.ibgp.multipath", strconv.FormatInt(data.MaximumPathsIbgpMultipath.Value, 10))
	}
	if !data.LabelModePerCe.Null && !data.LabelModePerCe.Unknown {
		if data.LabelModePerCe.Value {
			body, _ = sjson.Set(body, "label.mode.per-ce", map[string]string{})
		}
	}
	if !data.LabelModePerVrf.Null && !data.LabelModePerVrf.Unknown {
		if data.LabelModePerVrf.Value {
			body, _ = sjson.Set(body, "label.mode.per-vrf", map[string]string{})
		}
	}
	if !data.RedistributeConnected.Null && !data.RedistributeConnected.Unknown {
		if data.RedistributeConnected.Value {
			body, _ = sjson.Set(body, "redistribute.connected", map[string]string{})
		}
	}
	if !data.RedistributeConnectedMetric.Null && !data.RedistributeConnectedMetric.Unknown {
		body, _ = sjson.Set(body, "redistribute.connected.metric", strconv.FormatInt(data.RedistributeConnectedMetric.Value, 10))
	}
	if !data.RedistributeStatic.Null && !data.RedistributeStatic.Unknown {
		if data.RedistributeStatic.Value {
			body, _ = sjson.Set(body, "redistribute.static", map[string]string{})
		}
	}
	if !data.RedistributeStaticMetric.Null && !data.RedistributeStaticMetric.Unknown {
		body, _ = sjson.Set(body, "redistribute.static.metric", strconv.FormatInt(data.RedistributeStaticMetric.Value, 10))
	}
	return body
}

func (data *RouterBGPVRFAddressFamily) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "maximum-paths.ebgp.multipath"); value.Exists() {
		data.MaximumPathsEbgpMultipath.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "maximum-paths.eibgp.multipath"); value.Exists() {
		data.MaximumPathsEibgpMultipath.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "maximum-paths.ibgp.multipath"); value.Exists() {
		data.MaximumPathsIbgpMultipath.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "label.mode.per-ce"); value.Exists() {
		data.LabelModePerCe.Value = true
	}
	if value := gjson.GetBytes(res, "label.mode.per-vrf"); value.Exists() {
		data.LabelModePerVrf.Value = true
	}
	if value := gjson.GetBytes(res, "redistribute.connected"); value.Exists() {
		data.RedistributeConnected.Value = true
	}
	if value := gjson.GetBytes(res, "redistribute.connected.metric"); value.Exists() {
		data.RedistributeConnectedMetric.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "redistribute.static"); value.Exists() {
		data.RedistributeStatic.Value = true
	}
	if value := gjson.GetBytes(res, "redistribute.static.metric"); value.Exists() {
		data.RedistributeStaticMetric.Value = value.Int()
	}
}

func (data *RouterBGPVRFAddressFamily) fromPlan(plan RouterBGPVRFAddressFamily) {
	data.Device = plan.Device
	data.AsNumber.Value = plan.AsNumber.Value
	data.VrfName.Value = plan.VrfName.Value
	data.AfName.Value = plan.AfName.Value
}
