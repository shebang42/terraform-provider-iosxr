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

type resourceRouterOSPFVRFRedistributeISISType struct{}

func (t resourceRouterOSPFVRFRedistributeISISType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router OSPF VRF Redistribute ISIS configuration.",

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
			"vrf_name": {
				MarkdownDescription: helpers.NewAttributeDescription("Name for this OSPF vrf").String,
				Type:                types.StringType,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(1, 32, `[\w\-\.:,_@#%$\+=\|;]+`),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"instance_name": {
				MarkdownDescription: helpers.NewAttributeDescription("ISO IS-IS").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"level_1": {
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS level-1 routes only").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"level_2": {
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS level-2 routes only").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"level_1_2": {
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS level-1 and level-2 routes").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"tag": {
				MarkdownDescription: helpers.NewAttributeDescription("Set tag for routes redistributed into OSPF").AddIntegerRangeDescription(0, 4294967295).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 4294967295),
				},
			},
			"metric_type": {
				MarkdownDescription: helpers.NewAttributeDescription("OSPF exterior metric type for redistributed routes").AddStringEnumDescription("1", "2").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("1", "2"),
				},
			},
		},
	}, nil
}

func (t resourceRouterOSPFVRFRedistributeISISType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceRouterOSPFVRFRedistributeISIS{
		provider: provider,
	}, diags
}

type resourceRouterOSPFVRFRedistributeISIS struct {
	provider provider
}

func (r resourceRouterOSPFVRFRedistributeISIS) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state RouterOSPFVRFRedistributeISIS

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

func (r resourceRouterOSPFVRFRedistributeISIS) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state RouterOSPFVRFRedistributeISIS

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

func (r resourceRouterOSPFVRFRedistributeISIS) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state RouterOSPFVRFRedistributeISIS

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

func (r resourceRouterOSPFVRFRedistributeISIS) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state RouterOSPFVRFRedistributeISIS

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

func (r resourceRouterOSPFVRFRedistributeISIS) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
