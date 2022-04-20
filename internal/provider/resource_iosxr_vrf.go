// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/helpers"
)

type resourceVRFType struct{}

func (t resourceVRFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the VRF configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			"vrf_name": {
				MarkdownDescription: helpers.NewAttributeDescription("VRF name").String,
				Type:                types.StringType,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(1, 32, `[\w\-\.:,_@#%$\+=\|;]+`),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"description": {
				MarkdownDescription: helpers.NewAttributeDescription("A description for the VRF").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"vpn_id": {
				MarkdownDescription: helpers.NewAttributeDescription("VPN ID, (OUI:VPN-Index) format(hex), 4 bytes VPN_Index Part").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `([0-9a-f]{1,8}):([0-9a-f]{1,8})`),
				},
			},
			"address_family_ipv4_unicast": {
				MarkdownDescription: helpers.NewAttributeDescription("Unicast sub address family").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"address_family_ipv4_multicast": {
				MarkdownDescription: helpers.NewAttributeDescription("Multicast topology").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"address_family_ipv4_flowspec": {
				MarkdownDescription: helpers.NewAttributeDescription("Flowspec sub address family").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"address_family_ipv6_unicast": {
				MarkdownDescription: helpers.NewAttributeDescription("Unicast sub address family").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"address_family_ipv6_multicast": {
				MarkdownDescription: helpers.NewAttributeDescription("Multicast topology").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"address_family_ipv6_flowspec": {
				MarkdownDescription: helpers.NewAttributeDescription("Flowspec sub address family").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"rd_two_byte_as_as_number": {
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"rd_two_byte_as_index": {
				MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 4294967295),
				},
			},
			"rd_four_byte_as_as_number": {
				MarkdownDescription: helpers.NewAttributeDescription("4-byte AS number").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"rd_four_byte_as_index": {
				MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 4294967295),
				},
			},
			"rd_ip_address_ipv4_address": {
				MarkdownDescription: helpers.NewAttributeDescription("configure this node").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`, `[0-9\.]*`),
				},
			},
			"rd_ip_address_index": {
				MarkdownDescription: helpers.NewAttributeDescription("IPv4Address:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 65535),
				},
			},
		},
	}, nil
}

func (t resourceVRFType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceVRF{
		provider: provider,
	}, diags
}

type resourceVRF struct {
	provider provider
}

func (r resourceVRF) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state VRF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody()

	_, diags = r.provider.client.Set(ctx, plan.Device.Value, plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read object
	getResp, diags := r.provider.client.Get(ctx, plan.Device.Value, plan.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.fromBody(getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	state.fromPlan(plan)
	state.Id.Value = plan.getPath()

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceVRF) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state VRF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	getResp, diags := r.provider.client.Get(ctx, state.Device.Value, state.Id.Value)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.fromBody(getResp.Notification[0].Update[0].Val.GetJsonIetfVal())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceVRF) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state VRF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.Value))

	// Update object
	body := plan.toBody()

	_, diags = r.provider.client.Set(ctx, plan.Device.Value, plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read object
	getResp, diags := r.provider.client.Get(ctx, plan.Device.Value, plan.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.fromBody(getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	state.fromPlan(plan)
	state.Id.Value = plan.Id.Value

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceVRF) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state VRF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.Value))

	_, diags = r.provider.client.Set(ctx, state.Device.Value, state.getPath(), "", client.Delete)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.Value))

	resp.State.RemoveResource(ctx)
}

func (r resourceVRF) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
