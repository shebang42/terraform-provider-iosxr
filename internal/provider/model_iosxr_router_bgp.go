// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type RouterBGP struct {
	Device                      types.String `tfsdk:"device"`
	Id                          types.String `tfsdk:"id"`
	AsNumber                    types.String `tfsdk:"as_number"`
	DefaultInformationOriginate types.Bool   `tfsdk:"default_information_originate"`
	DefaultMetric               types.Int64  `tfsdk:"default_metric"`
	TimersBgpKeepaliveInterval  types.Int64  `tfsdk:"timers_bgp_keepalive_interval"`
	TimersBgpHoldtime           types.String `tfsdk:"timers_bgp_holdtime"`
	BfdMinimumInterval          types.Int64  `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier               types.Int64  `tfsdk:"bfd_multiplier"`
}

func (data RouterBGP) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]", data.AsNumber.Value)
}

func (data RouterBGP) toBody() string {
	body := "{}"
	if !data.DefaultInformationOriginate.Null && !data.DefaultInformationOriginate.Unknown {
		if data.DefaultInformationOriginate.Value {
			body, _ = sjson.Set(body, "default-information.originate", map[string]string{})
		}
	}
	if !data.DefaultMetric.Null && !data.DefaultMetric.Unknown {
		body, _ = sjson.Set(body, "default-metric", strconv.FormatInt(data.DefaultMetric.Value, 10))
	}
	if !data.TimersBgpKeepaliveInterval.Null && !data.TimersBgpKeepaliveInterval.Unknown {
		body, _ = sjson.Set(body, "timers.bgp.keepalive-interval", strconv.FormatInt(data.TimersBgpKeepaliveInterval.Value, 10))
	}
	if !data.TimersBgpHoldtime.Null && !data.TimersBgpHoldtime.Unknown {
		body, _ = sjson.Set(body, "timers.bgp.holdtime", data.TimersBgpHoldtime.Value)
	}
	if !data.BfdMinimumInterval.Null && !data.BfdMinimumInterval.Unknown {
		body, _ = sjson.Set(body, "bfd.minimum-interval", strconv.FormatInt(data.BfdMinimumInterval.Value, 10))
	}
	if !data.BfdMultiplier.Null && !data.BfdMultiplier.Unknown {
		body, _ = sjson.Set(body, "bfd.multiplier", strconv.FormatInt(data.BfdMultiplier.Value, 10))
	}
	return body
}

func (data *RouterBGP) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "default-information.originate"); value.Exists() {
		data.DefaultInformationOriginate.Value = true
	}
	if value := gjson.GetBytes(res, "default-metric"); value.Exists() {
		data.DefaultMetric.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "timers.bgp.keepalive-interval"); value.Exists() {
		data.TimersBgpKeepaliveInterval.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "timers.bgp.holdtime"); value.Exists() {
		data.TimersBgpHoldtime.Value = value.String()
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() {
		data.BfdMinimumInterval.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() {
		data.BfdMultiplier.Value = value.Int()
	}
}

func (data *RouterBGP) fromPlan(plan RouterBGP) {
	data.Device = plan.Device
	data.AsNumber.Value = plan.AsNumber.Value
}
