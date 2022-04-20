// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type RouterOSPFRedistributeOSPF struct {
	Device            types.String `tfsdk:"device"`
	Id                types.String `tfsdk:"id"`
	ProcessName       types.String `tfsdk:"process_name"`
	InstanceName      types.String `tfsdk:"instance_name"`
	MatchInternal     types.Bool   `tfsdk:"match_internal"`
	MatchExternal     types.Bool   `tfsdk:"match_external"`
	MatchNssaExternal types.Bool   `tfsdk:"match_nssa_external"`
	Tag               types.Int64  `tfsdk:"tag"`
	MetricType        types.String `tfsdk:"metric_type"`
}

func (data RouterOSPFRedistributeOSPF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-ospf-cfg:/router/ospf/processes/process[process-name=%s]/redistribute/ospf[instance-name=%s]", data.ProcessName.Value, data.InstanceName.Value)
}

func (data RouterOSPFRedistributeOSPF) toBody() string {
	body := "{}"

	if !data.MatchInternal.Null && !data.MatchInternal.Unknown {
		if data.MatchInternal.Value {
			body, _ = sjson.Set(body, "match.internal", map[string]string{})
		}
	}
	if !data.MatchExternal.Null && !data.MatchExternal.Unknown {
		if data.MatchExternal.Value {
			body, _ = sjson.Set(body, "match.external", map[string]string{})
		}
	}
	if !data.MatchNssaExternal.Null && !data.MatchNssaExternal.Unknown {
		if data.MatchNssaExternal.Value {
			body, _ = sjson.Set(body, "match.nssa-external", map[string]string{})
		}
	}
	if !data.Tag.Null && !data.Tag.Unknown {
		body, _ = sjson.Set(body, "tag", strconv.FormatInt(data.Tag.Value, 10))
	}
	if !data.MetricType.Null && !data.MetricType.Unknown {
		body, _ = sjson.Set(body, "metric-type", data.MetricType.Value)
	}

	return body
}

func (data *RouterOSPFRedistributeOSPF) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "match.internal"); value.Exists() {
		data.MatchInternal.Value = true
	}
	if value := gjson.GetBytes(res, "match.external"); value.Exists() {
		data.MatchExternal.Value = true
	}
	if value := gjson.GetBytes(res, "match.nssa-external"); value.Exists() {
		data.MatchNssaExternal.Value = true
	}
	if value := gjson.GetBytes(res, "tag"); value.Exists() {
		data.Tag.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "metric-type"); value.Exists() {
		data.MetricType.Value = value.String()
	}
}

func (data *RouterOSPFRedistributeOSPF) fromPlan(plan RouterOSPFRedistributeOSPF) {
	data.Device = plan.Device
	data.ProcessName.Value = plan.ProcessName.Value
	data.InstanceName.Value = plan.InstanceName.Value
}
