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

type resourceRouterBGPNeighborType struct{}

func (t resourceRouterBGPNeighborType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router BGP Neighbor configuration.",

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
			"as_number": {
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"neighbor_address": {
				MarkdownDescription: helpers.NewAttributeDescription("Neighbor address").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"remote_as": {
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"description": {
				MarkdownDescription: helpers.NewAttributeDescription("Neighbor specific description").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"ignore_connected_check": {
				MarkdownDescription: helpers.NewAttributeDescription("Bypass the directly connected nexthop check for single-hop eBGP peering").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"ebgp_multihop_maximum_hop_count": {
				MarkdownDescription: helpers.NewAttributeDescription("maximum hop count").AddIntegerRangeDescription(1, 255).String,
				Type:                types.Int64Type,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 255),
				},
			},
			"bfd_minimum_interval": {
				MarkdownDescription: helpers.NewAttributeDescription("Hello interval").AddIntegerRangeDescription(3, 30000).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(3, 30000),
				},
			},
			"bfd_multiplier": {
				MarkdownDescription: helpers.NewAttributeDescription("Detect multiplier").AddIntegerRangeDescription(2, 16).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(2, 16),
				},
			},
			"local_as": {
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"local_as_no_prepend": {
				MarkdownDescription: helpers.NewAttributeDescription("Do not prepend local AS to announcements from this neighbor").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"local_as_replace_as": {
				MarkdownDescription: helpers.NewAttributeDescription("Prepend only local AS to announcements to this neighbor").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"local_as_dual_as": {
				MarkdownDescription: helpers.NewAttributeDescription("Dual-AS mode").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"password": {
				MarkdownDescription: helpers.NewAttributeDescription("Specifies an ENCRYPTED password will follow").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `(!.+)|([^!].+)`),
				},
			},
			"shutdown": {
				MarkdownDescription: helpers.NewAttributeDescription("Administratively shut down this neighbor").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"timers_keepalive_interval": {
				MarkdownDescription: helpers.NewAttributeDescription("BGP timers").AddIntegerRangeDescription(0, 65535).String,
				Type:                types.Int64Type,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 65535),
				},
			},
			"timers_holdtime": {
				MarkdownDescription: helpers.NewAttributeDescription("Holdtime. Set 0 to disable keepalives/hold time.").String,
				Type:                types.StringType,
				Required:            true,
			},
			"update_source": {
				MarkdownDescription: helpers.NewAttributeDescription("Source of routing updates").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `[a-zA-Z0-9.:_/-]+`),
				},
			},
			"ttl_security": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable EBGP TTL security").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
		},
	}, nil
}

func (t resourceRouterBGPNeighborType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceRouterBGPNeighbor{
		provider: provider,
	}, diags
}

type resourceRouterBGPNeighbor struct {
	provider provider
}

func (r resourceRouterBGPNeighbor) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state RouterBGPNeighbor

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

func (r resourceRouterBGPNeighbor) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state RouterBGPNeighbor

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

func (r resourceRouterBGPNeighbor) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state RouterBGPNeighbor

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

func (r resourceRouterBGPNeighbor) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state RouterBGPNeighbor

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

func (r resourceRouterBGPNeighbor) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
