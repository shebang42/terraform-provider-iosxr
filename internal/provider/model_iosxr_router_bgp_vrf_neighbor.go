// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type RouterBGPVRFNeighbor struct {
	Device                          types.String `tfsdk:"device"`
	Id                              types.String `tfsdk:"id"`
	AsNumber                        types.String `tfsdk:"as_number"`
	VrfName                         types.String `tfsdk:"vrf_name"`
	NeighborAddress                 types.String `tfsdk:"neighbor_address"`
	RemoteAs                        types.String `tfsdk:"remote_as"`
	Description                     types.String `tfsdk:"description"`
	IgnoreConnectedCheck            types.Bool   `tfsdk:"ignore_connected_check"`
	EbgpMultihopMaximumHopCount     types.Int64  `tfsdk:"ebgp_multihop_maximum_hop_count"`
	BfdMinimumInterval              types.Int64  `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier                   types.Int64  `tfsdk:"bfd_multiplier"`
	LocalAsAsNumber                 types.String `tfsdk:"local_as"`
	LocalAsNoPrepend                types.Bool   `tfsdk:"local_as_no_prepend"`
	LocalAsNoPrependReplaceAs       types.Bool   `tfsdk:"local_as_replace_as"`
	LocalAsNoPrependReplaceAsDualAs types.Bool   `tfsdk:"local_as_dual_as"`
	PasswordEncrypted               types.String `tfsdk:"password"`
	Shutdown                        types.Bool   `tfsdk:"shutdown"`
	TimersKeepaliveInterval         types.Int64  `tfsdk:"timers_keepalive_interval"`
	TimersHoldtime                  types.String `tfsdk:"timers_holdtime"`
	UpdateSource                    types.String `tfsdk:"update_source"`
	TtlSecurity                     types.Bool   `tfsdk:"ttl_security"`
}

func (data RouterBGPVRFNeighbor) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]/vrfs/vrf[vrf-name=%s]/neighbors/neighbor[neighbor-address=%s]", data.AsNumber.Value, data.VrfName.Value, data.NeighborAddress.Value)
}

func (data RouterBGPVRFNeighbor) toBody() string {
	body := "{}"

	if !data.RemoteAs.Null && !data.RemoteAs.Unknown {
		body, _ = sjson.Set(body, "remote-as", data.RemoteAs.Value)
	}
	if !data.Description.Null && !data.Description.Unknown {
		body, _ = sjson.Set(body, "description", data.Description.Value)
	}
	if !data.IgnoreConnectedCheck.Null && !data.IgnoreConnectedCheck.Unknown {
		if data.IgnoreConnectedCheck.Value {
			body, _ = sjson.Set(body, "ignore-connected-check", map[string]string{})
		}
	}
	if !data.EbgpMultihopMaximumHopCount.Null && !data.EbgpMultihopMaximumHopCount.Unknown {
		body, _ = sjson.Set(body, "ebgp-multihop.maximum-hop-count", strconv.FormatInt(data.EbgpMultihopMaximumHopCount.Value, 10))
	}
	if !data.BfdMinimumInterval.Null && !data.BfdMinimumInterval.Unknown {
		body, _ = sjson.Set(body, "bfd.minimum-interval", strconv.FormatInt(data.BfdMinimumInterval.Value, 10))
	}
	if !data.BfdMultiplier.Null && !data.BfdMultiplier.Unknown {
		body, _ = sjson.Set(body, "bfd.multiplier", strconv.FormatInt(data.BfdMultiplier.Value, 10))
	}
	if !data.LocalAsAsNumber.Null && !data.LocalAsAsNumber.Unknown {
		body, _ = sjson.Set(body, "local-as.as-number", data.LocalAsAsNumber.Value)
	}
	if !data.LocalAsNoPrepend.Null && !data.LocalAsNoPrepend.Unknown {
		if data.LocalAsNoPrepend.Value {
			body, _ = sjson.Set(body, "local-as.no-prepend", map[string]string{})
		}
	}
	if !data.LocalAsNoPrependReplaceAs.Null && !data.LocalAsNoPrependReplaceAs.Unknown {
		if data.LocalAsNoPrependReplaceAs.Value {
			body, _ = sjson.Set(body, "local-as.no-prepend.replace-as", map[string]string{})
		}
	}
	if !data.LocalAsNoPrependReplaceAsDualAs.Null && !data.LocalAsNoPrependReplaceAsDualAs.Unknown {
		if data.LocalAsNoPrependReplaceAsDualAs.Value {
			body, _ = sjson.Set(body, "local-as.no-prepend.replace-as.dual-as", map[string]string{})
		}
	}
	if !data.PasswordEncrypted.Null && !data.PasswordEncrypted.Unknown {
		body, _ = sjson.Set(body, "password.encrypted", data.PasswordEncrypted.Value)
	}
	if !data.Shutdown.Null && !data.Shutdown.Unknown {
		if data.Shutdown.Value {
			body, _ = sjson.Set(body, "shutdown", map[string]string{})
		}
	}
	if !data.TimersKeepaliveInterval.Null && !data.TimersKeepaliveInterval.Unknown {
		body, _ = sjson.Set(body, "timers.keepalive-interval", strconv.FormatInt(data.TimersKeepaliveInterval.Value, 10))
	}
	if !data.TimersHoldtime.Null && !data.TimersHoldtime.Unknown {
		body, _ = sjson.Set(body, "timers.holdtime", data.TimersHoldtime.Value)
	}
	if !data.UpdateSource.Null && !data.UpdateSource.Unknown {
		body, _ = sjson.Set(body, "update-source", data.UpdateSource.Value)
	}
	if !data.TtlSecurity.Null && !data.TtlSecurity.Unknown {
		if data.TtlSecurity.Value {
			body, _ = sjson.Set(body, "ttl-security", map[string]string{})
		}
	}

	return body
}

func (data *RouterBGPVRFNeighbor) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "remote-as"); value.Exists() {
		data.RemoteAs.Value = value.String()
	}
	if value := gjson.GetBytes(res, "description"); value.Exists() {
		data.Description.Value = value.String()
	}
	if value := gjson.GetBytes(res, "ignore-connected-check"); value.Exists() {
		data.IgnoreConnectedCheck.Value = true
	}
	if value := gjson.GetBytes(res, "ebgp-multihop.maximum-hop-count"); value.Exists() {
		data.EbgpMultihopMaximumHopCount.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() {
		data.BfdMinimumInterval.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() {
		data.BfdMultiplier.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "local-as.as-number"); value.Exists() {
		data.LocalAsAsNumber.Value = value.String()
	}
	if value := gjson.GetBytes(res, "local-as.no-prepend"); value.Exists() {
		data.LocalAsNoPrepend.Value = true
	}
	if value := gjson.GetBytes(res, "local-as.no-prepend.replace-as"); value.Exists() {
		data.LocalAsNoPrependReplaceAs.Value = true
	}
	if value := gjson.GetBytes(res, "local-as.no-prepend.replace-as.dual-as"); value.Exists() {
		data.LocalAsNoPrependReplaceAsDualAs.Value = true
	}
	if value := gjson.GetBytes(res, "password.encrypted"); value.Exists() {
		data.PasswordEncrypted.Value = value.String()
	}
	if value := gjson.GetBytes(res, "shutdown"); value.Exists() {
		data.Shutdown.Value = true
	}
	if value := gjson.GetBytes(res, "timers.keepalive-interval"); value.Exists() {
		data.TimersKeepaliveInterval.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "timers.holdtime"); value.Exists() {
		data.TimersHoldtime.Value = value.String()
	}
	if value := gjson.GetBytes(res, "update-source"); value.Exists() {
		data.UpdateSource.Value = value.String()
	}
	if value := gjson.GetBytes(res, "ttl-security"); value.Exists() {
		data.TtlSecurity.Value = true
	}
}

func (data *RouterBGPVRFNeighbor) fromPlan(plan RouterBGPVRFNeighbor) {
	data.Device = plan.Device
	data.AsNumber.Value = plan.AsNumber.Value
	data.VrfName.Value = plan.VrfName.Value
	data.NeighborAddress.Value = plan.NeighborAddress.Value
}
