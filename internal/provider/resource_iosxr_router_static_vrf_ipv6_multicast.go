// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NewRouterStaticVRFIPv6MulticastResource() resource.Resource {
	return &RouterStaticVRFIPv6MulticastResource{}
}

type RouterStaticVRFIPv6MulticastResource struct {
	client *client.Client
}

func (r *RouterStaticVRFIPv6MulticastResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_static_vrf_ipv6_multicast"
}

func (r *RouterStaticVRFIPv6MulticastResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router Static VRF IPv6 Multicast configuration.",

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
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VRF Static route configuration subcommands").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"prefix_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination prefix").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`(([^:]+:){6}(([^:]+:[^:]+)|(.*\..*)))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?)(%.+)?`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F:\.]*`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"prefix_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination prefix length").AddIntegerRangeDescription(0, 128).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 128),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"nexthop_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("description of the static route").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 30),
							},
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set tag for this route").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"distance_metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Distance metric for this route").AddIntegerRangeDescription(1, 254).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 254),
							},
						},
						"permanent": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Permanent route").String,
							Optional:            true,
						},
						"track": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable object tracking for static route").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set metric for this route").AddIntegerRangeDescription(1, 16777214).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 16777214),
							},
						},
					},
				},
			},
			"nexthop_interface_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
							},
						},
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding router's address").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`(([^:]+:){6}(([^:]+:[^:]+)|(.*\..*)))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?)(%.+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F:\.]*`), ""),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("description of the static route").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 30),
							},
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set tag for this route").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"distance_metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Distance metric for this route").AddIntegerRangeDescription(1, 254).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 254),
							},
						},
						"permanent": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Permanent route").String,
							Optional:            true,
						},
						"track": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable object tracking for static route").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set metric for this route").AddIntegerRangeDescription(1, 16777214).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 16777214),
							},
						},
					},
				},
			},
			"nexthop_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Forwarding router's address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding router's address").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`(([^:]+:){6}(([^:]+:[^:]+)|(.*\..*)))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?)(%.+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F:\.]*`), ""),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("description of the static route").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 30),
							},
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set tag for this route").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"distance_metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Distance metric for this route").AddIntegerRangeDescription(1, 254).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 254),
							},
						},
						"permanent": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Permanent route").String,
							Optional:            true,
						},
						"track": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable object tracking for static route").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set metric for this route").AddIntegerRangeDescription(1, 16777214).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 16777214),
							},
						},
					},
				},
			},
			"vrfs": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination VRF").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vrf_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination VRF").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"nexthop_interfaces": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"interface_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
										},
									},
									"description": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("description of the static route").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 30),
										},
									},
									"tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set tag for this route").AddIntegerRangeDescription(1, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 4294967295),
										},
									},
									"distance_metric": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Distance metric for this route").AddIntegerRangeDescription(1, 254).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 254),
										},
									},
									"permanent": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Permanent route").String,
										Optional:            true,
									},
									"track": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable object tracking for static route").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 32),
											stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
										},
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set metric for this route").AddIntegerRangeDescription(1, 16777214).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 16777214),
										},
									},
								},
							},
						},
						"nexthop_interface_addresses": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"interface_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Forwarding interface").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
										},
									},
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Forwarding router's address").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?`), ""),
											stringvalidator.RegexMatches(regexp.MustCompile(`(([^:]+:){6}(([^:]+:[^:]+)|(.*\..*)))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?)(%.+)?`), ""),
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F:\.]*`), ""),
										},
									},
									"description": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("description of the static route").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 30),
										},
									},
									"tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set tag for this route").AddIntegerRangeDescription(1, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 4294967295),
										},
									},
									"distance_metric": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Distance metric for this route").AddIntegerRangeDescription(1, 254).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 254),
										},
									},
									"permanent": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Permanent route").String,
										Optional:            true,
									},
									"track": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable object tracking for static route").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 32),
											stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
										},
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set metric for this route").AddIntegerRangeDescription(1, 16777214).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 16777214),
										},
									},
									"bfd_fast_detect_minimum_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Hello interval").AddIntegerRangeDescription(3, 30000).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(3, 30000),
										},
									},
									"bfd_fast_detect_multiplier": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Detect multiplier").AddIntegerRangeDescription(1, 10).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 10),
										},
									},
								},
							},
						},
						"nexthop_addresses": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Forwarding router's address").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Forwarding router's address").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?`), ""),
											stringvalidator.RegexMatches(regexp.MustCompile(`(([^:]+:){6}(([^:]+:[^:]+)|(.*\..*)))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?)(%.+)?`), ""),
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F:\.]*`), ""),
										},
									},
									"description": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("description of the static route").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 30),
										},
									},
									"tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set tag for this route").AddIntegerRangeDescription(1, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 4294967295),
										},
									},
									"distance_metric": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Distance metric for this route").AddIntegerRangeDescription(1, 254).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 254),
										},
									},
									"permanent": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Permanent route").String,
										Optional:            true,
									},
									"track": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable object tracking for static route").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 32),
											stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
										},
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set metric for this route").AddIntegerRangeDescription(1, 16777214).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 16777214),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *RouterStaticVRFIPv6MulticastResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RouterStaticVRFIPv6MulticastResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouterStaticVRFIPv6Multicast

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

func (r *RouterStaticVRFIPv6MulticastResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouterStaticVRFIPv6Multicast

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	import_ := false
	if state.Id.ValueString() == "" {
		import_ = true
		state.Id = types.StringValue(state.getPath())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	getResp, diags := r.client.Get(ctx, state.Device.ValueString(), state.Id.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	respBody := getResp.Notification[0].Update[0].Val.GetJsonIetfVal()
	if import_ {
		state.fromBody(ctx, respBody)
	} else {
		state.updateFromBody(ctx, respBody)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterStaticVRFIPv6MulticastResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouterStaticVRFIPv6Multicast

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

	deletedListItems := plan.getDeletedItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("Removed items to delete: %+v", deletedListItems))

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

func (r *RouterStaticVRFIPv6MulticastResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouterStaticVRFIPv6Multicast

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

func (r *RouterStaticVRFIPv6MulticastResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 3 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <vrf_name>,<prefix_address>,<prefix_length>. Got: %q", req.ID),
		)
		return
	}
	value0 := idParts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vrf_name"), value0)...)
	value1 := idParts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("prefix_address"), value1)...)
	value2, _ := strconv.Atoi(idParts[2])
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("prefix_length"), value2)...)
}
