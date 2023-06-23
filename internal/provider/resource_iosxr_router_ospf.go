// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*RouterOSPFResource)(nil)

func NewRouterOSPFResource() resource.Resource {
	return &RouterOSPFResource{}
}

type RouterOSPFResource struct {
	client *client.Client
}

func (r *RouterOSPFResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_ospf"
}

func (r *RouterOSPFResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router OSPF configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"delete_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.").AddStringEnumDescription("all", "attributes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "attributes"),
				},
			},
			"process_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name for this OSPF process").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"mpls_ldp_sync": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable LDP IGP synchronization").String,
				Optional:            true,
			},
			"hello_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Time between HELLO packets").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"dead_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval after which a neighbor is declared dead").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Router priority").AddIntegerRangeDescription(0, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
			},
			"mtu_ignore_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ignores the MTU in DBD packets").String,
				Optional:            true,
			},
			"mtu_ignore_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable ignoring the MTU in DBD packets").String,
				Optional:            true,
			},
			"passive_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable passive").String,
				Optional:            true,
			},
			"passive_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable passive").String,
				Optional:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("configure this node").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
				},
			},
			"redistribute_connected": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Connected routes").String,
				Optional:            true,
			},
			"redistribute_connected_tag": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set tag for routes redistributed into OSPF").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"redistribute_connected_metric_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OSPF exterior metric type for redistributed routes").AddStringEnumDescription("1", "2").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1", "2"),
				},
			},
			"redistribute_static": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static routes").String,
				Optional:            true,
			},
			"redistribute_static_tag": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set tag for routes redistributed into OSPF").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"redistribute_static_metric_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OSPF exterior metric type for redistributed routes").AddStringEnumDescription("1", "2").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1", "2"),
				},
			},
			"bfd_fast_detect": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Fast detection").String,
				Optional:            true,
			},
			"bfd_minimum_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum interval").AddIntegerRangeDescription(3, 30000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 30000),
				},
			},
			"bfd_multiplier": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect multiplier").AddIntegerRangeDescription(2, 50).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 50),
				},
			},
			"default_information_originate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Distribute a default route").String,
				Optional:            true,
			},
			"default_information_originate_always": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Always advertise default route").String,
				Optional:            true,
			},
			"default_information_originate_metric_type": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("OSPF metric type for default routes").AddIntegerRangeDescription(1, 2).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2),
				},
			},
			"areas": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enter the OSPF area configuration submode").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"area_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enter the OSPF area configuration submode").String,
							Required:            true,
						},
					},
				},
			},
			"redistribute_bgp": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
							Required:            true,
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set tag for routes redistributed into OSPF").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"metric_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("OSPF exterior metric type for redistributed routes").AddStringEnumDescription("1", "2").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("1", "2"),
							},
						},
					},
				},
			},
			"redistribute_isis": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ISO IS-IS").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"instance_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ISO IS-IS").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
							},
						},
						"level_1": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IS-IS level-1 routes only").String,
							Optional:            true,
						},
						"level_2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IS-IS level-2 routes only").String,
							Optional:            true,
						},
						"level_1_2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IS-IS level-1 and level-2 routes").String,
							Optional:            true,
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set tag for routes redistributed into OSPF").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"metric_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("OSPF exterior metric type for redistributed routes").AddStringEnumDescription("1", "2").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("1", "2"),
							},
						},
					},
				},
			},
			"redistribute_ospf": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Open Shortest Path First (OSPF)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"instance_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Open Shortest Path First (OSPF)").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
							},
						},
						"match_internal": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute OSPF internal routes").String,
							Optional:            true,
						},
						"match_external": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute OSPF external routes").String,
							Optional:            true,
						},
						"match_nssa_external": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute OSPF NSSA external routes").String,
							Optional:            true,
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set tag for routes redistributed into OSPF").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"metric_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("OSPF exterior metric type for redistributed routes").AddStringEnumDescription("1", "2").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("1", "2"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *RouterOSPFResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RouterOSPFResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouterOSPF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterOSPFResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouterOSPF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	getResp, diags := r.client.Get(ctx, state.Device.ValueString(), state.Id.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.updateFromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterOSPFResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouterOSPF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Update object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterOSPFResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouterOSPF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	var ops []client.SetOperation
	deleteMode := "all"
	if state.DeleteMode.ValueString() == "all" {
		deleteMode = "all"
	} else if state.DeleteMode.ValueString() == "attributes" {
		deleteMode = "attributes"
	}

	if deleteMode == "all" {
		ops = append(ops, client.SetOperation{Path: state.Id.ValueString(), Body: "", Operation: client.Delete})
	} else {
		deletePaths := state.getDeletePaths(ctx)
		tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

		for _, i := range deletePaths {
			ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
		}
	}

	_, diags = r.client.Set(ctx, state.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *RouterOSPFResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
