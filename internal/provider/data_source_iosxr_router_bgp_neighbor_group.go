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

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RouterBGPNeighborGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterBGPNeighborGroupDataSource{}
)

func NewRouterBGPNeighborGroupDataSource() datasource.DataSource {
	return &RouterBGPNeighborGroupDataSource{}
}

type RouterBGPNeighborGroupDataSource struct {
	client *client.Client
}

func (d *RouterBGPNeighborGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_bgp_neighbor_group"
}

func (d *RouterBGPNeighborGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router BGP Neighbor Group configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: "bgp as-number",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Specify a Neighbor-group",
				Required:            true,
			},
			"remote_as": schema.StringAttribute{
				MarkdownDescription: "bgp as-number",
				Computed:            true,
			},
			"update_source": schema.StringAttribute{
				MarkdownDescription: "Source of routing updates",
				Computed:            true,
			},
			"advertisement_interval_seconds": schema.Int64Attribute{
				MarkdownDescription: "Minimum interval between sending BGP routing updates",
				Computed:            true,
			},
			"advertisement_interval_milliseconds": schema.Int64Attribute{
				MarkdownDescription: "time in milliseconds",
				Computed:            true,
			},
			"ao_key_chain_name": schema.StringAttribute{
				MarkdownDescription: "Name of the key chain - maximum 32 characters",
				Computed:            true,
			},
			"ao_include_tcp_options_enable": schema.BoolAttribute{
				MarkdownDescription: "Include other TCP options in the header",
				Computed:            true,
			},
			"bfd_minimum_interval": schema.Int64Attribute{
				MarkdownDescription: "Hello interval",
				Computed:            true,
			},
			"bfd_multiplier": schema.Int64Attribute{
				MarkdownDescription: "Detect multiplier",
				Computed:            true,
			},
			"bfd_fast_detect": schema.BoolAttribute{
				MarkdownDescription: "Enable Fast detection",
				Computed:            true,
			},
			"bfd_fast_detect_strict_mode": schema.BoolAttribute{
				MarkdownDescription: "Hold down neighbor session until BFD session is up",
				Computed:            true,
			},
			"bfd_fast_detect_inheritance_disable": schema.BoolAttribute{
				MarkdownDescription: "Prevent bfd settings from being inherited from the parent",
				Computed:            true,
			},
			"address_families": schema.ListNestedAttribute{
				MarkdownDescription: "Enter Address Family command mode",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"af_name": schema.StringAttribute{
							MarkdownDescription: "Enter Address Family command mode",
							Computed:            true,
						},
						"soft_reconfiguration_inbound_always": schema.BoolAttribute{
							MarkdownDescription: "Always use soft reconfig, even if route refresh is supported",
							Computed:            true,
						},
						"next_hop_self": schema.BoolAttribute{
							MarkdownDescription: "Disable the next hop calculation for this neighbor",
							Computed:            true,
						},
						"next_hop_self_inheritance_disable": schema.BoolAttribute{
							MarkdownDescription: "Prevent next-hop-self from being inherited from the parent",
							Computed:            true,
						},
						"route_reflector_client": schema.BoolAttribute{
							MarkdownDescription: "Configure a neighbor as Route Reflector client",
							Computed:            true,
						},
						"route_reflector_client_inheritance_disable": schema.BoolAttribute{
							MarkdownDescription: "Prevent route-reflector-client from being inherited from the parent",
							Computed:            true,
						},
						"route_policy_in": schema.StringAttribute{
							MarkdownDescription: "Apply route policy to inbound routes",
							Computed:            true,
						},
						"route_policy_out": schema.StringAttribute{
							MarkdownDescription: "Apply route policy to outbound routes",
							Computed:            true,
						},
					},
				},
			},
			"timers_keepalive_interval": schema.Int64Attribute{
				MarkdownDescription: "BGP timers",
				Computed:            true,
			},
			"timers_holdtime": schema.StringAttribute{
				MarkdownDescription: "Holdtime. Set 0 to disable keepalives/hold time.",
				Computed:            true,
			},
			"timers_minimum_acceptable_holdtime": schema.StringAttribute{
				MarkdownDescription: "Minimum acceptable holdtime from neighbor. Set 0 to disable keepalives/hold time.",
				Computed:            true,
			},
		},
	}
}

func (d *RouterBGPNeighborGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterBGPNeighborGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterBGPNeighborGroupData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, diags := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.fromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
