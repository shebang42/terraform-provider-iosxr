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

type resourceRouterOSPFType struct{}

func (t resourceRouterOSPFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router OSPF configuration.",

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
			"process_name": {
				MarkdownDescription: helpers.NewAttributeDescription("Name for this OSPF process").String,
				Type:                types.StringType,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(1, 32, `[\w\-\.:,_@#%$\+=\|;]+`),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"mpls_ldp_sync": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable LDP IGP synchronization").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"hello_interval": {
				MarkdownDescription: helpers.NewAttributeDescription("Time between HELLO packets").AddIntegerRangeDescription(1, 65535).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 65535),
				},
			},
			"dead_interval": {
				MarkdownDescription: helpers.NewAttributeDescription("Interval after which a neighbor is declared dead").AddIntegerRangeDescription(1, 65535).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 65535),
				},
			},
			"priority": {
				MarkdownDescription: helpers.NewAttributeDescription("Router priority").AddIntegerRangeDescription(0, 255).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 255),
				},
			},
			"mtu_ignore_enable": {
				MarkdownDescription: helpers.NewAttributeDescription("Ignores the MTU in DBD packets").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"mtu_ignore_disable": {
				MarkdownDescription: helpers.NewAttributeDescription("Disable ignoring the MTU in DBD packets").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"passive_enable": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable passive").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"passive_disable": {
				MarkdownDescription: helpers.NewAttributeDescription("Disable passive").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"router_id": {
				MarkdownDescription: helpers.NewAttributeDescription("configure this node").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`, `[0-9\.]*`),
				},
			},
			"redistribute_connected": {
				MarkdownDescription: helpers.NewAttributeDescription("Connected routes").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"redistribute_connected_tag": {
				MarkdownDescription: helpers.NewAttributeDescription("Set tag for routes redistributed into OSPF").AddIntegerRangeDescription(0, 4294967295).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 4294967295),
				},
			},
			"redistribute_connected_metric_type": {
				MarkdownDescription: helpers.NewAttributeDescription("OSPF exterior metric type for redistributed routes").AddStringEnumDescription("1", "2").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("1", "2"),
				},
			},
			"redistribute_static": {
				MarkdownDescription: helpers.NewAttributeDescription("Static routes").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"redistribute_static_tag": {
				MarkdownDescription: helpers.NewAttributeDescription("Set tag for routes redistributed into OSPF").AddIntegerRangeDescription(0, 4294967295).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 4294967295),
				},
			},
			"redistribute_static_metric_type": {
				MarkdownDescription: helpers.NewAttributeDescription("OSPF exterior metric type for redistributed routes").AddStringEnumDescription("1", "2").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("1", "2"),
				},
			},
			"bfd_fast_detect": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable Fast detection").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"bfd_minimum_interval": {
				MarkdownDescription: helpers.NewAttributeDescription("Minimum interval").AddIntegerRangeDescription(3, 30000).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(3, 30000),
				},
			},
			"bfd_multiplier": {
				MarkdownDescription: helpers.NewAttributeDescription("Detect multiplier").AddIntegerRangeDescription(2, 50).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(2, 50),
				},
			},
			"default_information_originate": {
				MarkdownDescription: helpers.NewAttributeDescription("Distribute a default route").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"default_information_originate_always": {
				MarkdownDescription: helpers.NewAttributeDescription("Always advertise default route").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"default_information_originate_metric_type": {
				MarkdownDescription: helpers.NewAttributeDescription("OSPF metric type for default routes").AddIntegerRangeDescription(1, 2).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 2),
				},
			},
		},
	}, nil
}

func (t resourceRouterOSPFType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceRouterOSPF{
		provider: provider,
	}, diags
}

type resourceRouterOSPF struct {
	provider provider
}

func (r resourceRouterOSPF) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state RouterOSPF

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

func (r resourceRouterOSPF) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state RouterOSPF

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

func (r resourceRouterOSPF) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state RouterOSPF

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

func (r resourceRouterOSPF) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state RouterOSPF

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

func (r resourceRouterOSPF) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
