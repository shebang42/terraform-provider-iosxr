// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type RouterBGPVRFAddressFamilyRedistributeOSPF struct {
	Device                    types.String `tfsdk:"device"`
	Id                        types.String `tfsdk:"id"`
	AsNumber                  types.String `tfsdk:"as_number"`
	VrfName                   types.String `tfsdk:"vrf_name"`
	AfName                    types.String `tfsdk:"af_name"`
	RouterTag                 types.String `tfsdk:"router_tag"`
	MatchInternal             types.Bool   `tfsdk:"match_internal"`
	MatchInternalExternal     types.Bool   `tfsdk:"match_internal_external"`
	MatchInternalNssaExternal types.Bool   `tfsdk:"match_internal_nssa_external"`
	MatchExternal             types.Bool   `tfsdk:"match_external"`
	MatchExternalNssaExternal types.Bool   `tfsdk:"match_external_nssa_external"`
	MatchNssaExternal         types.Bool   `tfsdk:"match_nssa_external"`
	Metric                    types.Int64  `tfsdk:"metric"`
}

func (data RouterBGPVRFAddressFamilyRedistributeOSPF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]/vrfs/vrf[vrf-name=%s]/address-families/address-family[af-name=%s]/redistribute/ospf[router-tag=%s]", data.AsNumber.Value, data.VrfName.Value, data.AfName.Value, data.RouterTag.Value)
}

func (data RouterBGPVRFAddressFamilyRedistributeOSPF) toBody() string {
	body := "{}"
	if !data.MatchInternal.Null && !data.MatchInternal.Unknown {
		if data.MatchInternal.Value {
			body, _ = sjson.Set(body, "match.internal", map[string]string{})
		}
	}
	if !data.MatchInternalExternal.Null && !data.MatchInternalExternal.Unknown {
		if data.MatchInternalExternal.Value {
			body, _ = sjson.Set(body, "match.internal.external", map[string]string{})
		}
	}
	if !data.MatchInternalNssaExternal.Null && !data.MatchInternalNssaExternal.Unknown {
		if data.MatchInternalNssaExternal.Value {
			body, _ = sjson.Set(body, "match.internal.nssa-external", map[string]string{})
		}
	}
	if !data.MatchExternal.Null && !data.MatchExternal.Unknown {
		if data.MatchExternal.Value {
			body, _ = sjson.Set(body, "match.external", map[string]string{})
		}
	}
	if !data.MatchExternalNssaExternal.Null && !data.MatchExternalNssaExternal.Unknown {
		if data.MatchExternalNssaExternal.Value {
			body, _ = sjson.Set(body, "match.external.nssa-external", map[string]string{})
		}
	}
	if !data.MatchNssaExternal.Null && !data.MatchNssaExternal.Unknown {
		if data.MatchNssaExternal.Value {
			body, _ = sjson.Set(body, "match.nssa-external", map[string]string{})
		}
	}
	if !data.Metric.Null && !data.Metric.Unknown {
		body, _ = sjson.Set(body, "metric", strconv.FormatInt(data.Metric.Value, 10))
	}
	return body
}

func (data *RouterBGPVRFAddressFamilyRedistributeOSPF) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "match.internal"); value.Exists() {
		data.MatchInternal.Value = true
	}
	if value := gjson.GetBytes(res, "match.internal.external"); value.Exists() {
		data.MatchInternalExternal.Value = true
	}
	if value := gjson.GetBytes(res, "match.internal.nssa-external"); value.Exists() {
		data.MatchInternalNssaExternal.Value = true
	}
	if value := gjson.GetBytes(res, "match.external"); value.Exists() {
		data.MatchExternal.Value = true
	}
	if value := gjson.GetBytes(res, "match.external.nssa-external"); value.Exists() {
		data.MatchExternalNssaExternal.Value = true
	}
	if value := gjson.GetBytes(res, "match.nssa-external"); value.Exists() {
		data.MatchNssaExternal.Value = true
	}
	if value := gjson.GetBytes(res, "metric"); value.Exists() {
		data.Metric.Value = value.Int()
	}
}

func (data *RouterBGPVRFAddressFamilyRedistributeOSPF) fromPlan(plan RouterBGPVRFAddressFamilyRedistributeOSPF) {
	data.Device = plan.Device
	data.AsNumber.Value = plan.AsNumber.Value
	data.VrfName.Value = plan.VrfName.Value
	data.AfName.Value = plan.AfName.Value
	data.RouterTag.Value = plan.RouterTag.Value
}
