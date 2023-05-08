// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/helpers"
)

var _ resource.Resource = (*SNMPServerResource)(nil)

func NewSNMPServerResource() resource.Resource {
	return &SNMPServerResource{}
}

type SNMPServerResource struct {
	client *client.Client
}

func (r *SNMPServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_server"
}

func (r *SNMPServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the SNMP Server configuration.",

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
			"rf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP RF-MIB traps").String,
				Optional:            true,
			},
			"bfd": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable BFD traps").String,
				Optional:            true,
			},
			"ntp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP Cisco Ntp traps").String,
				Optional:            true,
			},
			"ethernet_oam_events": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable all OAM event traps").String,
				Optional:            true,
			},
			"copy_complete": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable CISCO-CONFIG-COPY-MIB ccCopyCompletion traps").String,
				Optional:            true,
			},
			"traps_snmp_linkup": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMPv2-MIB linkUp traps").String,
				Optional:            true,
			},
			"traps_snmp_linkdown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMPv2-MIB linDownp traps").String,
				Optional:            true,
			},
			"power": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP entity power traps").String,
				Optional:            true,
			},
			"config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP config traps").String,
				Optional:            true,
			},
			"entity": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP entity traps").String,
				Optional:            true,
			},
			"system": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP SYSTEMMIB-MIB traps").String,
				Optional:            true,
			},
			"bridgemib": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP Trap for Bridge MIB").String,
				Optional:            true,
			},
			"entity_state_operstatus": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable entity oper status enable notification").String,
				Optional:            true,
			},
			"entity_redundancy_all": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable all CISCO-ENTITY-REDUNDANCY-MIB traps").String,
				Optional:            true,
			},
			"trap_source_both": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Assign an interface for the source address of all traps").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
				},
			},
			"l2vpn_all": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable all L2VPN traps").String,
				Optional:            true,
			},
			"l2vpn_vc_up": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable VC up traps").String,
				Optional:            true,
			},
			"l2vpn_vc_down": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable VC down traps").String,
				Optional:            true,
			},
			"sensor": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP entity sensor traps").String,
				Optional:            true,
			},
			"fru_ctrl": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP entity FRU control traps").String,
				Optional:            true,
			},
			"isis_all": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable all IS-IS traps").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_database_overload": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisDatabaseOverload").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_manual_address_drops": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisManualAddressDrops").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_corrupted_lsp_detected": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisCorruptedLSPDetected").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_attempt_to_exceed_max_sequence": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisAttemptToExceedMaxSequence").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_id_len_mismatch": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisIDLenMismatch").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_max_area_addresses_mismatch": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisMaxAreaAddressesMismatch").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_own_lsp_purge": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisOwnLSPPurge").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_sequence_number_skip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisSequenceNumberSkip").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_authentication_type_failure": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisAuthenticationTypeFailure").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_authentication_failure": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisAuthenticationFailure").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_version_skew": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisVersionSkew").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_area_mismatch": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisAreaMismatch").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_rejected_adjacency": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisRejectedAdjacency").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_lsp_too_large_to_propagate": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisLSPTooLargeToPropagate").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_orig_lsp_buff_size_mismatch": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisOrigLSPBuffSizeMismatch").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_protocols_supported_mismatch": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisProtocolsSupportedMismatch").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_adjacency_change": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisAdjacencyChange").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"isis_lsp_error_detected": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("isisLSPErrorDetected").AddStringEnumDescription("disable", "enable").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "enable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"bgp_cbgp2_updown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable CISCO-BGP4-MIB v2 up/down traps").String,
				Optional:            true,
			},
			"bgp_bgp4_mib_updown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable CISCO-BGP4-MIB v2 up/down traps").String,
				Optional:            true,
			},
			"users": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the user").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"user_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the user").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"group_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Group to which the user belongs").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"v3_auth_md5_encryption_aes": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specifies an aes-128 ENCRYPTED authentication password").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-zA-Z]+`), ""),
							},
						},
						"v3_auth_md5_encryption_default": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specifies an default ENCRYPTED authentication password").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(!.+)|([^!].+)`), ""),
							},
						},
					},
				},
			},
			"groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the group").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the group").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"v3_priv": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("group using authPriv Security Level").String,
							Optional:            true,
						},
						"v3_read": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("specify a read view for this group").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"v3_write": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("specify a write view for this group").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"v3_context": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Attach a SNMP context").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"v3_notify": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("specify a notify view for the group").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"v3_ipv4": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of Access-list").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"v3_ipv6": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of Access-list").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
					},
				},
			},
		},
	}
}

func (r *SNMPServerResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *SNMPServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SNMPServer

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

func (r *SNMPServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SNMPServer

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

func (r *SNMPServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SNMPServer

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

func (r *SNMPServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SNMPServer

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	_, diags = r.client.Set(ctx, state.Device.ValueString(), client.SetOperation{Path: state.getPath(), Body: "", Operation: client.Delete})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *SNMPServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
